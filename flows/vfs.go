package flows

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"path"
	actions_proto "www.velocidex.com/golang/velociraptor/actions/proto"
	config "www.velocidex.com/golang/velociraptor/config"
	constants "www.velocidex.com/golang/velociraptor/constants"
	crypto_proto "www.velocidex.com/golang/velociraptor/crypto/proto"
	datastore "www.velocidex.com/golang/velociraptor/datastore"
	flows_proto "www.velocidex.com/golang/velociraptor/flows/proto"
	"www.velocidex.com/golang/velociraptor/responder"
	urns "www.velocidex.com/golang/velociraptor/urns"
)

const (
	_VFSListDirectory_ListDir          uint64 = 1
	_VFSListDirectory_RecursiveListDir uint64 = 2
)

type VFSListDirectory struct {
	state *flows_proto.VFSListRequestState
	rows  []map[string]interface{}
}

func (self *VFSListDirectory) Load(flow_obj *AFF4FlowObject) error {
	message := flow_obj.GetState()
	if message == nil {
		message = &flows_proto.VFSListRequestState{
			Current: &actions_proto.VQLResponse{},
		}
	}
	self.state = message.(*flows_proto.VFSListRequestState)
	return json.Unmarshal([]byte(self.state.Current.Response), &self.rows)
}

func (self *VFSListDirectory) Save(flow_obj *AFF4FlowObject) error {
	s, err := json.Marshal(self.rows)
	if err != nil {
		return err
	}
	self.state.Current.Response = string(s)
	flow_obj.SetState(self.state)
	return nil
}

func (self *VFSListDirectory) Start(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject,
	args proto.Message) (*string, error) {

	vfs_args, ok := args.(*flows_proto.VFSListRequest)
	if !ok {
		return nil, errors.New("Expected args of type VQLCollectorArgs")
	}

	glob := "'/*'"
	next_state := _VFSListDirectory_ListDir
	if vfs_args.RecursionDepth > 0 {
		glob = fmt.Sprintf("'/**%d'", vfs_args.RecursionDepth)
		next_state = _VFSListDirectory_RecursiveListDir
	}

	vql_collector_args := &actions_proto.VQLCollectorArgs{
		// Injecting the path in the environment avoids the
		// need to escape it within the query itself and it is
		// therefore more robust.
		Env: []*actions_proto.VQLEnv{
			{Key: "path", Value: path.Join("/", vfs_args.VfsPath)},
		},
		Query: []*actions_proto.VQLRequest{
			{
				Name: vfs_args.VfsPath,
				VQL: "SELECT IsDir, FullPath as _FullPath, " +
					"Name, Size, Mode, " +
					"timestamp(epoch=Sys.Mtim.Sec) as mtime, " +
					"timestamp(epoch=Sys.Atim.Sec) as atime, " +
					"timestamp(epoch=Sys.Ctim.Sec) as ctime " +
					"from glob(globs=path + " + glob + ")",
			},
		},
		MaxRow: uint64(10000),
	}

	db, err := datastore.GetDB(config_obj)
	if err != nil {
		return nil, err
	}

	flow_id := GetNewFlowIdForClient(flow_obj.RunnerArgs.ClientId)
	err = db.QueueMessageForClient(
		config_obj, flow_obj.RunnerArgs.ClientId,
		flow_id,
		"VQLClientAction",
		vql_collector_args, next_state)
	if err != nil {
		return nil, err
	}

	self.state = &flows_proto.VFSListRequestState{
		Current: &actions_proto.VQLResponse{
			Query: vql_collector_args.Query[0],
		},
	}

	return &flow_id, nil
}

func (self *VFSListDirectory) ProcessMessage(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject,
	message *crypto_proto.GrrMessage) error {

	switch message.RequestId {
	case _VFSListDirectory_ListDir:
		return self.processSingleDirectoryListing(
			config_obj, flow_obj, message)

	case _VFSListDirectory_RecursiveListDir:
		return self.processRecursiveDirectoryListing(
			config_obj, flow_obj, message)
	}

	return nil
}

func (self *VFSListDirectory) processSingleDirectoryListing(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject,
	message *crypto_proto.GrrMessage) error {

	var tmp_args ptypes.DynamicAny
	err := ptypes.UnmarshalAny(flow_obj.RunnerArgs.Args, &tmp_args)
	if err != nil {
		return err
	}

	vfs_args := tmp_args.Message.(*flows_proto.VFSListRequest)
	err = flow_obj.FailIfError(message)
	if err != nil {
		return err
	}

	db, err := datastore.GetDB(config_obj)
	if err != nil {
		return err
	}

	if flow_obj.IsRequestComplete(message) {
		flow_obj.Complete()
		return nil
	}

	if message.ArgsRdfName != "VQLResponse" {
		return errors.New("Unexpected response type " + message.ArgsRdfName)
	}

	data := make(map[string][]byte)
	data[constants.VFS_FILE_LISTING] = message.Args

	urn := urns.BuildURN(
		flow_obj.RunnerArgs.ClientId, "vfs",
		vfs_args.VfsPath)

	fmt.Printf("Storing in urn %v", urn)

	return db.SetSubjectData(config_obj, urn, 0, data)
}

// Handle recursive VFS traversal. NOTE: This algorithm relies on the
// fact that the recursive glob (** wildcard) returns files in breadth
// first order. We keep track of the previous directory and add rows
// to the current collection as long as they belong to the current
// directory. When we see a file which belongs in another directory,
// we can flush the current collection and start a new one.
func (self *VFSListDirectory) processRecursiveDirectoryListing(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject,
	message *crypto_proto.GrrMessage) error {

	if flow_obj.IsRequestComplete(message) {
		err := self.flush_state(config_obj, flow_obj)
		if err != nil {
			return err
		}
		flow_obj.Complete()
		return nil
	}

	vql_response, ok := responder.ExtractGrrMessagePayload(
		message).(*actions_proto.VQLResponse)
	if !ok {
		return errors.New("Unexpected response type " + message.ArgsRdfName)
	}

	// Inspect each row and check if it belongs to the current
	// collection.
	var rows []map[string]interface{}
	err := json.Unmarshal([]byte(vql_response.Response), &rows)
	if err != nil {
		return err
	}

	for _, row := range rows {
		full_path, ok := (row["_FullPath"]).(string)
		if ok {
			// This row does not belong in the current
			// collection - flush the collection and start
			// a new one.
			if path.Dir(full_path) != self.state.VfsPath {
				// VfsPath == "" represents the first
				// collection before the first row is
				// processed.
				if self.state.VfsPath != "" {
					err := self.flush_state(config_obj, flow_obj)
					if err != nil {
						return err
					}
				}
				self.state.VfsPath = path.Dir(full_path)
				self.state.Current = vql_response
			}
			self.rows = append(self.rows, row)
		}
	}

	return nil
}

// Flush the current state into the database and clear it for the next directory.
func (self *VFSListDirectory) flush_state(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject) error {
	err := self.Save(flow_obj)
	if err != nil {
		return err
	}
	self.rows = nil

	urn := urns.BuildURN(
		flow_obj.RunnerArgs.ClientId, "vfs",
		self.state.VfsPath)

	fmt.Printf("Storing in urn %v\n", urn)
	s, err := proto.Marshal(self.state.Current)
	if err != nil {
		return err
	}

	data := make(map[string][]byte)
	data[constants.VFS_FILE_LISTING] = s

	db, err := datastore.GetDB(config_obj)
	if err != nil {
		return err
	}

	return db.SetSubjectData(config_obj, urn, 0, data)
}

func init() {
	impl := VFSListDirectory{}
	default_args, _ := ptypes.MarshalAny(&flows_proto.VFSListRequest{})
	desc := &flows_proto.FlowDescriptor{
		Name:         "VFSListDirectory",
		FriendlyName: "List VFS Directory",
		Category:     "Collectors",
		Doc:          "List a single directory in the client's filesystem.",
		ArgsType:     "VFSListRequest",
		DefaultArgs:  default_args,
	}
	RegisterImplementation(desc, &impl)
}

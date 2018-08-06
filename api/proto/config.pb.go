// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientConfig struct {
	Name                   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description            string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Version                string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Commit                 string   `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	BuildTime              string   `protobuf:"bytes,5,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	Labels                 []string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	PrivateKey             string   `protobuf:"bytes,7,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	ServerUrls             []string `protobuf:"bytes,8,rep,name=server_urls,json=serverUrls,proto3" json:"server_urls,omitempty"`
	CaCertificate          string   `protobuf:"bytes,11,opt,name=ca_certificate,json=caCertificate,proto3" json:"ca_certificate,omitempty"`
	Nonce                  string   `protobuf:"bytes,12,opt,name=nonce,proto3" json:"nonce,omitempty"`
	WritebackLinux         string   `protobuf:"bytes,9,opt,name=writeback_linux,json=writebackLinux,proto3" json:"writeback_linux,omitempty"`
	WritebackWindows       string   `protobuf:"bytes,10,opt,name=writeback_windows,json=writebackWindows,proto3" json:"writeback_windows,omitempty"`
	HuntLastTimestamp      uint64   `protobuf:"varint,13,opt,name=hunt_last_timestamp,json=huntLastTimestamp,proto3" json:"hunt_last_timestamp,omitempty"`
	LastServerSerialNumber uint64   `protobuf:"varint,14,opt,name=last_server_serial_number,json=lastServerSerialNumber,proto3" json:"last_server_serial_number,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ClientConfig) Reset()         { *m = ClientConfig{} }
func (m *ClientConfig) String() string { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()    {}
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{0}
}
func (m *ClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfig.Unmarshal(m, b)
}
func (m *ClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfig.Marshal(b, m, deterministic)
}
func (dst *ClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfig.Merge(dst, src)
}
func (m *ClientConfig) XXX_Size() int {
	return xxx_messageInfo_ClientConfig.Size(m)
}
func (m *ClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfig proto.InternalMessageInfo

func (m *ClientConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClientConfig) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ClientConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClientConfig) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *ClientConfig) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

func (m *ClientConfig) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ClientConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *ClientConfig) GetServerUrls() []string {
	if m != nil {
		return m.ServerUrls
	}
	return nil
}

func (m *ClientConfig) GetCaCertificate() string {
	if m != nil {
		return m.CaCertificate
	}
	return ""
}

func (m *ClientConfig) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *ClientConfig) GetWritebackLinux() string {
	if m != nil {
		return m.WritebackLinux
	}
	return ""
}

func (m *ClientConfig) GetWritebackWindows() string {
	if m != nil {
		return m.WritebackWindows
	}
	return ""
}

func (m *ClientConfig) GetHuntLastTimestamp() uint64 {
	if m != nil {
		return m.HuntLastTimestamp
	}
	return 0
}

func (m *ClientConfig) GetLastServerSerialNumber() uint64 {
	if m != nil {
		return m.LastServerSerialNumber
	}
	return 0
}

type APIConfig struct {
	BindAddress          string   `protobuf:"bytes,1,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort             uint32   `protobuf:"varint,2,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIConfig) Reset()         { *m = APIConfig{} }
func (m *APIConfig) String() string { return proto.CompactTextString(m) }
func (*APIConfig) ProtoMessage()    {}
func (*APIConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{1}
}
func (m *APIConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIConfig.Unmarshal(m, b)
}
func (m *APIConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIConfig.Marshal(b, m, deterministic)
}
func (dst *APIConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIConfig.Merge(dst, src)
}
func (m *APIConfig) XXX_Size() int {
	return xxx_messageInfo_APIConfig.Size(m)
}
func (m *APIConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_APIConfig.DiscardUnknown(m)
}

var xxx_messageInfo_APIConfig proto.InternalMessageInfo

func (m *APIConfig) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *APIConfig) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

type GUIConfig struct {
	BindAddress          string   `protobuf:"bytes,1,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort             uint32   `protobuf:"varint,2,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	InternalCidr         []string `protobuf:"bytes,3,rep,name=internal_cidr,json=internalCidr,proto3" json:"internal_cidr,omitempty"`
	VpnCidr              []string `protobuf:"bytes,4,rep,name=vpn_cidr,json=vpnCidr,proto3" json:"vpn_cidr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GUIConfig) Reset()         { *m = GUIConfig{} }
func (m *GUIConfig) String() string { return proto.CompactTextString(m) }
func (*GUIConfig) ProtoMessage()    {}
func (*GUIConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{2}
}
func (m *GUIConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GUIConfig.Unmarshal(m, b)
}
func (m *GUIConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GUIConfig.Marshal(b, m, deterministic)
}
func (dst *GUIConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GUIConfig.Merge(dst, src)
}
func (m *GUIConfig) XXX_Size() int {
	return xxx_messageInfo_GUIConfig.Size(m)
}
func (m *GUIConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GUIConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GUIConfig proto.InternalMessageInfo

func (m *GUIConfig) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *GUIConfig) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *GUIConfig) GetInternalCidr() []string {
	if m != nil {
		return m.InternalCidr
	}
	return nil
}

func (m *GUIConfig) GetVpnCidr() []string {
	if m != nil {
		return m.VpnCidr
	}
	return nil
}

type CAConfig struct {
	PrivateKey           string   `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CAConfig) Reset()         { *m = CAConfig{} }
func (m *CAConfig) String() string { return proto.CompactTextString(m) }
func (*CAConfig) ProtoMessage()    {}
func (*CAConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{3}
}
func (m *CAConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CAConfig.Unmarshal(m, b)
}
func (m *CAConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CAConfig.Marshal(b, m, deterministic)
}
func (dst *CAConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CAConfig.Merge(dst, src)
}
func (m *CAConfig) XXX_Size() int {
	return xxx_messageInfo_CAConfig.Size(m)
}
func (m *CAConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CAConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CAConfig proto.InternalMessageInfo

func (m *CAConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

type FrontendConfig struct {
	BindAddress          string   `protobuf:"bytes,1,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort             uint32   `protobuf:"varint,2,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	Certificate          string   `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	PrivateKey           string   `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	ClientLeaseTime      uint32   `protobuf:"varint,5,opt,name=client_lease_time,json=clientLeaseTime,proto3" json:"client_lease_time,omitempty"`
	DnsName              string   `protobuf:"bytes,6,opt,name=dns_name,json=dnsName,proto3" json:"dns_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontendConfig) Reset()         { *m = FrontendConfig{} }
func (m *FrontendConfig) String() string { return proto.CompactTextString(m) }
func (*FrontendConfig) ProtoMessage()    {}
func (*FrontendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{4}
}
func (m *FrontendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontendConfig.Unmarshal(m, b)
}
func (m *FrontendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontendConfig.Marshal(b, m, deterministic)
}
func (dst *FrontendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendConfig.Merge(dst, src)
}
func (m *FrontendConfig) XXX_Size() int {
	return xxx_messageInfo_FrontendConfig.Size(m)
}
func (m *FrontendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendConfig proto.InternalMessageInfo

func (m *FrontendConfig) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *FrontendConfig) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *FrontendConfig) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *FrontendConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *FrontendConfig) GetClientLeaseTime() uint32 {
	if m != nil {
		return m.ClientLeaseTime
	}
	return 0
}

func (m *FrontendConfig) GetDnsName() string {
	if m != nil {
		return m.DnsName
	}
	return ""
}

type DatastoreConfig struct {
	Implementation       string   `protobuf:"bytes,1,opt,name=implementation,proto3" json:"implementation,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	FilestoreDirectory   string   `protobuf:"bytes,3,opt,name=filestore_directory,json=filestoreDirectory,proto3" json:"filestore_directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatastoreConfig) Reset()         { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()    {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{5}
}
func (m *DatastoreConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatastoreConfig.Unmarshal(m, b)
}
func (m *DatastoreConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatastoreConfig.Marshal(b, m, deterministic)
}
func (dst *DatastoreConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatastoreConfig.Merge(dst, src)
}
func (m *DatastoreConfig) XXX_Size() int {
	return xxx_messageInfo_DatastoreConfig.Size(m)
}
func (m *DatastoreConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatastoreConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatastoreConfig proto.InternalMessageInfo

func (m *DatastoreConfig) GetImplementation() string {
	if m != nil {
		return m.Implementation
	}
	return ""
}

func (m *DatastoreConfig) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *DatastoreConfig) GetFilestoreDirectory() string {
	if m != nil {
		return m.FilestoreDirectory
	}
	return ""
}

type FlowsConfig struct {
	InterrogateAdditionalQueries *proto1.VQLCollectorArgs `protobuf:"bytes,1,opt,name=interrogate_additional_queries,json=interrogateAdditionalQueries,proto3" json:"interrogate_additional_queries,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                 `json:"-"`
	XXX_unrecognized             []byte                   `json:"-"`
	XXX_sizecache                int32                    `json:"-"`
}

func (m *FlowsConfig) Reset()         { *m = FlowsConfig{} }
func (m *FlowsConfig) String() string { return proto.CompactTextString(m) }
func (*FlowsConfig) ProtoMessage()    {}
func (*FlowsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{6}
}
func (m *FlowsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowsConfig.Unmarshal(m, b)
}
func (m *FlowsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowsConfig.Marshal(b, m, deterministic)
}
func (dst *FlowsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowsConfig.Merge(dst, src)
}
func (m *FlowsConfig) XXX_Size() int {
	return xxx_messageInfo_FlowsConfig.Size(m)
}
func (m *FlowsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FlowsConfig proto.InternalMessageInfo

func (m *FlowsConfig) GetInterrogateAdditionalQueries() *proto1.VQLCollectorArgs {
	if m != nil {
		return m.InterrogateAdditionalQueries
	}
	return nil
}

type Config struct {
	Client               *ClientConfig    `protobuf:"bytes,1,opt,name=Client,proto3" json:"Client,omitempty"`
	API                  *APIConfig       `protobuf:"bytes,2,opt,name=API,proto3" json:"API,omitempty"`
	GUI                  *GUIConfig       `protobuf:"bytes,3,opt,name=GUI,proto3" json:"GUI,omitempty"`
	CA                   *CAConfig        `protobuf:"bytes,4,opt,name=CA,proto3" json:"CA,omitempty"`
	Frontend             *FrontendConfig  `protobuf:"bytes,5,opt,name=Frontend,proto3" json:"Frontend,omitempty"`
	Datastore            *DatastoreConfig `protobuf:"bytes,6,opt,name=Datastore,proto3" json:"Datastore,omitempty"`
	Flows                *FlowsConfig     `protobuf:"bytes,7,opt,name=Flows,proto3" json:"Flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ee6bacce45ad81b4, []int{7}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetClient() *ClientConfig {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Config) GetAPI() *APIConfig {
	if m != nil {
		return m.API
	}
	return nil
}

func (m *Config) GetGUI() *GUIConfig {
	if m != nil {
		return m.GUI
	}
	return nil
}

func (m *Config) GetCA() *CAConfig {
	if m != nil {
		return m.CA
	}
	return nil
}

func (m *Config) GetFrontend() *FrontendConfig {
	if m != nil {
		return m.Frontend
	}
	return nil
}

func (m *Config) GetDatastore() *DatastoreConfig {
	if m != nil {
		return m.Datastore
	}
	return nil
}

func (m *Config) GetFlows() *FlowsConfig {
	if m != nil {
		return m.Flows
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientConfig)(nil), "proto.ClientConfig")
	proto.RegisterType((*APIConfig)(nil), "proto.APIConfig")
	proto.RegisterType((*GUIConfig)(nil), "proto.GUIConfig")
	proto.RegisterType((*CAConfig)(nil), "proto.CAConfig")
	proto.RegisterType((*FrontendConfig)(nil), "proto.FrontendConfig")
	proto.RegisterType((*DatastoreConfig)(nil), "proto.DatastoreConfig")
	proto.RegisterType((*FlowsConfig)(nil), "proto.FlowsConfig")
	proto.RegisterType((*Config)(nil), "proto.Config")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_ee6bacce45ad81b4) }

var fileDescriptor_config_ee6bacce45ad81b4 = []byte{
	// 1787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4b, 0x73, 0x5c, 0x47,
	0x15, 0xae, 0x2b, 0xc9, 0xb2, 0xd4, 0x23, 0xc9, 0x76, 0x9b, 0x24, 0x83, 0x89, 0x71, 0x67, 0x12,
	0xb0, 0x1c, 0x52, 0x57, 0xd2, 0xc4, 0x31, 0xb1, 0x0b, 0x57, 0x32, 0x1a, 0x3f, 0x10, 0x51, 0x54,
	0xca, 0xb5, 0x15, 0x53, 0x29, 0xa8, 0xa1, 0xe7, 0xde, 0x33, 0x33, 0x6d, 0xf7, 0x74, 0xdf, 0x74,
	0xf7, 0x68, 0xac, 0x1d, 0xb0, 0x82, 0x3f, 0xc0, 0xab, 0x60, 0xc3, 0xa3, 0x8a, 0x05, 0x0b, 0xf6,
	0x6c, 0xf9, 0x1d, 0x2c, 0x60, 0xc3, 0x9e, 0x2d, 0x45, 0x51, 0x7d, 0xba, 0x67, 0xe6, 0xce, 0xc4,
	0x14, 0x2c, 0xb2, 0x60, 0x35, 0x53, 0xdd, 0xdf, 0xf9, 0xce, 0x77, 0xba, 0xcf, 0xa3, 0x2f, 0xd9,
	0xc8, 0xb5, 0xea, 0x89, 0x7e, 0x5a, 0x1a, 0xed, 0x34, 0x3d, 0x87, 0x3f, 0x57, 0xee, 0x8c, 0xc7,
	0xe3, 0xf4, 0x14, 0xa4, 0xce, 0x45, 0x01, 0xcf, 0xd3, 0x5c, 0x0f, 0x77, 0xfa, 0x5a, 0x72, 0xd5,
	0xdf, 0x09, 0x8b, 0x86, 0x97, 0x4e, 0x9b, 0x1d, 0x04, 0xef, 0x58, 0x18, 0x72, 0xe5, 0x44, 0x1e,
	0x28, 0xae, 0xdc, 0xfd, 0xdf, 0x6c, 0x79, 0xee, 0x84, 0x56, 0x36, 0x72, 0x9c, 0x7e, 0x2a, 0x83,
	0x79, 0xe3, 0x1f, 0x35, 0xb2, 0xd1, 0x96, 0x02, 0x94, 0x6b, 0xa3, 0x30, 0x9a, 0x92, 0x15, 0xc5,
	0x87, 0x50, 0x4f, 0x58, 0xb2, 0xbd, 0xbe, 0x7f, 0xe5, 0xaf, 0xff, 0xfc, 0xdb, 0x9f, 0x93, 0x2f,
	0x50, 0xfa, 0x78, 0x00, 0x2c, 0x47, 0xdc, 0x75, 0xcb, 0x3c, 0x20, 0xcd, 0x10, 0x47, 0x8f, 0x48,
	0xad, 0x00, 0x9b, 0x1b, 0x51, 0x7a, 0x07, 0xf5, 0x25, 0x34, 0x7b, 0x0b, 0xcd, 0xbe, 0x4a, 0xdf,
	0x98, 0x33, 0x93, 0x5a, 0xf5, 0x59, 0x05, 0xcc, 0xac, 0x33, 0x42, 0xf5, 0xb3, 0x2a, 0x01, 0xbd,
	0x4b, 0xce, 0x9f, 0x82, 0xb1, 0x9e, 0x6b, 0x19, 0xb9, 0x5e, 0x47, 0xae, 0xab, 0xf4, 0x4b, 0x73,
	0x5c, 0x11, 0x33, 0xa1, 0x98, 0xd8, 0xd0, 0xf7, 0xc8, 0x6a, 0xae, 0x87, 0x43, 0xe1, 0xea, 0x2b,
	0x68, 0x7d, 0x1d, 0xad, 0x5f, 0xa3, 0xd7, 0xe6, 0xac, 0xfb, 0xc2, 0xb1, 0x00, 0x8b, 0x04, 0x69,
	0x16, 0xcd, 0x68, 0x8b, 0x90, 0xee, 0x48, 0xc8, 0xa2, 0xe3, 0xc4, 0x10, 0xea, 0xe7, 0x90, 0xa4,
	0x81, 0x24, 0xaf, 0xd2, 0x2b, 0x4f, 0x06, 0xa0, 0x98, 0x9b, 0x32, 0xb1, 0x31, 0xb7, 0xcc, 0xa3,
	0x5d, 0x9a, 0xad, 0xa3, 0xd5, 0x63, 0x31, 0x04, 0xfa, 0x83, 0x84, 0xac, 0x4a, 0xde, 0x05, 0x69,
	0xeb, 0xab, 0x6c, 0x79, 0x7b, 0x7d, 0x5f, 0xa0, 0x7d, 0x4e, 0x79, 0x8b, 0x49, 0x61, 0x1d, 0xd3,
	0x3d, 0x16, 0xf6, 0xab, 0x5c, 0x03, 0x6e, 0x53, 0xf6, 0x78, 0x20, 0x2c, 0xe3, 0x52, 0xea, 0xb1,
	0x65, 0x16, 0x24, 0xe4, 0x0e, 0x0a, 0xd6, 0x37, 0x7a, 0x54, 0x5a, 0x6f, 0x15, 0xa0, 0x96, 0x39,
	0xcd, 0xba, 0xc0, 0x1c, 0x37, 0x7d, 0x70, 0x1e, 0x21, 0x14, 0x1b, 0x8c, 0x94, 0xb3, 0x69, 0x16,
	0x1d, 0xd3, 0x8c, 0xd4, 0x4a, 0x23, 0x4e, 0xb9, 0x83, 0xce, 0x33, 0x38, 0xab, 0x9f, 0xc7, 0x38,
	0xf6, 0x50, 0xc7, 0xd7, 0xe8, 0x8d, 0xb9, 0xc3, 0x88, 0x38, 0xf6, 0x0c, 0xce, 0x3c, 0xd1, 0xf1,
	0xfd, 0x0f, 0x19, 0xa8, 0x5c, 0x17, 0x78, 0x2c, 0x24, 0xee, 0x7e, 0x00, 0x67, 0x94, 0x93, 0x9a,
	0x05, 0x73, 0x0a, 0xa6, 0x33, 0x32, 0xd2, 0xd6, 0xd7, 0x30, 0xb6, 0xf7, 0x91, 0xf3, 0x0e, 0x7d,
	0x77, 0x16, 0x5b, 0x00, 0xb1, 0x93, 0xec, 0x70, 0x2e, 0xc0, 0xb1, 0x90, 0x92, 0x39, 0x73, 0xe6,
	0xd5, 0xe7, 0x5a, 0x29, 0xc8, 0x1d, 0x73, 0x3a, 0xcd, 0x48, 0xc0, 0x9f, 0x18, 0x69, 0xe9, 0x8f,
	0x13, 0xb2, 0x95, 0xf3, 0x4e, 0x0e, 0xc6, 0x89, 0x9e, 0xc8, 0xb9, 0x83, 0x7a, 0x0d, 0xa5, 0x77,
	0xd1, 0xcd, 0x77, 0xe8, 0x27, 0x5e, 0x7a, 0xbb, 0x75, 0xdd, 0xb2, 0x0a, 0x26, 0xe8, 0x0d, 0xe7,
	0x70, 0x7c, 0xff, 0x43, 0x7f, 0x8c, 0x8b, 0x5e, 0x47, 0xd6, 0x31, 0xae, 0xce, 0x26, 0xda, 0xac,
	0xe8, 0x2b, 0x28, 0xd8, 0x58, 0xb8, 0x01, 0x73, 0xfe, 0xcc, 0xdb, 0xad, 0x34, 0xdb, 0xcc, 0x79,
	0x7b, 0x46, 0x4a, 0xff, 0x94, 0x90, 0x73, 0x4a, 0xab, 0x1c, 0xea, 0x1b, 0x28, 0xe1, 0xb7, 0x09,
	0x6a, 0xf8, 0x55, 0x42, 0x7f, 0x91, 0xb4, 0x98, 0x1d, 0x70, 0x03, 0x05, 0x43, 0xc0, 0x67, 0xa2,
	0x2c, 0x0d, 0x58, 0x50, 0x2e, 0x08, 0x89, 0x2e, 0x71, 0xc3, 0xc0, 0x53, 0x1f, 0xf6, 0xe4, 0x1e,
	0xc7, 0x03, 0x91, 0x0f, 0x58, 0xa1, 0x99, 0xd2, 0x6e, 0x62, 0x14, 0xf4, 0xf8, 0xcb, 0xb0, 0x9a,
	0x09, 0xc7, 0x86, 0x5e, 0x7a, 0x17, 0x18, 0x0c, 0xbb, 0x50, 0xc4, 0x20, 0x5d, 0xf5, 0xfe, 0x42,
	0x47, 0x19, 0x19, 0xee, 0xeb, 0x27, 0xcd, 0x82, 0x66, 0x9a, 0x93, 0x0b, 0x63, 0x23, 0x1c, 0x74,
	0x79, 0xfe, 0xac, 0x23, 0x85, 0x1a, 0x3d, 0xaf, 0xaf, 0x63, 0x18, 0x77, 0x30, 0x8a, 0x9b, 0xb4,
	0xd9, 0x62, 0x25, 0x77, 0x03, 0x36, 0x1e, 0x80, 0x01, 0x86, 0x88, 0x99, 0x2e, 0xaf, 0x15, 0xcd,
	0x99, 0xd4, 0x39, 0x97, 0xcc, 0x3a, 0xee, 0x20, 0xcd, 0xb6, 0xa6, 0x94, 0x87, 0x1e, 0x4f, 0x05,
	0xb9, 0x34, 0x73, 0x32, 0x16, 0xaa, 0xd0, 0x63, 0x5b, 0x27, 0xe8, 0xe6, 0x1b, 0xe8, 0xe6, 0x16,
	0xbd, 0x39, 0xe7, 0x26, 0x62, 0xfe, 0xab, 0xa3, 0x8b, 0x53, 0xda, 0x27, 0xc1, 0x82, 0x7e, 0x8f,
	0x5c, 0xf6, 0x29, 0xde, 0x91, 0xdc, 0x3a, 0xac, 0x4d, 0xeb, 0xf8, 0xb0, 0xac, 0x6f, 0xb2, 0x64,
	0x7b, 0x65, 0x7f, 0x17, 0x9d, 0xbd, 0x49, 0xb7, 0xfd, 0x79, 0x7b, 0x04, 0x96, 0x03, 0x9b, 0xc2,
	0xc2, 0x99, 0xc6, 0x1b, 0x32, 0x5c, 0xa5, 0xd9, 0x25, 0x0f, 0x38, 0xe4, 0xd6, 0x3d, 0x9e, 0x60,
	0xe8, 0xdf, 0x13, 0xf2, 0x45, 0x64, 0x8f, 0x49, 0x6e, 0xc1, 0x08, 0x2e, 0x3b, 0x6a, 0x34, 0xec,
	0x82, 0xa9, 0x6f, 0xa1, 0xa3, 0x3f, 0x84, 0x1c, 0xf8, 0x5d, 0x42, 0x7f, 0x9d, 0x4c, 0x7d, 0x55,
	0x33, 0x31, 0x18, 0xb1, 0x60, 0xc4, 0xc6, 0xc0, 0x2c, 0x1f, 0xb3, 0x9e, 0xd1, 0x43, 0xbc, 0xb3,
	0xc0, 0x9c, 0xb2, 0x76, 0x0c, 0xdf, 0x40, 0x6f, 0x64, 0x61, 0xbe, 0x20, 0x22, 0xc8, 0x86, 0xe4,
	0xd4, 0xb2, 0xf0, 0xe9, 0x5a, 0x65, 0xb5, 0x8c, 0x5b, 0x56, 0x72, 0x53, 0xad, 0x35, 0x9f, 0x2e,
	0x46, 0x3b, 0x4c, 0x01, 0x86, 0x9d, 0x3d, 0xd7, 0x32, 0xcd, 0x5e, 0xf6, 0xfa, 0x1e, 0x21, 0xe4,
	0x11, 0x52, 0x1c, 0x21, 0x43, 0xe3, 0x2f, 0x09, 0x59, 0x6f, 0x1d, 0x1f, 0xc4, 0x96, 0xff, 0xf3,
	0x84, 0x6c, 0x74, 0x85, 0x2a, 0x3a, 0xbc, 0x28, 0x0c, 0x58, 0x1b, 0x7b, 0xff, 0x29, 0x86, 0x5a,
	0x52, 0xd5, 0x0a, 0xcb, 0xd8, 0x75, 0x84, 0x2a, 0x58, 0x3f, 0x3b, 0x6e, 0x33, 0x50, 0x45, 0xa9,
	0x45, 0xc8, 0x71, 0x61, 0x99, 0x1d, 0xe8, 0x91, 0x2c, 0xd8, 0xc8, 0x8e, 0xb8, 0x94, 0x67, 0x4c,
	0x2b, 0x79, 0xe6, 0x93, 0x76, 0xaf, 0xf9, 0xf5, 0x74, 0x37, 0xdd, 0x4d, 0xf7, 0xde, 0x62, 0xda,
	0x0d, 0xc0, 0x8c, 0x85, 0x05, 0xbf, 0x6e, 0x47, 0x06, 0xc3, 0x2e, 0x8d, 0x2e, 0xc1, 0x48, 0x5f,
	0x97, 0xb9, 0x5f, 0x12, 0x2e, 0xcd, 0x6a, 0xde, 0x49, 0xf4, 0x49, 0xdf, 0x21, 0xeb, 0x28, 0xad,
	0xd4, 0xc6, 0xe1, 0x70, 0xd9, 0xdc, 0xaf, 0xa3, 0x2e, 0x4a, 0x2f, 0x1e, 0x6b, 0xe3, 0xa6, 0xa2,
	0x7c, 0x27, 0x59, 0xf3, 0xff, 0xfc, 0x6a, 0xe3, 0x87, 0x2b, 0x64, 0xfd, 0xe1, 0xc9, 0x24, 0xc0,
	0x9f, 0xbd, 0x38, 0xc0, 0x11, 0x12, 0x69, 0x3a, 0x5c, 0x0c, 0xf0, 0xe1, 0xc9, 0xc1, 0xff, 0x77,
	0x7c, 0xf4, 0xa7, 0x09, 0xd9, 0x14, 0xca, 0x81, 0x51, 0x5c, 0x76, 0x72, 0x51, 0x98, 0xfa, 0x32,
	0x76, 0x63, 0x83, 0xb6, 0x92, 0x3e, 0x6d, 0x1f, 0xdc, 0xcb, 0x58, 0x0c, 0x17, 0x70, 0x76, 0x4c,
	0xe0, 0x4c, 0x81, 0x1b, 0x6b, 0xf3, 0xcc, 0xb2, 0x6d, 0x48, 0xfb, 0x29, 0xdb, 0xbb, 0xdd, 0x4c,
	0xf7, 0x6e, 0xbd, 0xeb, 0xc3, 0xd8, 0xd9, 0xbb, 0x75, 0x23, 0x65, 0x4f, 0x80, 0x61, 0x52, 0xfa,
	0xb0, 0x7d, 0x3a, 0x0e, 0xf4, 0x18, 0x93, 0x77, 0x42, 0xb0, 0x03, 0xcf, 0x23, 0x93, 0xc8, 0x7d,
	0xc3, 0xd9, 0x98, 0x6c, 0xb4, 0x45, 0x61, 0xe8, 0x8f, 0x12, 0xb2, 0x76, 0x5a, 0xaa, 0x20, 0x6a,
	0x05, 0x45, 0x0d, 0x51, 0x54, 0x9f, 0xc2, 0x67, 0x45, 0x9d, 0x96, 0xea, 0x73, 0xd7, 0x73, 0xfe,
	0xb4, 0x54, 0x5e, 0x4a, 0xe3, 0x37, 0x09, 0x59, 0x6b, 0xb7, 0x62, 0x0e, 0xfc, 0x24, 0x99, 0x9f,
	0x88, 0x0b, 0x29, 0xe0, 0x8b, 0xb9, 0x3a, 0x08, 0x75, 0x0f, 0x7d, 0xb4, 0x5b, 0x2f, 0x98, 0x2d,
	0xc2, 0x86, 0xa6, 0x9c, 0x6b, 0x63, 0xc0, 0x96, 0x5a, 0x4d, 0x07, 0x09, 0xcc, 0xf5, 0x02, 0xa1,
	0x62, 0xa1, 0xa7, 0xf3, 0x23, 0x6d, 0x6e, 0xaa, 0x36, 0x7e, 0xb9, 0x42, 0xb6, 0x1e, 0x18, 0xad,
	0x1c, 0xa8, 0x22, 0x6a, 0xfd, 0xfe, 0x8b, 0xf3, 0xf5, 0xbb, 0x28, 0xf6, 0x09, 0x3d, 0x59, 0xcc,
	0xd7, 0x5e, 0x34, 0xaf, 0x24, 0xed, 0xa4, 0xc3, 0x54, 0xfa, 0x0a, 0x1e, 0xe0, 0x04, 0x81, 0xc9,
	0xa9, 0xa5, 0x64, 0x3d, 0x6d, 0xd8, 0x53, 0xdd, 0xb5, 0x9f, 0x4f, 0x5e, 0x1e, 0x91, 0x5a, 0x75,
	0x76, 0x2f, 0xcf, 0xbf, 0x06, 0xbf, 0xfd, 0xce, 0xee, 0x6d, 0x56, 0x19, 0xb1, 0xfe, 0x94, 0xa7,
	0xda, 0x63, 0x6f, 0xcc, 0xaa, 0x04, 0xb4, 0x37, 0x7f, 0x69, 0xe1, 0x4d, 0x77, 0x1f, 0xf9, 0xde,
	0xa3, 0x77, 0xff, 0xc3, 0xa5, 0x2d, 0x50, 0x2e, 0xdc, 0xa0, 0x8f, 0x78, 0xc8, 0xdd, 0xfc, 0xd3,
	0xe6, 0x63, 0x72, 0x29, 0x4c, 0x87, 0x8e, 0x04, 0x6e, 0x61, 0xf6, 0xf8, 0xdb, 0xdc, 0x7f, 0x13,
	0xbd, 0xbd, 0x41, 0x1b, 0xdf, 0xd4, 0xe3, 0xf0, 0x84, 0x75, 0x9a, 0x21, 0x6a, 0x32, 0x50, 0x86,
	0x60, 0x2d, 0xef, 0x83, 0x4d, 0xb3, 0x0b, 0x61, 0xe5, 0xd0, 0xef, 0xe2, 0x53, 0xf0, 0x7d, 0xb2,
	0x56, 0x28, 0xdb, 0xc1, 0x17, 0xf5, 0x2a, 0x8a, 0xff, 0x0a, 0xd2, 0x5d, 0xa3, 0x57, 0xbd, 0xf8,
	0x7b, 0x47, 0x8f, 0xf0, 0x31, 0xbd, 0xa8, 0x3c, 0xcd, 0xce, 0x17, 0xca, 0x1e, 0xf1, 0x21, 0x34,
	0xfe, 0xb8, 0x44, 0x2e, 0xdc, 0xe3, 0x8e, 0x5b, 0xa7, 0x0d, 0xc4, 0xfc, 0xb0, 0x64, 0x4b, 0x0c,
	0x4b, 0x09, 0x43, 0x50, 0xa1, 0xe5, 0xc7, 0x04, 0xf9, 0x00, 0xb9, 0xef, 0xd3, 0x5d, 0xcf, 0x5d,
	0xe5, 0x2d, 0x26, 0x04, 0x6c, 0xde, 0xce, 0x87, 0x32, 0xb2, 0x90, 0x36, 0x2f, 0x3d, 0x10, 0x12,
	0xf6, 0xb9, 0x05, 0xef, 0xea, 0x91, 0x47, 0x66, 0x0b, 0x2e, 0xe8, 0x21, 0x59, 0xf3, 0x23, 0xba,
	0xf2, 0xca, 0x9f, 0x9f, 0xba, 0x71, 0x8f, 0x6d, 0x6b, 0x83, 0x13, 0xff, 0x46, 0xc8, 0xb9, 0x8a,
	0xf3, 0x34, 0x9b, 0x32, 0xd0, 0x4f, 0xc8, 0xe5, 0x9e, 0x90, 0x80, 0xeb, 0x9d, 0x42, 0x18, 0xc8,
	0x9d, 0x36, 0x67, 0x31, 0x61, 0x6e, 0x20, 0xf1, 0xeb, 0xf4, 0x35, 0xbc, 0x60, 0xff, 0x7a, 0xf0,
	0x75, 0x8f, 0xfa, 0x47, 0xa5, 0xd4, 0xdc, 0x5f, 0x26, 0x5a, 0xa7, 0x19, 0x9d, 0xb2, 0xdc, 0x9b,
	0x90, 0x34, 0x7e, 0x9f, 0x90, 0xda, 0x03, 0xff, 0x84, 0x9e, 0xb5, 0xff, 0x2f, 0x63, 0xb3, 0x30,
	0xba, 0xef, 0x33, 0x89, 0x17, 0x85, 0xf0, 0x22, 0xb8, 0xec, 0x7c, 0x3a, 0x02, 0x23, 0x20, 0x14,
	0x58, 0xad, 0xf9, 0x4a, 0xf8, 0x28, 0x4a, 0x3f, 0xfe, 0xe8, 0xb0, 0xad, 0xa5, 0x44, 0xba, 0x96,
	0xe9, 0xdb, 0xfd, 0xdb, 0x28, 0xe8, 0x6d, 0xba, 0xd7, 0x9a, 0x9a, 0xb2, 0x68, 0xea, 0xa5, 0x09,
	0x6b, 0x47, 0xd8, 0xee, 0x0f, 0x66, 0x5e, 0x58, 0xcf, 0x2b, 0x48, 0xb3, 0x57, 0x2b, 0x9e, 0x67,
	0xd6, 0x1f, 0x05, 0xe3, 0xc6, 0xbf, 0x96, 0xc8, 0x6a, 0x54, 0xf9, 0x2d, 0xb2, 0x1a, 0x8a, 0x36,
	0x8a, 0xb9, 0x1c, 0xc5, 0x54, 0xbf, 0xce, 0xf6, 0xaf, 0xa2, 0x90, 0x57, 0xe8, 0x4b, 0x61, 0x75,
	0xf1, 0xed, 0x17, 0x19, 0xe8, 0x31, 0x59, 0x6e, 0x1d, 0x1f, 0xe0, 0x35, 0xd5, 0x9a, 0x17, 0x23,
	0xd1, 0x74, 0xe0, 0xcf, 0x0a, 0xb2, 0x5d, 0x35, 0xc7, 0x56, 0x80, 0xb3, 0xbd, 0x75, 0x5c, 0x99,
	0x7f, 0x99, 0xa7, 0xa2, 0x0d, 0xb2, 0xfc, 0xf0, 0xe4, 0x00, 0xef, 0x67, 0xc6, 0x38, 0x9d, 0xb0,
	0x99, 0xdf, 0xa4, 0xd7, 0xc8, 0x52, 0xbb, 0x85, 0x35, 0x5a, 0x6b, 0x5e, 0x98, 0xa8, 0x8f, 0xfd,
	0x37, 0x5b, 0x6a, 0xb7, 0xe8, 0x1e, 0x59, 0x9b, 0x74, 0x3a, 0x2c, 0xae, 0x5a, 0xf3, 0xa5, 0x08,
	0x9b, 0x6f, 0x80, 0xd9, 0x14, 0x46, 0x6f, 0x92, 0xf5, 0x69, 0xf6, 0x63, 0x05, 0xd5, 0x9a, 0x2f,
	0x47, 0x9b, 0x85, 0xaa, 0xc8, 0x66, 0x40, 0xba, 0x4d, 0xce, 0x61, 0x02, 0xe0, 0x77, 0x4f, 0xad,
	0x49, 0x27, 0x5e, 0x66, 0x49, 0x91, 0x05, 0x40, 0x77, 0x15, 0x77, 0xde, 0xfe, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0x5f, 0xe7, 0x26, 0x98, 0x0f, 0x00, 0x00,
}

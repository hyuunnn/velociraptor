// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vfs.proto

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

type VFSListRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	VfsPath              string   `protobuf:"bytes,2,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	RecursionDepth       uint64   `protobuf:"varint,3,opt,name=recursion_depth,json=recursionDepth,proto3" json:"recursion_depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VFSListRequest) Reset()         { *m = VFSListRequest{} }
func (m *VFSListRequest) String() string { return proto.CompactTextString(m) }
func (*VFSListRequest) ProtoMessage()    {}
func (*VFSListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vfs_39ad82f312d4da57, []int{0}
}
func (m *VFSListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VFSListRequest.Unmarshal(m, b)
}
func (m *VFSListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VFSListRequest.Marshal(b, m, deterministic)
}
func (dst *VFSListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VFSListRequest.Merge(dst, src)
}
func (m *VFSListRequest) XXX_Size() int {
	return xxx_messageInfo_VFSListRequest.Size(m)
}
func (m *VFSListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VFSListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VFSListRequest proto.InternalMessageInfo

func (m *VFSListRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VFSListRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *VFSListRequest) GetRecursionDepth() uint64 {
	if m != nil {
		return m.RecursionDepth
	}
	return 0
}

type VFSListRequestState struct {
	VfsPath              string              `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Current              *proto1.VQLResponse `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VFSListRequestState) Reset()         { *m = VFSListRequestState{} }
func (m *VFSListRequestState) String() string { return proto.CompactTextString(m) }
func (*VFSListRequestState) ProtoMessage()    {}
func (*VFSListRequestState) Descriptor() ([]byte, []int) {
	return fileDescriptor_vfs_39ad82f312d4da57, []int{1}
}
func (m *VFSListRequestState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VFSListRequestState.Unmarshal(m, b)
}
func (m *VFSListRequestState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VFSListRequestState.Marshal(b, m, deterministic)
}
func (dst *VFSListRequestState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VFSListRequestState.Merge(dst, src)
}
func (m *VFSListRequestState) XXX_Size() int {
	return xxx_messageInfo_VFSListRequestState.Size(m)
}
func (m *VFSListRequestState) XXX_DiscardUnknown() {
	xxx_messageInfo_VFSListRequestState.DiscardUnknown(m)
}

var xxx_messageInfo_VFSListRequestState proto.InternalMessageInfo

func (m *VFSListRequestState) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *VFSListRequestState) GetCurrent() *proto1.VQLResponse {
	if m != nil {
		return m.Current
	}
	return nil
}

func init() {
	proto.RegisterType((*VFSListRequest)(nil), "proto.VFSListRequest")
	proto.RegisterType((*VFSListRequestState)(nil), "proto.VFSListRequestState")
}

func init() { proto.RegisterFile("vfs.proto", fileDescriptor_vfs_39ad82f312d4da57) }

var fileDescriptor_vfs_39ad82f312d4da57 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0xcb, 0x13, 0x31,
	0x10, 0xc6, 0xd9, 0xaa, 0xfd, 0x13, 0xa1, 0x42, 0xbc, 0xd4, 0x5e, 0x3a, 0x16, 0xc4, 0x22, 0xb2,
	0x5b, 0x2a, 0x5e, 0x04, 0x3d, 0x2c, 0xa5, 0x50, 0x28, 0xa2, 0xa9, 0xf6, 0xd8, 0x12, 0xb3, 0xb3,
	0xdd, 0xc0, 0x36, 0xd9, 0x26, 0xd3, 0xad, 0x7e, 0x3f, 0x3f, 0x89, 0x7e, 0x09, 0x0f, 0x1e, 0xa4,
	0x59, 0xab, 0xef, 0x7b, 0x7b, 0x4f, 0x81, 0x99, 0xe7, 0xf7, 0x63, 0x9e, 0xb0, 0x5e, 0x9d, 0xfb,
	0xb8, 0x72, 0x96, 0x2c, 0x7f, 0x10, 0x9e, 0xe1, 0x9b, 0xf3, 0xf9, 0x1c, 0xd7, 0x58, 0x5a, 0xa5,
	0x33, 0xfc, 0x1a, 0x2b, 0x7b, 0x48, 0xf6, 0xb6, 0x94, 0x66, 0x9f, 0x34, 0x43, 0x27, 0x2b, 0xb2,
	0x2e, 0x09, 0xe1, 0xc4, 0xe3, 0x41, 0x1a, 0xd2, 0xaa, 0x51, 0x0c, 0xdf, 0xde, 0x8d, 0x95, 0x8a,
	0xb4, 0x35, 0xfe, 0xaf, 0xa3, 0x3e, 0x96, 0x0d, 0x3e, 0xfe, 0x15, 0xb1, 0xfe, 0x66, 0xb1, 0x5e,
	0x69, 0x4f, 0x02, 0x8f, 0x27, 0xf4, 0xc4, 0x53, 0xd6, 0x53, 0xa5, 0x46, 0x43, 0x3b, 0x9d, 0x0d,
	0x22, 0x88, 0x26, 0xbd, 0xf4, 0xd9, 0x8f, 0xdf, 0x3f, 0xbf, 0x47, 0x23, 0xd6, 0x16, 0xf3, 0xc5,
	0x67, 0xf1, 0x9e, 0xf3, 0x4f, 0x05, 0x42, 0x13, 0x02, 0xb2, 0x50, 0x6a, 0x4f, 0xc3, 0xa8, 0x25,
	0xba, 0xcd, 0x68, 0x99, 0xf1, 0x94, 0x75, 0xeb, 0xdc, 0xef, 0x2a, 0x49, 0xc5, 0xa0, 0x15, 0x14,
	0xcf, 0x83, 0xe2, 0x29, 0x1f, 0x5d, 0xd0, 0xcd, 0x62, 0x0d, 0x1f, 0x24, 0x15, 0x57, 0x18, 0x72,
	0x5d, 0xa2, 0x07, 0x6d, 0x20, 0x16, 0x9d, 0x3a, 0xf7, 0x97, 0x1d, 0xdf, 0xb2, 0x47, 0x0e, 0xd5,
	0xc9, 0x79, 0x6d, 0xcd, 0x2e, 0xc3, 0x8a, 0x8a, 0xc1, 0x3d, 0x88, 0x26, 0xf7, 0xd3, 0xd7, 0x41,
	0x95, 0xf0, 0x17, 0xcb, 0x1c, 0xde, 0xc1, 0x14, 0xa8, 0x40, 0x03, 0x4d, 0x14, 0x81, 0x0a, 0xed,
	0xe1, 0x20, 0xcd, 0x37, 0xc8, 0xb4, 0x43, 0x45, 0xd6, 0x69, 0xf4, 0xf1, 0x2c, 0x9a, 0x8a, 0xfe,
	0x3f, 0xdb, 0xfc, 0x22, 0x1b, 0x6f, 0xd9, 0xe3, 0xdb, 0xcd, 0xd7, 0x24, 0x09, 0xf9, 0x93, 0x1b,
	0xa7, 0x87, 0xf6, 0xff, 0x2f, 0x7a, 0xc9, 0x3a, 0xea, 0xe4, 0x1c, 0x1a, 0x0a, 0xa5, 0x1e, 0xce,
	0x78, 0xf3, 0x8b, 0xf1, 0xe6, 0xe3, 0x4a, 0xa0, 0xaf, 0xac, 0xf1, 0x28, 0xae, 0x91, 0x2f, 0xed,
	0xb0, 0x7b, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0x39, 0x8a, 0x6a, 0xc1, 0xf0, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/raft/raft.proto

package raft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Metadata struct {
	BindAddr             string   `protobuf:"bytes,1,opt,name=bind_addr,json=bindAddr,proto3" json:"bind_addr,omitempty"`
	GrpcAddr             string   `protobuf:"bytes,2,opt,name=grpc_addr,json=grpcAddr,proto3" json:"grpc_addr,omitempty"`
	HttpAddr             string   `protobuf:"bytes,3,opt,name=http_addr,json=httpAddr,proto3" json:"http_addr,omitempty"`
	DataDir              string   `protobuf:"bytes,4,opt,name=data_dir,json=dataDir,proto3" json:"data_dir,omitempty"`
	Leader               bool     `protobuf:"varint,5,opt,name=leader,proto3" json:"leader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_028aa12295c796d4, []int{0}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

func (m *Metadata) GetGrpcAddr() string {
	if m != nil {
		return m.GrpcAddr
	}
	return ""
}

func (m *Metadata) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

func (m *Metadata) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

func (m *Metadata) GetLeader() bool {
	if m != nil {
		return m.Leader
	}
	return false
}

type Node struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_028aa12295c796d4, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Cluster struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nodes                map[string]*Metadata `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_028aa12295c796d4, []int{2}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cluster) GetNodes() map[string]*Metadata {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "raft.Metadata")
	proto.RegisterType((*Node)(nil), "raft.Node")
	proto.RegisterType((*Cluster)(nil), "raft.Cluster")
	proto.RegisterMapType((map[string]*Metadata)(nil), "raft.Cluster.NodesEntry")
}

func init() { proto.RegisterFile("protobuf/raft/raft.proto", fileDescriptor_028aa12295c796d4) }

var fileDescriptor_028aa12295c796d4 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0xf4, 0x5f, 0x7a, 0x0b, 0xe5, 0x63, 0x16, 0x1f, 0xa3, 0x6e, 0x42, 0x50, 0x0c,
	0x2e, 0x12, 0xa8, 0x1b, 0x71, 0x67, 0x55, 0x70, 0xa3, 0x8b, 0x2c, 0xdd, 0x94, 0x49, 0xef, 0xb4,
	0x1d, 0x9a, 0x64, 0xc2, 0x64, 0x22, 0xf4, 0x39, 0xf4, 0x81, 0x65, 0x66, 0x52, 0x45, 0xc4, 0x4d,
	0xb8, 0xe7, 0xfc, 0x2e, 0x27, 0xe7, 0x32, 0xc0, 0x1a, 0xad, 0x8c, 0x2a, 0xba, 0x4d, 0xa6, 0xf9,
	0xc6, 0xb8, 0x4f, 0xea, 0x2c, 0x3a, 0xb4, 0x73, 0xfc, 0x41, 0x20, 0x7c, 0x16, 0x86, 0x23, 0x37,
	0x9c, 0x9e, 0xc1, 0xb4, 0x90, 0x35, 0xae, 0x38, 0xa2, 0x66, 0x24, 0x22, 0xc9, 0x34, 0x0f, 0xad,
	0x71, 0x87, 0xa8, 0x2d, 0xdc, 0xea, 0x66, 0xed, 0x61, 0xe0, 0xa1, 0x35, 0x8e, 0x70, 0x67, 0x4c,
	0xe3, 0xe1, 0xc0, 0x43, 0x6b, 0x38, 0x78, 0x02, 0xa1, 0x8d, 0x5f, 0xa1, 0xd4, 0x6c, 0xe8, 0xd8,
	0xc4, 0xea, 0x07, 0xa9, 0xe9, 0x7f, 0x18, 0x97, 0x82, 0xa3, 0xd0, 0x6c, 0x14, 0x91, 0x24, 0xcc,
	0x7b, 0x15, 0x2f, 0x61, 0xf8, 0xa2, 0x50, 0xd0, 0x39, 0x04, 0x12, 0xfb, 0x2a, 0x81, 0x44, 0x7a,
	0x05, 0x61, 0xd5, 0xb7, 0x75, 0x1d, 0x66, 0x8b, 0x79, 0xea, 0x6e, 0x3a, 0xde, 0x90, 0x7f, 0xf1,
	0xf8, 0x9d, 0xc0, 0xe4, 0xbe, 0xec, 0x5a, 0x23, 0xf4, 0xaf, 0x9c, 0x14, 0x46, 0xb5, 0x42, 0xd1,
	0xb2, 0x20, 0x1a, 0x24, 0xb3, 0x05, 0xf3, 0x21, 0xfd, 0x76, 0x6a, 0x7f, 0xdd, 0x3e, 0xd6, 0x46,
	0x1f, 0x72, 0xbf, 0x76, 0xfa, 0x04, 0xf0, 0x6d, 0xd2, 0x7f, 0x30, 0xd8, 0x8b, 0x43, 0x1f, 0x67,
	0x47, 0x7a, 0x0e, 0xa3, 0x37, 0x5e, 0x76, 0xe2, 0x8f, 0x52, 0x1e, 0xde, 0x06, 0x37, 0x64, 0x79,
	0xf9, 0x7a, 0xb1, 0x95, 0x66, 0xd7, 0x15, 0xe9, 0x5a, 0x55, 0x59, 0xa5, 0xda, 0x6e, 0xcf, 0xb3,
	0xa2, 0xe4, 0xad, 0xc9, 0x7e, 0x3c, 0x55, 0x31, 0x76, 0xf2, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0x56, 0x1d, 0xb1, 0xc2, 0x01, 0x00, 0x00,
}

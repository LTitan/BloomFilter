// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MachineInfo struct {
	Cpu                  uint32   `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	CpuUsage             float32  `protobuf:"fixed32,2,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	Memory               uint64   `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	MemoryUsage          float32  `protobuf:"fixed32,4,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	Host                 string   `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineInfo) Reset()         { *m = MachineInfo{} }
func (m *MachineInfo) String() string { return proto.CompactTextString(m) }
func (*MachineInfo) ProtoMessage()    {}
func (*MachineInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *MachineInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineInfo.Unmarshal(m, b)
}
func (m *MachineInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineInfo.Marshal(b, m, deterministic)
}
func (m *MachineInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineInfo.Merge(m, src)
}
func (m *MachineInfo) XXX_Size() int {
	return xxx_messageInfo_MachineInfo.Size(m)
}
func (m *MachineInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MachineInfo proto.InternalMessageInfo

func (m *MachineInfo) GetCpu() uint32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *MachineInfo) GetCpuUsage() float32 {
	if m != nil {
		return m.CpuUsage
	}
	return 0
}

func (m *MachineInfo) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *MachineInfo) GetMemoryUsage() float32 {
	if m != nil {
		return m.MemoryUsage
	}
	return 0
}

func (m *MachineInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *MachineInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Reply struct {
	Recv                 bool     `protobuf:"varint,1,opt,name=recv,proto3" json:"recv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetRecv() bool {
	if m != nil {
		return m.Recv
	}
	return false
}

type ApplyReply struct {
	Recv                 bool     `protobuf:"varint,1,opt,name=recv,proto3" json:"recv,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyReply) Reset()         { *m = ApplyReply{} }
func (m *ApplyReply) String() string { return proto.CompactTextString(m) }
func (*ApplyReply) ProtoMessage()    {}
func (*ApplyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}

func (m *ApplyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyReply.Unmarshal(m, b)
}
func (m *ApplyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyReply.Marshal(b, m, deterministic)
}
func (m *ApplyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyReply.Merge(m, src)
}
func (m *ApplyReply) XXX_Size() int {
	return xxx_messageInfo_ApplyReply.Size(m)
}
func (m *ApplyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyReply.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyReply proto.InternalMessageInfo

func (m *ApplyReply) GetRecv() bool {
	if m != nil {
		return m.Recv
	}
	return false
}

func (m *ApplyReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ApplyRequest struct {
	Size                 uint64   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRequest) Reset()         { *m = ApplyRequest{} }
func (m *ApplyRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyRequest) ProtoMessage()    {}
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{3}
}

func (m *ApplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRequest.Unmarshal(m, b)
}
func (m *ApplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRequest.Marshal(b, m, deterministic)
}
func (m *ApplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRequest.Merge(m, src)
}
func (m *ApplyRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyRequest.Size(m)
}
func (m *ApplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRequest proto.InternalMessageInfo

func (m *ApplyRequest) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type AddRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{4}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddRequest) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type DeleteRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{5}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type QueryRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{6}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *QueryRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type QueryManyRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryManyRequest) Reset()         { *m = QueryManyRequest{} }
func (m *QueryManyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryManyRequest) ProtoMessage()    {}
func (*QueryManyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{7}
}

func (m *QueryManyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryManyRequest.Unmarshal(m, b)
}
func (m *QueryManyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryManyRequest.Marshal(b, m, deterministic)
}
func (m *QueryManyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryManyRequest.Merge(m, src)
}
func (m *QueryManyRequest) XXX_Size() int {
	return xxx_messageInfo_QueryManyRequest.Size(m)
}
func (m *QueryManyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryManyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryManyRequest proto.InternalMessageInfo

func (m *QueryManyRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *QueryManyRequest) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type QueryManyReply struct {
	Has                  bool     `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
	Result               []bool   `protobuf:"varint,2,rep,packed,name=result,proto3" json:"result,omitempty"`
	Err                  string   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryManyReply) Reset()         { *m = QueryManyReply{} }
func (m *QueryManyReply) String() string { return proto.CompactTextString(m) }
func (*QueryManyReply) ProtoMessage()    {}
func (*QueryManyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{8}
}

func (m *QueryManyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryManyReply.Unmarshal(m, b)
}
func (m *QueryManyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryManyReply.Marshal(b, m, deterministic)
}
func (m *QueryManyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryManyReply.Merge(m, src)
}
func (m *QueryManyReply) XXX_Size() int {
	return xxx_messageInfo_QueryManyReply.Size(m)
}
func (m *QueryManyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryManyReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryManyReply proto.InternalMessageInfo

func (m *QueryManyReply) GetHas() bool {
	if m != nil {
		return m.Has
	}
	return false
}

func (m *QueryManyReply) GetResult() []bool {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *QueryManyReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*MachineInfo)(nil), "rpc.MachineInfo")
	proto.RegisterType((*Reply)(nil), "rpc.Reply")
	proto.RegisterType((*ApplyReply)(nil), "rpc.ApplyReply")
	proto.RegisterType((*ApplyRequest)(nil), "rpc.ApplyRequest")
	proto.RegisterType((*AddRequest)(nil), "rpc.AddRequest")
	proto.RegisterType((*DeleteRequest)(nil), "rpc.DeleteRequest")
	proto.RegisterType((*QueryRequest)(nil), "rpc.QueryRequest")
	proto.RegisterType((*QueryManyRequest)(nil), "rpc.QueryManyRequest")
	proto.RegisterType((*QueryManyReply)(nil), "rpc.QueryManyReply")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0x4d, 0x5b, 0x9a, 0x49, 0x17, 0x8a, 0xf9, 0x50, 0xd4, 0xbd, 0x64, 0x2d, 0x0e, 0x11,
	0x88, 0x1e, 0x8a, 0x54, 0x71, 0xe0, 0x52, 0x84, 0x04, 0x48, 0xec, 0x01, 0x57, 0x9c, 0x91, 0x49,
	0x87, 0xb6, 0xc2, 0x9b, 0x18, 0xc7, 0x89, 0x14, 0xfe, 0x0c, 0x7f, 0x93, 0x23, 0xf2, 0xc4, 0x2c,
	0x69, 0xa5, 0x1e, 0xb8, 0xbd, 0x79, 0x7e, 0x6f, 0x66, 0xec, 0x97, 0x40, 0x2c, 0x77, 0x58, 0xd8,
	0x85, 0x36, 0xa5, 0x2d, 0x59, 0x68, 0x74, 0xce, 0x7f, 0x05, 0x10, 0x5f, 0xcb, 0x7c, 0x7f, 0x28,
	0xf0, 0x43, 0xf1, 0xad, 0x64, 0x33, 0x08, 0x73, 0x5d, 0x27, 0x41, 0x1a, 0x64, 0x17, 0xc2, 0x41,
	0x76, 0x09, 0x51, 0xae, 0xeb, 0x2f, 0x75, 0x25, 0x77, 0x98, 0x0c, 0xd2, 0x20, 0x1b, 0x88, 0x49,
	0xae, 0xeb, 0xcf, 0xae, 0x66, 0x4f, 0x60, 0x7c, 0x83, 0x37, 0xa5, 0x69, 0x93, 0x30, 0x0d, 0xb2,
	0xa1, 0xf0, 0x15, 0xbb, 0x82, 0x69, 0x87, 0xbc, 0x6f, 0x48, 0xbe, 0xb8, 0xe3, 0x3a, 0x2b, 0x83,
	0xe1, 0xbe, 0xac, 0x6c, 0x32, 0x4a, 0x83, 0x2c, 0x12, 0x84, 0x1d, 0xa7, 0x4b, 0x63, 0x93, 0x31,
	0x8d, 0x27, 0xcc, 0x2f, 0x61, 0x24, 0x50, 0xab, 0xd6, 0x1d, 0x1a, 0xcc, 0x1b, 0xda, 0x6d, 0x22,
	0x08, 0xf3, 0x25, 0xc0, 0x5a, 0x6b, 0xd5, 0x9e, 0x55, 0xb8, 0x0b, 0x7d, 0xc7, 0x96, 0x16, 0x8f,
	0x84, 0x83, 0x9c, 0xc3, 0xd4, 0x7b, 0x7e, 0xd4, 0xd8, 0x0d, 0xad, 0x0e, 0x3f, 0x91, 0x5c, 0x43,
	0x41, 0x98, 0xaf, 0x00, 0xd6, 0xdb, 0xed, 0x5f, 0x85, 0xef, 0x11, 0xdc, 0xf6, 0x70, 0xf7, 0x6e,
	0xa4, 0xaa, 0xb1, 0x4a, 0x06, 0x69, 0x98, 0x45, 0xc2, 0x57, 0xfc, 0x0a, 0x2e, 0xde, 0xa2, 0x42,
	0x8b, 0x67, 0xad, 0x7c, 0x05, 0xd3, 0x4f, 0x35, 0x9a, 0xf6, 0x7c, 0xf3, 0x47, 0x30, 0xa2, 0x76,
	0x7e, 0xe9, 0xae, 0xe0, 0xaf, 0x61, 0x46, 0xbe, 0x6b, 0x59, 0xb4, 0xff, 0xbf, 0xd8, 0x47, 0xb8,
	0xd7, 0x73, 0xbb, 0xc7, 0x9a, 0x41, 0xb8, 0x97, 0x95, 0x7f, 0x2b, 0x07, 0x9d, 0xd7, 0x60, 0x55,
	0x2b, 0x4b, 0xde, 0x89, 0xf0, 0x95, 0x53, 0xa2, 0x31, 0x94, 0x70, 0x24, 0x1c, 0x5c, 0xae, 0xe0,
	0xee, 0x3b, 0x83, 0x68, 0xd1, 0xb0, 0xe7, 0x10, 0xbd, 0x47, 0x69, 0xec, 0x1b, 0x94, 0x96, 0xcd,
	0x16, 0x46, 0xe7, 0x8b, 0xde, 0xf7, 0x34, 0x07, 0x62, 0x68, 0x22, 0xbf, 0xb3, 0xfc, 0x1d, 0x40,
	0xbc, 0x51, 0xb2, 0xc1, 0x0d, 0x9a, 0x06, 0x0d, 0x7b, 0x01, 0x23, 0x8a, 0x82, 0x3d, 0x20, 0x59,
	0x3f, 0x96, 0xf9, 0xfd, 0x3e, 0x45, 0x76, 0xf6, 0x14, 0xc2, 0xf5, 0x76, 0xcb, 0xfc, 0xc9, 0x6d,
	0x3e, 0xc7, 0x43, 0xd8, 0x33, 0x18, 0x77, 0x19, 0x30, 0x46, 0xfc, 0x51, 0x20, 0x27, 0xda, 0x05,
	0xc4, 0xf4, 0x2c, 0x9b, 0x43, 0xb1, 0x53, 0xe8, 0xd7, 0xe8, 0xc7, 0x73, 0xa2, 0x7f, 0x05, 0x13,
	0x3a, 0x5d, 0x2b, 0xc5, 0x1e, 0xff, 0x13, 0xf7, 0x32, 0x99, 0x3f, 0x3c, 0xa5, 0xc9, 0xf9, 0x75,
	0x4c, 0x3f, 0xdd, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xfc, 0x7b, 0xff, 0x83, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	HeartBeat(ctx context.Context, in *MachineInfo, opts ...grpc.CallOption) (*Reply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) HeartBeat(ctx context.Context, in *MachineInfo, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	HeartBeat(context.Context, *MachineInfo) (*Reply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) HeartBeat(ctx context.Context, req *MachineInfo) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).HeartBeat(ctx, req.(*MachineInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _Greeter_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

// SlaveServerClient is the client API for SlaveServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SlaveServerClient interface {
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyReply, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*Reply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Reply, error)
	QuerySingle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*Reply, error)
	QueryAll(ctx context.Context, in *QueryManyRequest, opts ...grpc.CallOption) (*QueryManyReply, error)
}

type slaveServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSlaveServerClient(cc grpc.ClientConnInterface) SlaveServerClient {
	return &slaveServerClient{cc}
}

func (c *slaveServerClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyReply, error) {
	out := new(ApplyReply)
	err := c.cc.Invoke(ctx, "/rpc.SlaveServer/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveServerClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/rpc.SlaveServer/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveServerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/rpc.SlaveServer/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveServerClient) QuerySingle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/rpc.SlaveServer/QuerySingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveServerClient) QueryAll(ctx context.Context, in *QueryManyRequest, opts ...grpc.CallOption) (*QueryManyReply, error) {
	out := new(QueryManyReply)
	err := c.cc.Invoke(ctx, "/rpc.SlaveServer/QueryAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlaveServerServer is the server API for SlaveServer service.
type SlaveServerServer interface {
	Apply(context.Context, *ApplyRequest) (*ApplyReply, error)
	Add(context.Context, *AddRequest) (*Reply, error)
	Delete(context.Context, *DeleteRequest) (*Reply, error)
	QuerySingle(context.Context, *QueryRequest) (*Reply, error)
	QueryAll(context.Context, *QueryManyRequest) (*QueryManyReply, error)
}

// UnimplementedSlaveServerServer can be embedded to have forward compatible implementations.
type UnimplementedSlaveServerServer struct {
}

func (*UnimplementedSlaveServerServer) Apply(ctx context.Context, req *ApplyRequest) (*ApplyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (*UnimplementedSlaveServerServer) Add(ctx context.Context, req *AddRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedSlaveServerServer) Delete(ctx context.Context, req *DeleteRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedSlaveServerServer) QuerySingle(ctx context.Context, req *QueryRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySingle not implemented")
}
func (*UnimplementedSlaveServerServer) QueryAll(ctx context.Context, req *QueryManyRequest) (*QueryManyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAll not implemented")
}

func RegisterSlaveServerServer(s *grpc.Server, srv SlaveServerServer) {
	s.RegisterService(&_SlaveServer_serviceDesc, srv)
}

func _SlaveServer_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveServerServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.SlaveServer/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveServerServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveServer_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveServerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.SlaveServer/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveServerServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveServer_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveServerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.SlaveServer/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveServerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveServer_QuerySingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveServerServer).QuerySingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.SlaveServer/QuerySingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveServerServer).QuerySingle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveServer_QueryAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveServerServer).QueryAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.SlaveServer/QueryAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveServerServer).QueryAll(ctx, req.(*QueryManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SlaveServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.SlaveServer",
	HandlerType: (*SlaveServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _SlaveServer_Apply_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _SlaveServer_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SlaveServer_Delete_Handler,
		},
		{
			MethodName: "QuerySingle",
			Handler:    _SlaveServer_QuerySingle_Handler,
		},
		{
			MethodName: "QueryAll",
			Handler:    _SlaveServer_QueryAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

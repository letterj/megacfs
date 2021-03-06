// Code generated by protoc-gen-go.
// source: group_api.proto
// DO NOT EDIT!

/*
Package groupproto is a generated protocol buffer package.

It is generated from these files:
	group_api.proto

It has these top-level messages:
	DeleteRequest
	DeleteResponse
	LookupGroupRequest
	LookupGroupResponse
	LookupGroupItem
	LookupRequest
	LookupResponse
	ReadGroupRequest
	ReadGroupResponse
	ReadGroupItem
	ReadRequest
	ReadResponse
	WriteRequest
	WriteResponse
*/
package groupproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type DeleteRequest struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	KeyA           uint64 `protobuf:"varint,2,opt,name=KeyA" json:"KeyA,omitempty"`
	KeyB           uint64 `protobuf:"varint,3,opt,name=KeyB" json:"KeyB,omitempty"`
	ChildKeyA      uint64 `protobuf:"varint,4,opt,name=ChildKeyA" json:"ChildKeyA,omitempty"`
	ChildKeyB      uint64 `protobuf:"varint,5,opt,name=ChildKeyB" json:"ChildKeyB,omitempty"`
	TimestampMicro int64  `protobuf:"varint,6,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteRequest) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *DeleteRequest) GetKeyA() uint64 {
	if m != nil {
		return m.KeyA
	}
	return 0
}

func (m *DeleteRequest) GetKeyB() uint64 {
	if m != nil {
		return m.KeyB
	}
	return 0
}

func (m *DeleteRequest) GetChildKeyA() uint64 {
	if m != nil {
		return m.ChildKeyA
	}
	return 0
}

func (m *DeleteRequest) GetChildKeyB() uint64 {
	if m != nil {
		return m.ChildKeyB
	}
	return 0
}

func (m *DeleteRequest) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

type DeleteResponse struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	TimestampMicro int64  `protobuf:"varint,2,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
	Err            string `protobuf:"bytes,3,opt,name=Err" json:"Err,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeleteResponse) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *DeleteResponse) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

func (m *DeleteResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type LookupGroupRequest struct {
	RPCID uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	KeyA  uint64 `protobuf:"varint,2,opt,name=KeyA" json:"KeyA,omitempty"`
	KeyB  uint64 `protobuf:"varint,3,opt,name=KeyB" json:"KeyB,omitempty"`
}

func (m *LookupGroupRequest) Reset()                    { *m = LookupGroupRequest{} }
func (m *LookupGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupGroupRequest) ProtoMessage()               {}
func (*LookupGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LookupGroupRequest) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *LookupGroupRequest) GetKeyA() uint64 {
	if m != nil {
		return m.KeyA
	}
	return 0
}

func (m *LookupGroupRequest) GetKeyB() uint64 {
	if m != nil {
		return m.KeyB
	}
	return 0
}

type LookupGroupResponse struct {
	RPCID uint32             `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	Items []*LookupGroupItem `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
	Err   string             `protobuf:"bytes,3,opt,name=Err" json:"Err,omitempty"`
}

func (m *LookupGroupResponse) Reset()                    { *m = LookupGroupResponse{} }
func (m *LookupGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupGroupResponse) ProtoMessage()               {}
func (*LookupGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LookupGroupResponse) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *LookupGroupResponse) GetItems() []*LookupGroupItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *LookupGroupResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type LookupGroupItem struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	ChildKeyA      uint64 `protobuf:"varint,2,opt,name=ChildKeyA" json:"ChildKeyA,omitempty"`
	ChildKeyB      uint64 `protobuf:"varint,3,opt,name=ChildKeyB" json:"ChildKeyB,omitempty"`
	TimestampMicro int64  `protobuf:"varint,4,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
	Length         uint32 `protobuf:"varint,5,opt,name=Length" json:"Length,omitempty"`
}

func (m *LookupGroupItem) Reset()                    { *m = LookupGroupItem{} }
func (m *LookupGroupItem) String() string            { return proto.CompactTextString(m) }
func (*LookupGroupItem) ProtoMessage()               {}
func (*LookupGroupItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LookupGroupItem) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *LookupGroupItem) GetChildKeyA() uint64 {
	if m != nil {
		return m.ChildKeyA
	}
	return 0
}

func (m *LookupGroupItem) GetChildKeyB() uint64 {
	if m != nil {
		return m.ChildKeyB
	}
	return 0
}

func (m *LookupGroupItem) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

func (m *LookupGroupItem) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type LookupRequest struct {
	RPCID     uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	KeyA      uint64 `protobuf:"varint,2,opt,name=KeyA" json:"KeyA,omitempty"`
	KeyB      uint64 `protobuf:"varint,3,opt,name=KeyB" json:"KeyB,omitempty"`
	ChildKeyA uint64 `protobuf:"varint,4,opt,name=ChildKeyA" json:"ChildKeyA,omitempty"`
	ChildKeyB uint64 `protobuf:"varint,5,opt,name=ChildKeyB" json:"ChildKeyB,omitempty"`
}

func (m *LookupRequest) Reset()                    { *m = LookupRequest{} }
func (m *LookupRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()               {}
func (*LookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LookupRequest) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *LookupRequest) GetKeyA() uint64 {
	if m != nil {
		return m.KeyA
	}
	return 0
}

func (m *LookupRequest) GetKeyB() uint64 {
	if m != nil {
		return m.KeyB
	}
	return 0
}

func (m *LookupRequest) GetChildKeyA() uint64 {
	if m != nil {
		return m.ChildKeyA
	}
	return 0
}

func (m *LookupRequest) GetChildKeyB() uint64 {
	if m != nil {
		return m.ChildKeyB
	}
	return 0
}

type LookupResponse struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	TimestampMicro int64  `protobuf:"varint,2,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
	Length         uint32 `protobuf:"varint,3,opt,name=Length" json:"Length,omitempty"`
	Err            string `protobuf:"bytes,4,opt,name=Err" json:"Err,omitempty"`
}

func (m *LookupResponse) Reset()                    { *m = LookupResponse{} }
func (m *LookupResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()               {}
func (*LookupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LookupResponse) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *LookupResponse) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

func (m *LookupResponse) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *LookupResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ReadGroupRequest struct {
	RPCID uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	KeyA  uint64 `protobuf:"varint,2,opt,name=KeyA" json:"KeyA,omitempty"`
	KeyB  uint64 `protobuf:"varint,3,opt,name=KeyB" json:"KeyB,omitempty"`
}

func (m *ReadGroupRequest) Reset()                    { *m = ReadGroupRequest{} }
func (m *ReadGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadGroupRequest) ProtoMessage()               {}
func (*ReadGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadGroupRequest) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *ReadGroupRequest) GetKeyA() uint64 {
	if m != nil {
		return m.KeyA
	}
	return 0
}

func (m *ReadGroupRequest) GetKeyB() uint64 {
	if m != nil {
		return m.KeyB
	}
	return 0
}

type ReadGroupResponse struct {
	RPCID uint32           `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	Items []*ReadGroupItem `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
	Err   string           `protobuf:"bytes,3,opt,name=Err" json:"Err,omitempty"`
}

func (m *ReadGroupResponse) Reset()                    { *m = ReadGroupResponse{} }
func (m *ReadGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadGroupResponse) ProtoMessage()               {}
func (*ReadGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReadGroupResponse) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *ReadGroupResponse) GetItems() []*ReadGroupItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ReadGroupResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ReadGroupItem struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	ChildKeyA      uint64 `protobuf:"varint,2,opt,name=ChildKeyA" json:"ChildKeyA,omitempty"`
	ChildKeyB      uint64 `protobuf:"varint,3,opt,name=ChildKeyB" json:"ChildKeyB,omitempty"`
	TimestampMicro int64  `protobuf:"varint,4,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
	Value          []byte `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *ReadGroupItem) Reset()                    { *m = ReadGroupItem{} }
func (m *ReadGroupItem) String() string            { return proto.CompactTextString(m) }
func (*ReadGroupItem) ProtoMessage()               {}
func (*ReadGroupItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReadGroupItem) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *ReadGroupItem) GetChildKeyA() uint64 {
	if m != nil {
		return m.ChildKeyA
	}
	return 0
}

func (m *ReadGroupItem) GetChildKeyB() uint64 {
	if m != nil {
		return m.ChildKeyB
	}
	return 0
}

func (m *ReadGroupItem) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

func (m *ReadGroupItem) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ReadRequest struct {
	RPCID     uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	KeyA      uint64 `protobuf:"varint,2,opt,name=KeyA" json:"KeyA,omitempty"`
	KeyB      uint64 `protobuf:"varint,3,opt,name=KeyB" json:"KeyB,omitempty"`
	ChildKeyA uint64 `protobuf:"varint,4,opt,name=ChildKeyA" json:"ChildKeyA,omitempty"`
	ChildKeyB uint64 `protobuf:"varint,5,opt,name=ChildKeyB" json:"ChildKeyB,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReadRequest) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *ReadRequest) GetKeyA() uint64 {
	if m != nil {
		return m.KeyA
	}
	return 0
}

func (m *ReadRequest) GetKeyB() uint64 {
	if m != nil {
		return m.KeyB
	}
	return 0
}

func (m *ReadRequest) GetChildKeyA() uint64 {
	if m != nil {
		return m.ChildKeyA
	}
	return 0
}

func (m *ReadRequest) GetChildKeyB() uint64 {
	if m != nil {
		return m.ChildKeyB
	}
	return 0
}

type ReadResponse struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	TimestampMicro int64  `protobuf:"varint,2,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
	Value          []byte `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
	Err            string `protobuf:"bytes,4,opt,name=Err" json:"Err,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReadResponse) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *ReadResponse) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

func (m *ReadResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ReadResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type WriteRequest struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	KeyA           uint64 `protobuf:"varint,2,opt,name=KeyA" json:"KeyA,omitempty"`
	KeyB           uint64 `protobuf:"varint,3,opt,name=KeyB" json:"KeyB,omitempty"`
	ChildKeyA      uint64 `protobuf:"varint,4,opt,name=ChildKeyA" json:"ChildKeyA,omitempty"`
	ChildKeyB      uint64 `protobuf:"varint,5,opt,name=ChildKeyB" json:"ChildKeyB,omitempty"`
	Value          []byte `protobuf:"bytes,6,opt,name=Value,proto3" json:"Value,omitempty"`
	TimestampMicro int64  `protobuf:"varint,7,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *WriteRequest) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *WriteRequest) GetKeyA() uint64 {
	if m != nil {
		return m.KeyA
	}
	return 0
}

func (m *WriteRequest) GetKeyB() uint64 {
	if m != nil {
		return m.KeyB
	}
	return 0
}

func (m *WriteRequest) GetChildKeyA() uint64 {
	if m != nil {
		return m.ChildKeyA
	}
	return 0
}

func (m *WriteRequest) GetChildKeyB() uint64 {
	if m != nil {
		return m.ChildKeyB
	}
	return 0
}

func (m *WriteRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *WriteRequest) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

type WriteResponse struct {
	RPCID          uint32 `protobuf:"varint,1,opt,name=RPCID" json:"RPCID,omitempty"`
	TimestampMicro int64  `protobuf:"varint,2,opt,name=TimestampMicro" json:"TimestampMicro,omitempty"`
	Err            string `protobuf:"bytes,3,opt,name=Err" json:"Err,omitempty"`
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *WriteResponse) GetRPCID() uint32 {
	if m != nil {
		return m.RPCID
	}
	return 0
}

func (m *WriteResponse) GetTimestampMicro() int64 {
	if m != nil {
		return m.TimestampMicro
	}
	return 0
}

func (m *WriteResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteRequest)(nil), "groupproto.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "groupproto.DeleteResponse")
	proto.RegisterType((*LookupGroupRequest)(nil), "groupproto.LookupGroupRequest")
	proto.RegisterType((*LookupGroupResponse)(nil), "groupproto.LookupGroupResponse")
	proto.RegisterType((*LookupGroupItem)(nil), "groupproto.LookupGroupItem")
	proto.RegisterType((*LookupRequest)(nil), "groupproto.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "groupproto.LookupResponse")
	proto.RegisterType((*ReadGroupRequest)(nil), "groupproto.ReadGroupRequest")
	proto.RegisterType((*ReadGroupResponse)(nil), "groupproto.ReadGroupResponse")
	proto.RegisterType((*ReadGroupItem)(nil), "groupproto.ReadGroupItem")
	proto.RegisterType((*ReadRequest)(nil), "groupproto.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "groupproto.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "groupproto.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "groupproto.WriteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GroupStore service

type GroupStoreClient interface {
	Delete(ctx context.Context, opts ...grpc.CallOption) (GroupStore_DeleteClient, error)
	LookupGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_LookupGroupClient, error)
	Lookup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_LookupClient, error)
	ReadGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_ReadGroupClient, error)
	Read(ctx context.Context, opts ...grpc.CallOption) (GroupStore_ReadClient, error)
	Write(ctx context.Context, opts ...grpc.CallOption) (GroupStore_WriteClient, error)
}

type groupStoreClient struct {
	cc *grpc.ClientConn
}

func NewGroupStoreClient(cc *grpc.ClientConn) GroupStoreClient {
	return &groupStoreClient{cc}
}

func (c *groupStoreClient) Delete(ctx context.Context, opts ...grpc.CallOption) (GroupStore_DeleteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[0], c.cc, "/groupproto.GroupStore/Delete", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreDeleteClient{stream}
	return x, nil
}

type GroupStore_DeleteClient interface {
	Send(*DeleteRequest) error
	Recv() (*DeleteResponse, error)
	grpc.ClientStream
}

type groupStoreDeleteClient struct {
	grpc.ClientStream
}

func (x *groupStoreDeleteClient) Send(m *DeleteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreDeleteClient) Recv() (*DeleteResponse, error) {
	m := new(DeleteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) LookupGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_LookupGroupClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[1], c.cc, "/groupproto.GroupStore/LookupGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreLookupGroupClient{stream}
	return x, nil
}

type GroupStore_LookupGroupClient interface {
	Send(*LookupGroupRequest) error
	Recv() (*LookupGroupResponse, error)
	grpc.ClientStream
}

type groupStoreLookupGroupClient struct {
	grpc.ClientStream
}

func (x *groupStoreLookupGroupClient) Send(m *LookupGroupRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreLookupGroupClient) Recv() (*LookupGroupResponse, error) {
	m := new(LookupGroupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) Lookup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_LookupClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[2], c.cc, "/groupproto.GroupStore/Lookup", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreLookupClient{stream}
	return x, nil
}

type GroupStore_LookupClient interface {
	Send(*LookupRequest) error
	Recv() (*LookupResponse, error)
	grpc.ClientStream
}

type groupStoreLookupClient struct {
	grpc.ClientStream
}

func (x *groupStoreLookupClient) Send(m *LookupRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreLookupClient) Recv() (*LookupResponse, error) {
	m := new(LookupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) ReadGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_ReadGroupClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[3], c.cc, "/groupproto.GroupStore/ReadGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreReadGroupClient{stream}
	return x, nil
}

type GroupStore_ReadGroupClient interface {
	Send(*ReadGroupRequest) error
	Recv() (*ReadGroupResponse, error)
	grpc.ClientStream
}

type groupStoreReadGroupClient struct {
	grpc.ClientStream
}

func (x *groupStoreReadGroupClient) Send(m *ReadGroupRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreReadGroupClient) Recv() (*ReadGroupResponse, error) {
	m := new(ReadGroupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) Read(ctx context.Context, opts ...grpc.CallOption) (GroupStore_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[4], c.cc, "/groupproto.GroupStore/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreReadClient{stream}
	return x, nil
}

type GroupStore_ReadClient interface {
	Send(*ReadRequest) error
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type groupStoreReadClient struct {
	grpc.ClientStream
}

func (x *groupStoreReadClient) Send(m *ReadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) Write(ctx context.Context, opts ...grpc.CallOption) (GroupStore_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[5], c.cc, "/groupproto.GroupStore/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreWriteClient{stream}
	return x, nil
}

type GroupStore_WriteClient interface {
	Send(*WriteRequest) error
	Recv() (*WriteResponse, error)
	grpc.ClientStream
}

type groupStoreWriteClient struct {
	grpc.ClientStream
}

func (x *groupStoreWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreWriteClient) Recv() (*WriteResponse, error) {
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GroupStore service

type GroupStoreServer interface {
	Delete(GroupStore_DeleteServer) error
	LookupGroup(GroupStore_LookupGroupServer) error
	Lookup(GroupStore_LookupServer) error
	ReadGroup(GroupStore_ReadGroupServer) error
	Read(GroupStore_ReadServer) error
	Write(GroupStore_WriteServer) error
}

func RegisterGroupStoreServer(s *grpc.Server, srv GroupStoreServer) {
	s.RegisterService(&_GroupStore_serviceDesc, srv)
}

func _GroupStore_Delete_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).Delete(&groupStoreDeleteServer{stream})
}

type GroupStore_DeleteServer interface {
	Send(*DeleteResponse) error
	Recv() (*DeleteRequest, error)
	grpc.ServerStream
}

type groupStoreDeleteServer struct {
	grpc.ServerStream
}

func (x *groupStoreDeleteServer) Send(m *DeleteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreDeleteServer) Recv() (*DeleteRequest, error) {
	m := new(DeleteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_LookupGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).LookupGroup(&groupStoreLookupGroupServer{stream})
}

type GroupStore_LookupGroupServer interface {
	Send(*LookupGroupResponse) error
	Recv() (*LookupGroupRequest, error)
	grpc.ServerStream
}

type groupStoreLookupGroupServer struct {
	grpc.ServerStream
}

func (x *groupStoreLookupGroupServer) Send(m *LookupGroupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreLookupGroupServer) Recv() (*LookupGroupRequest, error) {
	m := new(LookupGroupRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_Lookup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).Lookup(&groupStoreLookupServer{stream})
}

type GroupStore_LookupServer interface {
	Send(*LookupResponse) error
	Recv() (*LookupRequest, error)
	grpc.ServerStream
}

type groupStoreLookupServer struct {
	grpc.ServerStream
}

func (x *groupStoreLookupServer) Send(m *LookupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreLookupServer) Recv() (*LookupRequest, error) {
	m := new(LookupRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_ReadGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).ReadGroup(&groupStoreReadGroupServer{stream})
}

type GroupStore_ReadGroupServer interface {
	Send(*ReadGroupResponse) error
	Recv() (*ReadGroupRequest, error)
	grpc.ServerStream
}

type groupStoreReadGroupServer struct {
	grpc.ServerStream
}

func (x *groupStoreReadGroupServer) Send(m *ReadGroupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreReadGroupServer) Recv() (*ReadGroupRequest, error) {
	m := new(ReadGroupRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).Read(&groupStoreReadServer{stream})
}

type GroupStore_ReadServer interface {
	Send(*ReadResponse) error
	Recv() (*ReadRequest, error)
	grpc.ServerStream
}

type groupStoreReadServer struct {
	grpc.ServerStream
}

func (x *groupStoreReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreReadServer) Recv() (*ReadRequest, error) {
	m := new(ReadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).Write(&groupStoreWriteServer{stream})
}

type GroupStore_WriteServer interface {
	Send(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type groupStoreWriteServer struct {
	grpc.ServerStream
}

func (x *groupStoreWriteServer) Send(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GroupStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "groupproto.GroupStore",
	HandlerType: (*GroupStoreServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Delete",
			Handler:       _GroupStore_Delete_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "LookupGroup",
			Handler:       _GroupStore_LookupGroup_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Lookup",
			Handler:       _GroupStore_Lookup_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReadGroup",
			Handler:       _GroupStore_ReadGroup_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Read",
			Handler:       _GroupStore_Read_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Write",
			Handler:       _GroupStore_Write_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "group_api.proto",
}

func init() { proto.RegisterFile("group_api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0xcd, 0x0f, 0xea, 0x69, 0xd3, 0x0d, 0x33, 0x41, 0x56, 0x06, 0x54, 0xb9, 0x40, 0xb9,
	0x2a, 0x30, 0x1e, 0x00, 0x91, 0xad, 0x42, 0x13, 0x03, 0x4d, 0x66, 0x82, 0xcb, 0x51, 0x98, 0xb5,
	0x45, 0xb4, 0x4b, 0x48, 0x52, 0x09, 0x9e, 0x00, 0xf1, 0x14, 0x5c, 0xf0, 0x04, 0x3c, 0x04, 0x3c,
	0x17, 0x8a, 0xd3, 0xa4, 0x76, 0xeb, 0x94, 0x9b, 0xa0, 0x72, 0x67, 0xfb, 0x7c, 0x3a, 0x39, 0x9f,
	0xcf, 0xf1, 0x17, 0x6c, 0x5d, 0x24, 0xd1, 0x2c, 0x3e, 0x1b, 0xc7, 0xe1, 0x30, 0x4e, 0xa2, 0x2c,
	0xa2, 0x10, 0x07, 0x62, 0xed, 0xfd, 0x24, 0x70, 0x0e, 0xf9, 0x84, 0x67, 0x9c, 0xf1, 0x4f, 0x33,
	0x9e, 0x66, 0x74, 0x07, 0x16, 0x3b, 0x39, 0x38, 0x3a, 0x74, 0xc9, 0x80, 0xf8, 0x0e, 0x2b, 0x36,
	0x94, 0xc2, 0x7c, 0xc1, 0xbf, 0x3c, 0x73, 0x5b, 0x03, 0xe2, 0x9b, 0x4c, 0xac, 0xe7, 0x67, 0x81,
	0x6b, 0x54, 0x67, 0x01, 0xdd, 0x43, 0xfb, 0xe0, 0x32, 0x9c, 0x9c, 0x8b, 0x62, 0x53, 0x00, 0x8b,
	0x03, 0x19, 0x0d, 0x5c, 0x4b, 0x45, 0x03, 0xfa, 0x00, 0xbd, 0xd3, 0x70, 0xca, 0xd3, 0x6c, 0x3c,
	0x8d, 0x5f, 0x86, 0x1f, 0x92, 0xc8, 0xb5, 0x07, 0xc4, 0x37, 0xd8, 0xd2, 0xa9, 0xf7, 0x0e, 0xbd,
	0x52, 0x72, 0x1a, 0x47, 0x57, 0x29, 0xaf, 0xd1, 0xbc, 0xca, 0xd7, 0xd2, 0xf1, 0xd1, 0x6d, 0x18,
	0xa3, 0x24, 0x11, 0x6d, 0xb4, 0x59, 0xbe, 0xf4, 0x18, 0xe8, 0x71, 0x14, 0x7d, 0x9c, 0xc5, 0xcf,
	0xf3, 0x9b, 0x6a, 0xe4, 0x66, 0xbc, 0x18, 0x37, 0x15, 0xce, 0xb5, 0xd2, 0x1f, 0xc3, 0x3a, 0xca,
	0xf8, 0x34, 0x75, 0x5b, 0x03, 0xc3, 0xef, 0xec, 0xdf, 0x19, 0x2e, 0x2c, 0x1b, 0x4a, 0x2c, 0x79,
	0x0d, 0x2b, 0x2a, 0x35, 0x5d, 0xfc, 0x20, 0xd8, 0x5a, 0x2a, 0xae, 0xf9, 0x9c, 0xe2, 0x5a, 0x6b,
	0xad, 0x6b, 0xc6, 0xdf, 0x5d, 0x33, 0xb5, 0xb7, 0x7c, 0x0b, 0xf6, 0x31, 0xbf, 0xba, 0xc8, 0x2e,
	0x85, 0xf1, 0x0e, 0x9b, 0xef, 0xbc, 0x6f, 0x04, 0x4e, 0xa1, 0x72, 0xe3, 0x09, 0xf4, 0x3e, 0xa3,
	0x57, 0x4a, 0x69, 0x24, 0x59, 0x8b, 0x9e, 0x0d, 0xb9, 0xe7, 0xd2, 0x2b, 0x73, 0xe1, 0xd5, 0x09,
	0xb6, 0x19, 0x1f, 0x9f, 0x37, 0x98, 0xb7, 0x09, 0x6e, 0x48, 0x8c, 0x6b, 0xdb, 0x79, 0xa8, 0xa6,
	0x6d, 0x57, 0x4e, 0x5b, 0xc5, 0xb1, 0x3e, 0x6b, 0xdf, 0x09, 0x1c, 0xa5, 0x74, 0x83, 0x49, 0xdb,
	0x81, 0xf5, 0x66, 0x3c, 0x99, 0x71, 0xe1, 0x6f, 0x97, 0x15, 0x1b, 0xef, 0x2b, 0x41, 0x27, 0x57,
	0xb8, 0xf9, 0x94, 0x65, 0xe8, 0x16, 0x42, 0x1a, 0xc9, 0x58, 0xd5, 0xad, 0x21, 0x75, 0xab, 0x49,
	0xd8, 0x6f, 0x82, 0xee, 0xdb, 0x24, 0xfc, 0x1f, 0x06, 0x7d, 0x25, 0xd9, 0x96, 0x25, 0xaf, 0x36,
	0x7c, 0x5d, 0x3b, 0xfe, 0xcf, 0xe0, 0xcc, 0xfb, 0xf8, 0x37, 0xd3, 0x7f, 0xff, 0x97, 0x01, 0x88,
	0x1c, 0xbf, 0xce, 0xa2, 0x84, 0xd3, 0x11, 0xec, 0xe2, 0x77, 0x43, 0x95, 0x87, 0xa1, 0xfc, 0x35,
	0xfb, 0x7d, 0x1d, 0x54, 0xe8, 0xf3, 0xae, 0xf9, 0xe4, 0x11, 0xa1, 0xa7, 0xe8, 0x48, 0xc3, 0x98,
	0xde, 0xab, 0x19, 0xe9, 0x25, 0xe1, 0xfd, 0x5a, 0x5c, 0x61, 0x1d, 0xc1, 0x2e, 0x40, 0x55, 0x9c,
	0x32, 0x50, 0x55, 0x71, 0xea, 0x80, 0x9b, 0xd3, 0xbc, 0x42, 0xbb, 0x7a, 0xbd, 0x74, 0x4f, 0xfb,
	0xfe, 0x4b, 0xb2, 0xbb, 0x35, 0xa8, 0xc2, 0xf7, 0x14, 0x66, 0x0e, 0xd1, 0xdb, 0xcb, 0xc5, 0x25,
	0x8b, 0xbb, 0x0a, 0x28, 0x04, 0x01, 0x2c, 0x61, 0x32, 0x55, 0x0a, 0xe5, 0xfc, 0xf6, 0x77, 0x35,
	0x88, 0xcc, 0xf1, 0xde, 0x16, 0xd0, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x0f, 0x69,
	0x11, 0x01, 0x09, 0x00, 0x00,
}

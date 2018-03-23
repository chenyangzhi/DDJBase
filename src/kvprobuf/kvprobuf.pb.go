// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kvprobuf.proto

/*
Package kvprobuf is a generated protocol buffer package.

It is generated from these files:
	kvprobuf.proto

It has these top-level messages:
	Slice
	AscendIterateResponse
	PutResponse
	GetResponse
	DeleteResponse
*/
package kvprobuf

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

type Slice struct {
	Key          uint64 `protobuf:"varint,1,opt,name=Key" json:"Key,omitempty"`
	Value        []byte `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	IncludeStart bool   `protobuf:"varint,3,opt,name=IncludeStart" json:"IncludeStart,omitempty"`
}

func (m *Slice) Reset()                    { *m = Slice{} }
func (m *Slice) String() string            { return proto.CompactTextString(m) }
func (*Slice) ProtoMessage()               {}
func (*Slice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Slice) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Slice) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Slice) GetIncludeStart() bool {
	if m != nil {
		return m.IncludeStart
	}
	return false
}

type AscendIterateResponse struct {
	Slice []*Slice `protobuf:"bytes,1,rep,name=slice" json:"slice,omitempty"`
}

func (m *AscendIterateResponse) Reset()                    { *m = AscendIterateResponse{} }
func (m *AscendIterateResponse) String() string            { return proto.CompactTextString(m) }
func (*AscendIterateResponse) ProtoMessage()               {}
func (*AscendIterateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AscendIterateResponse) GetSlice() []*Slice {
	if m != nil {
		return m.Slice
	}
	return nil
}

type PutResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PutResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetResponse struct {
	Slice *Slice `protobuf:"bytes,1,opt,name=slice" json:"slice,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResponse) GetSlice() *Slice {
	if m != nil {
		return m.Slice
	}
	return nil
}

type DeleteResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Value   []byte `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeleteResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Slice)(nil), "kvprobuf.Slice")
	proto.RegisterType((*AscendIterateResponse)(nil), "kvprobuf.AscendIterateResponse")
	proto.RegisterType((*PutResponse)(nil), "kvprobuf.PutResponse")
	proto.RegisterType((*GetResponse)(nil), "kvprobuf.GetResponse")
	proto.RegisterType((*DeleteResponse)(nil), "kvprobuf.DeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DDJkv service

type DDJkvClient interface {
	AscendIterate(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*AscendIterateResponse, error)
	Put(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*PutResponse, error)
	Get(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type dDJkvClient struct {
	cc *grpc.ClientConn
}

func NewDDJkvClient(cc *grpc.ClientConn) DDJkvClient {
	return &dDJkvClient{cc}
}

func (c *dDJkvClient) AscendIterate(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*AscendIterateResponse, error) {
	out := new(AscendIterateResponse)
	err := grpc.Invoke(ctx, "/kvprobuf.DDJkv/AscendIterate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDJkvClient) Put(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/kvprobuf.DDJkv/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDJkvClient) Get(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/kvprobuf.DDJkv/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDJkvClient) Delete(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/kvprobuf.DDJkv/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DDJkv service

type DDJkvServer interface {
	AscendIterate(context.Context, *Slice) (*AscendIterateResponse, error)
	Put(context.Context, *Slice) (*PutResponse, error)
	Get(context.Context, *Slice) (*GetResponse, error)
	Delete(context.Context, *Slice) (*DeleteResponse, error)
}

func RegisterDDJkvServer(s *grpc.Server, srv DDJkvServer) {
	s.RegisterService(&_DDJkv_serviceDesc, srv)
}

func _DDJkv_AscendIterate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDJkvServer).AscendIterate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvprobuf.DDJkv/AscendIterate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDJkvServer).AscendIterate(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDJkv_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDJkvServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvprobuf.DDJkv/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDJkvServer).Put(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDJkv_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDJkvServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvprobuf.DDJkv/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDJkvServer).Get(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDJkv_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDJkvServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvprobuf.DDJkv/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDJkvServer).Delete(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

var _DDJkv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvprobuf.DDJkv",
	HandlerType: (*DDJkvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AscendIterate",
			Handler:    _DDJkv_AscendIterate_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _DDJkv_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DDJkv_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DDJkv_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvprobuf.proto",
}

func init() { proto.RegisterFile("kvprobuf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0xcc, 0x36, 0x26, 0x86, 0x97, 0x5a, 0x65, 0xb1, 0xb0, 0x78, 0x31, 0x2c, 0x88, 0x39, 0x55,
	0xa8, 0x7a, 0x15, 0x3f, 0x02, 0xa5, 0xf5, 0x52, 0x36, 0xe0, 0x3d, 0x4d, 0x9f, 0x20, 0x0d, 0x49,
	0xc8, 0xee, 0x16, 0xfc, 0x3b, 0xfe, 0x37, 0xff, 0x87, 0x34, 0xa1, 0x4d, 0x62, 0xd3, 0xde, 0x76,
	0x66, 0x99, 0x37, 0xf3, 0x86, 0x07, 0x83, 0xd5, 0x3a, 0x2f, 0xb2, 0x85, 0xfe, 0x1c, 0xe5, 0x45,
	0xa6, 0x32, 0xea, 0x6c, 0x31, 0x0f, 0xc1, 0x0a, 0x93, 0xaf, 0x18, 0xe9, 0x05, 0x98, 0xef, 0xf8,
	0xcd, 0x88, 0x47, 0xfc, 0x13, 0xb1, 0x79, 0xd2, 0x4b, 0xb0, 0x3e, 0xa2, 0x44, 0x23, 0xeb, 0x79,
	0xc4, 0xef, 0x8b, 0x0a, 0x50, 0x0e, 0xfd, 0x69, 0x1a, 0x27, 0x7a, 0x89, 0xa1, 0x8a, 0x0a, 0xc5,
	0x4c, 0x8f, 0xf8, 0x8e, 0x68, 0x71, 0xfc, 0x09, 0x86, 0x2f, 0x32, 0xc6, 0x74, 0x39, 0x55, 0x58,
	0x44, 0x0a, 0x05, 0xca, 0x3c, 0x4b, 0x25, 0xd2, 0x1b, 0xb0, 0xe4, 0xc6, 0x8d, 0x11, 0xcf, 0xf4,
	0xdd, 0xf1, 0xf9, 0x68, 0x97, 0xab, 0x0c, 0x21, 0xaa, 0x5f, 0x7e, 0x0b, 0xee, 0x5c, 0xab, 0x9d,
	0x8a, 0xc1, 0x69, 0xa8, 0xe3, 0x18, 0xa5, 0x2c, 0xe3, 0x39, 0x62, 0x0b, 0xf9, 0x03, 0xb8, 0x13,
	0x54, 0x5d, 0xe3, 0xc9, 0x91, 0xf1, 0xcf, 0x30, 0x08, 0x30, 0xc1, 0x46, 0xae, 0x83, 0x0e, 0xdd,
	0x25, 0x8c, 0x7f, 0x09, 0x58, 0x41, 0x30, 0x5b, 0xad, 0xe9, 0x1b, 0x9c, 0xb5, 0x56, 0xa5, 0xff,
	0x4d, 0xaf, 0xae, 0x6b, 0xa2, 0xb3, 0x14, 0x6e, 0xd0, 0x3b, 0x30, 0xe7, 0x5a, 0xed, 0x4b, 0x87,
	0x35, 0xd1, 0xe8, 0xa3, 0x12, 0x4c, 0xf0, 0xb8, 0xa0, 0xd1, 0x0b, 0x37, 0xe8, 0x23, 0xd8, 0xd5,
	0xca, 0xfb, 0x1a, 0x56, 0x13, 0xed, 0x56, 0xb8, 0xf1, 0x6a, 0xff, 0xf4, 0xcc, 0x20, 0x98, 0x2d,
	0xec, 0xf2, 0x6c, 0xee, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x86, 0xc9, 0xce, 0x48, 0x02,
	0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: KVStore.proto

package KVStore

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

type Error_Type int32

const (
	Error_NO_ERROR  Error_Type = 0
	Error_GET_ERROR Error_Type = 1
	Error_PUT_ERROR Error_Type = 2
	Error_DEL_ERROR Error_Type = 3
)

var Error_Type_name = map[int32]string{
	0: "NO_ERROR",
	1: "GET_ERROR",
	2: "PUT_ERROR",
	3: "DEL_ERROR",
}

var Error_Type_value = map[string]int32{
	"NO_ERROR":  0,
	"GET_ERROR": 1,
	"PUT_ERROR": 2,
	"DEL_ERROR": 3,
}

func (x Error_Type) String() string {
	return proto.EnumName(Error_Type_name, int32(x))
}

func (Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4a01cb53464fd04b, []int{4, 0}
}

type GetReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a01cb53464fd04b, []int{0}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetRes struct {
	Val                  string   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	Err                  *Error   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRes) Reset()         { *m = GetRes{} }
func (m *GetRes) String() string { return proto.CompactTextString(m) }
func (*GetRes) ProtoMessage()    {}
func (*GetRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a01cb53464fd04b, []int{1}
}

func (m *GetRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRes.Unmarshal(m, b)
}
func (m *GetRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRes.Marshal(b, m, deterministic)
}
func (m *GetRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRes.Merge(m, src)
}
func (m *GetRes) XXX_Size() int {
	return xxx_messageInfo_GetRes.Size(m)
}
func (m *GetRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetRes proto.InternalMessageInfo

func (m *GetRes) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func (m *GetRes) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type PutReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  string   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutReq) Reset()         { *m = PutReq{} }
func (m *PutReq) String() string { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()    {}
func (*PutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a01cb53464fd04b, []int{2}
}

func (m *PutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutReq.Unmarshal(m, b)
}
func (m *PutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutReq.Marshal(b, m, deterministic)
}
func (m *PutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutReq.Merge(m, src)
}
func (m *PutReq) XXX_Size() int {
	return xxx_messageInfo_PutReq.Size(m)
}
func (m *PutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PutReq.DiscardUnknown(m)
}

var xxx_messageInfo_PutReq proto.InternalMessageInfo

func (m *PutReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutReq) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type PutRes struct {
	Err                  *Error   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRes) Reset()         { *m = PutRes{} }
func (m *PutRes) String() string { return proto.CompactTextString(m) }
func (*PutRes) ProtoMessage()    {}
func (*PutRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a01cb53464fd04b, []int{3}
}

func (m *PutRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRes.Unmarshal(m, b)
}
func (m *PutRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRes.Marshal(b, m, deterministic)
}
func (m *PutRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRes.Merge(m, src)
}
func (m *PutRes) XXX_Size() int {
	return xxx_messageInfo_PutRes.Size(m)
}
func (m *PutRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRes.DiscardUnknown(m)
}

var xxx_messageInfo_PutRes proto.InternalMessageInfo

func (m *PutRes) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type Error struct {
	Type                 Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=Error_Type" json:"type,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a01cb53464fd04b, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_NO_ERROR
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("Error_Type", Error_Type_name, Error_Type_value)
	proto.RegisterType((*GetReq)(nil), "GetReq")
	proto.RegisterType((*GetRes)(nil), "GetRes")
	proto.RegisterType((*PutReq)(nil), "PutReq")
	proto.RegisterType((*PutRes)(nil), "PutRes")
	proto.RegisterType((*Error)(nil), "Error")
}

func init() {
	proto.RegisterFile("KVStore.proto", fileDescriptor_4a01cb53464fd04b)
}

var fileDescriptor_4a01cb53464fd04b = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xb3, 0x49, 0x4d, 0xec, 0xd4, 0x4a, 0xd9, 0x53, 0xcc, 0xc5, 0xb2, 0xa7, 0x1e, 0x24,
	0x87, 0xe8, 0x5d, 0x04, 0x43, 0x0e, 0x8a, 0x0d, 0x6b, 0xf5, 0x2a, 0x11, 0x06, 0x0f, 0x56, 0x12,
	0x77, 0xb7, 0x81, 0xfc, 0x00, 0xff, 0xb7, 0x4c, 0x32, 0xeb, 0xc5, 0xde, 0xe6, 0xdb, 0xef, 0xf1,
	0x16, 0x1e, 0x2c, 0x1f, 0x5e, 0x9f, 0x5d, 0x6b, 0x30, 0xef, 0x4c, 0xeb, 0x5a, 0x95, 0x41, 0x5c,
	0xa1, 0xd3, 0xf8, 0x2d, 0x57, 0x10, 0x7d, 0xe2, 0x90, 0x8a, 0xb5, 0xd8, 0xcc, 0x35, 0x9d, 0xea,
	0x86, 0x9d, 0x25, 0xd7, 0x37, 0x7b, 0xef, 0xfa, 0x66, 0x2f, 0x53, 0x88, 0xd0, 0x98, 0x34, 0x5c,
	0x8b, 0xcd, 0xa2, 0x88, 0xf3, 0xd2, 0x98, 0xd6, 0x68, 0x7a, 0x52, 0x57, 0x10, 0xd7, 0x87, 0xe3,
	0x8d, 0xbe, 0x27, 0xfc, 0xeb, 0x51, 0x8a, 0xd3, 0xd6, 0x37, 0x8a, 0xff, 0x8d, 0x3f, 0x02, 0x4e,
	0x46, 0x94, 0x97, 0x30, 0x73, 0x43, 0x87, 0x63, 0xe8, 0xbc, 0x58, 0x4c, 0xa1, 0x7c, 0x37, 0x74,
	0xa8, 0x47, 0x21, 0x53, 0x48, 0xbe, 0xd0, 0xda, 0xe6, 0x03, 0xf9, 0x13, 0x8f, 0xea, 0x0e, 0x66,
	0x94, 0x93, 0x67, 0x70, 0xfa, 0xb4, 0x7d, 0x2b, 0xb5, 0xde, 0xea, 0x55, 0x20, 0x97, 0x30, 0xaf,
	0xca, 0x1d, 0xa3, 0x20, 0xac, 0x5f, 0x3c, 0x86, 0x84, 0xf7, 0xe5, 0x23, 0x63, 0x54, 0xdc, 0x42,
	0xc2, 0xe3, 0xc9, 0x0b, 0x88, 0x2a, 0x74, 0x32, 0xc9, 0xa7, 0xf1, 0x32, 0x3e, 0xac, 0x0a, 0x48,
	0xd5, 0x07, 0x52, 0xd3, 0x0a, 0x19, 0x1f, 0x56, 0x05, 0xef, 0xf1, 0xb8, 0xf9, 0xf5, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x62, 0x21, 0x14, 0x9b, 0x84, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVStoreClient is the client API for KVStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVStoreClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error)
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRes, error)
}

type kVStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKVStoreClient(cc grpc.ClientConnInterface) KVStoreClient {
	return &kVStoreClient{cc}
}

func (c *kVStoreClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error) {
	out := new(GetRes)
	err := c.cc.Invoke(ctx, "/KVStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRes, error) {
	out := new(PutRes)
	err := c.cc.Invoke(ctx, "/KVStore/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVStoreServer is the server API for KVStore service.
type KVStoreServer interface {
	Get(context.Context, *GetReq) (*GetRes, error)
	Put(context.Context, *PutReq) (*PutRes, error)
}

// UnimplementedKVStoreServer can be embedded to have forward compatible implementations.
type UnimplementedKVStoreServer struct {
}

func (*UnimplementedKVStoreServer) Get(ctx context.Context, req *GetReq) (*GetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVStoreServer) Put(ctx context.Context, req *PutReq) (*PutRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}

func RegisterKVStoreServer(s *grpc.Server, srv KVStoreServer) {
	s.RegisterService(&_KVStore_serviceDesc, srv)
}

func _KVStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVStore/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KVStore",
	HandlerType: (*KVStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KVStore_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KVStore_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "KVStore.proto",
}

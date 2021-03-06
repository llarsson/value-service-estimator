// Code generated by protoc-gen-go. DO NOT EDIT.
// source: value.proto

package value

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	log "log"
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

// Wrapper for the value that this service handles.
type Value struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{0}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// gRPC workaround object, since we do not have a "null" keyword.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d8b663a521ecf69, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Value)(nil), "Value")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("value.proto", fileDescriptor_6d8b663a521ecf69) }

var fileDescriptor_6d8b663a521ecf69 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4b, 0xcc, 0x29,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe5, 0x62, 0x0d, 0x03, 0x71, 0x85, 0x44,
	0xb8, 0x58, 0xc1, 0xe2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x12, 0x3b, 0x17,
	0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x91, 0x17, 0x17, 0x0f, 0x58, 0x5d, 0x70, 0x6a, 0x51, 0x59,
	0x66, 0x72, 0xaa, 0x90, 0x0c, 0x17, 0x87, 0x7b, 0x6a, 0x09, 0x44, 0x2b, 0x9b, 0x1e, 0x58, 0x8d,
	0x14, 0x9b, 0x1e, 0x98, 0xaf, 0xc4, 0x00, 0x92, 0x0d, 0x46, 0xc8, 0x82, 0x69, 0x29, 0xa8, 0x2a,
	0x25, 0x86, 0x24, 0x36, 0xb0, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x31, 0x67,
	0x81, 0x89, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ValueServiceClient is the client API for ValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValueServiceClient interface {
	// Gets the value that this service handles.
	GetValue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Value, error)
	// Sets the value that this service handles.
	SetValue(ctx context.Context, in *Value, opts ...grpc.CallOption) (*Empty, error)
}

type valueServiceClient struct {
	cc *grpc.ClientConn
}

func NewValueServiceClient(cc *grpc.ClientConn) ValueServiceClient {
	return &valueServiceClient{cc}
}

func (c *valueServiceClient) GetValue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/ValueService/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueServiceClient) SetValue(ctx context.Context, in *Value, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ValueService/SetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValueServiceServer is the server API for ValueService service.
type ValueServiceServer interface {
	// Gets the value that this service handles.
	GetValue(context.Context, *Empty) (*Value, error)
	// Sets the value that this service handles.
	SetValue(context.Context, *Value) (*Empty, error)
}

// UnimplementedValueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedValueServiceServer struct {
}

func (*UnimplementedValueServiceServer) GetValue(ctx context.Context, req *Empty) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (*UnimplementedValueServiceServer) SetValue(ctx context.Context, req *Value) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}

// ValueServiceProxy creates a proxy for the
type ValueServiceProxy struct {
	Client ValueServiceClient
}

func (p *ValueServiceProxy) GetValue(ctx context.Context, req *Empty) (*Value, error) {
	response, err := p.Client.GetValue(ctx, req)
	if err != nil {
		log.Printf("Failed to call upstream ValueService.GetValue")
		return nil, err
	}

	return response, nil
}

func (p *ValueServiceProxy) SetValue(ctx context.Context, req *Value) (*Empty, error) {
	response, err := p.Client.SetValue(ctx, req)
	if err != nil {
		log.Printf("Failed to call upstream ValueService.SetValue")
		return nil, err
	}

	return response, nil
}

func RegisterValueServiceServer(s *grpc.Server, srv ValueServiceServer) {
	s.RegisterService(&_ValueService_serviceDesc, srv)
}

func _ValueService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ValueService/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueServiceServer).GetValue(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueService_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueServiceServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ValueService/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueServiceServer).SetValue(ctx, req.(*Value))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ValueService",
	HandlerType: (*ValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _ValueService_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _ValueService_SetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "value.proto",
}

// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

/*
Package login is a generated protocol buffer package.

It is generated from these files:
	login.proto

It has these top-level messages:
	LoginRequest
	LoginReply
*/
package login

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

// grpc客户端请求时传输给服务端的对象
type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// grpc服务端返回给客户端的对象
type LoginReply struct {
	Status   bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	ErrorMsg string `protobuf:"bytes,2,opt,name=errorMsg" json:"errorMsg,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "login.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "login.LoginReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Loginer service

type LoginerClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type loginerClient struct {
	cc *grpc.ClientConn
}

func NewLoginerClient(cc *grpc.ClientConn) LoginerClient {
	return &loginerClient{cc}
}

func (c *loginerClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/login.Loginer/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Loginer service

type LoginerServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterLoginerServer(s *grpc.Server, srv LoginerServer) {
	s.RegisterService(&_Loginer_serviceDesc, srv)
}

func _Loginer_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Loginer/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginerServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Loginer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "login.Loginer",
	HandlerType: (*LoginerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Loginer_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xdc, 0xb8, 0x78, 0x7c,
	0x40, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x8e, 0xd2, 0xe2, 0xd4,
	0xa2, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x1f, 0x24, 0x57,
	0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x94, 0x22, 0xc1, 0x04, 0x91, 0x83, 0xf1, 0x95, 0x1c, 0xb8,
	0xb8, 0xa0, 0xe6, 0x14, 0xe4, 0x54, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16,
	0x83, 0xcd, 0xe0, 0x08, 0x82, 0xf2, 0x40, 0x26, 0xa4, 0x16, 0x15, 0xe5, 0x17, 0xf9, 0x16, 0xa7,
	0xc3, 0x4c, 0x80, 0xf1, 0x8d, 0x6c, 0xb8, 0xd8, 0xc1, 0x26, 0xa4, 0x16, 0x09, 0x19, 0x72, 0xb1,
	0x82, 0x99, 0x42, 0xc2, 0x7a, 0x10, 0x27, 0x23, 0x3b, 0x51, 0x4a, 0x10, 0x55, 0xb0, 0x20, 0xa7,
	0x52, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x2b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x7e,
	0x0e, 0x04, 0xe4, 0x00, 0x00, 0x00,
}

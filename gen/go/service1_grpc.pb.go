// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: service1.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestService1Client is the client API for TestService1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestService1Client interface {
	Check(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type testService1Client struct {
	cc grpc.ClientConnInterface
}

func NewTestService1Client(cc grpc.ClientConnInterface) TestService1Client {
	return &testService1Client{cc}
}

func (c *testService1Client) Check(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/test.my_service1.v1.TestService1/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestService1Server is the server API for TestService1 service.
// All implementations must embed UnimplementedTestService1Server
// for forward compatibility
type TestService1Server interface {
	Check(context.Context, *StringMessage) (*StringMessage, error)
	mustEmbedUnimplementedTestService1Server()
}

// UnimplementedTestService1Server must be embedded to have forward compatible implementations.
type UnimplementedTestService1Server struct {
}

func (UnimplementedTestService1Server) Check(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedTestService1Server) mustEmbedUnimplementedTestService1Server() {}

// UnsafeTestService1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestService1Server will
// result in compilation errors.
type UnsafeTestService1Server interface {
	mustEmbedUnimplementedTestService1Server()
}

func RegisterTestService1Server(s grpc.ServiceRegistrar, srv TestService1Server) {
	s.RegisterService(&TestService1_ServiceDesc, srv)
}

func _TestService1_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestService1Server).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.my_service1.v1.TestService1/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestService1Server).Check(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService1_ServiceDesc is the grpc.ServiceDesc for TestService1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.my_service1.v1.TestService1",
	HandlerType: (*TestService1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _TestService1_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service1.proto",
}

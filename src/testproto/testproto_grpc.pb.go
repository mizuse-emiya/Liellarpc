// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testproto

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

// GoTestClient is the client API for GoTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoTestClient interface {
	// rpc请求 请求的函数 (发送请求参数) returns (返回响应的参数)
	GetAddressResponse(ctx context.Context, in *SendAddress, opts ...grpc.CallOption) (*GetResponse, error)
}

type goTestClient struct {
	cc grpc.ClientConnInterface
}

func NewGoTestClient(cc grpc.ClientConnInterface) GoTestClient {
	return &goTestClient{cc}
}

func (c *goTestClient) GetAddressResponse(ctx context.Context, in *SendAddress, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/testproto.GoTest/GetAddressResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoTestServer is the server API for GoTest service.
// All implementations must embed UnimplementedGoTestServer
// for forward compatibility
type GoTestServer interface {
	// rpc请求 请求的函数 (发送请求参数) returns (返回响应的参数)
	GetAddressResponse(context.Context, *SendAddress) (*GetResponse, error)
	mustEmbedUnimplementedGoTestServer()
}

// UnimplementedGoTestServer must be embedded to have forward compatible implementations.
type UnimplementedGoTestServer struct {
}

func (UnimplementedGoTestServer) GetAddressResponse(context.Context, *SendAddress) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressResponse not implemented")
}
func (UnimplementedGoTestServer) mustEmbedUnimplementedGoTestServer() {}

// UnsafeGoTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoTestServer will
// result in compilation errors.
type UnsafeGoTestServer interface {
	mustEmbedUnimplementedGoTestServer()
}

func RegisterGoTestServer(s grpc.ServiceRegistrar, srv GoTestServer) {
	s.RegisterService(&GoTest_ServiceDesc, srv)
}

func _GoTest_GetAddressResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTestServer).GetAddressResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testproto.GoTest/GetAddressResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTestServer).GetAddressResponse(ctx, req.(*SendAddress))
	}
	return interceptor(ctx, in, info, handler)
}

// GoTest_ServiceDesc is the grpc.ServiceDesc for GoTest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoTest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testproto.GoTest",
	HandlerType: (*GoTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddressResponse",
			Handler:    _GoTest_GetAddressResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testproto.proto",
}
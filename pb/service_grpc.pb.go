// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// CalculaterClient is the client API for Calculater service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculaterClient interface {
	Calculate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type calculaterClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculaterClient(cc grpc.ClientConnInterface) CalculaterClient {
	return &calculaterClient{cc}
}

func (c *calculaterClient) Calculate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Calculater/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculaterServer is the server API for Calculater service.
// All implementations must embed UnimplementedCalculaterServer
// for forward compatibility
type CalculaterServer interface {
	Calculate(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedCalculaterServer()
}

// UnimplementedCalculaterServer must be embedded to have forward compatible implementations.
type UnimplementedCalculaterServer struct {
}

func (UnimplementedCalculaterServer) Calculate(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculaterServer) mustEmbedUnimplementedCalculaterServer() {}

// UnsafeCalculaterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculaterServer will
// result in compilation errors.
type UnsafeCalculaterServer interface {
	mustEmbedUnimplementedCalculaterServer()
}

func RegisterCalculaterServer(s grpc.ServiceRegistrar, srv CalculaterServer) {
	s.RegisterService(&Calculater_ServiceDesc, srv)
}

func _Calculater_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculaterServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculater/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculaterServer).Calculate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculater_ServiceDesc is the grpc.ServiceDesc for Calculater service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculater_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Calculater",
	HandlerType: (*CalculaterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Calculater_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0--rc1
// source: perc.proto

package perc

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

// PercClient is the client API for Perc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PercClient interface {
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	CalculateFractional(ctx context.Context, in *CalculateFractionalRequest, opts ...grpc.CallOption) (*CalculateFractionalResponse, error)
}

type percClient struct {
	cc grpc.ClientConnInterface
}

func NewPercClient(cc grpc.ClientConnInterface) PercClient {
	return &percClient{cc}
}

func (c *percClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/perc.Perc/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *percClient) CalculateFractional(ctx context.Context, in *CalculateFractionalRequest, opts ...grpc.CallOption) (*CalculateFractionalResponse, error) {
	out := new(CalculateFractionalResponse)
	err := c.cc.Invoke(ctx, "/perc.Perc/CalculateFractional", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PercServer is the server API for Perc service.
// All implementations must embed UnimplementedPercServer
// for forward compatibility
type PercServer interface {
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	CalculateFractional(context.Context, *CalculateFractionalRequest) (*CalculateFractionalResponse, error)
	mustEmbedUnimplementedPercServer()
}

// UnimplementedPercServer must be embedded to have forward compatible implementations.
type UnimplementedPercServer struct {
}

func (UnimplementedPercServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedPercServer) CalculateFractional(context.Context, *CalculateFractionalRequest) (*CalculateFractionalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateFractional not implemented")
}
func (UnimplementedPercServer) mustEmbedUnimplementedPercServer() {}

// UnsafePercServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PercServer will
// result in compilation errors.
type UnsafePercServer interface {
	mustEmbedUnimplementedPercServer()
}

func RegisterPercServer(s grpc.ServiceRegistrar, srv PercServer) {
	s.RegisterService(&Perc_ServiceDesc, srv)
}

func _Perc_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PercServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/perc.Perc/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PercServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perc_CalculateFractional_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateFractionalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PercServer).CalculateFractional(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/perc.Perc/CalculateFractional",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PercServer).CalculateFractional(ctx, req.(*CalculateFractionalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Perc_ServiceDesc is the grpc.ServiceDesc for Perc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Perc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "perc.Perc",
	HandlerType: (*PercServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Perc_Calculate_Handler,
		},
		{
			MethodName: "CalculateFractional",
			Handler:    _Perc_CalculateFractional_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "perc.proto",
}

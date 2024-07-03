// Code generated with goa v3.17.1, DO NOT EDIT.
//
// calc protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/error/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: goagen_error_calc.proto

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Calc_Divide_FullMethodName = "/calc.Calc/Divide"
)

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service is the calc service interface.
type CalcClient interface {
	// Divide implements divide.
	Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error)
}

type calcClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcClient(cc grpc.ClientConnInterface) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DivideResponse)
	err := c.cc.Invoke(ctx, Calc_Divide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
// All implementations must embed UnimplementedCalcServer
// for forward compatibility
//
// Service is the calc service interface.
type CalcServer interface {
	// Divide implements divide.
	Divide(context.Context, *DivideRequest) (*DivideResponse, error)
	mustEmbedUnimplementedCalcServer()
}

// UnimplementedCalcServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (UnimplementedCalcServer) Divide(context.Context, *DivideRequest) (*DivideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedCalcServer) mustEmbedUnimplementedCalcServer() {}

// UnsafeCalcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServer will
// result in compilation errors.
type UnsafeCalcServer interface {
	mustEmbedUnimplementedCalcServer()
}

func RegisterCalcServer(s grpc.ServiceRegistrar, srv CalcServer) {
	s.RegisterService(&Calc_ServiceDesc, srv)
}

func _Calc_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calc_Divide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Divide(ctx, req.(*DivideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Calc_ServiceDesc is the grpc.ServiceDesc for Calc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calc.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Divide",
			Handler:    _Calc_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_error_calc.proto",
}

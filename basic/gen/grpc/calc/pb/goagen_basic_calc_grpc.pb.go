// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: goagen_basic_calc.proto

package calcpb

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

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcClient interface {
	// Multiply implements multiply.
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
}

type calcClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcClient(cc grpc.ClientConnInterface) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, "/calc.Calc/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
// All implementations must embed UnimplementedCalcServer
// for forward compatibility
type CalcServer interface {
	// Multiply implements multiply.
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
	mustEmbedUnimplementedCalcServer()
}

// UnimplementedCalcServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (UnimplementedCalcServer) Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
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

func _Calc_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Multiply(ctx, req.(*MultiplyRequest))
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
			MethodName: "Multiply",
			Handler:    _Calc_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_basic_calc.proto",
}

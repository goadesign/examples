// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: goagen_multiauth_secured_service.proto

package secured_servicepb

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

// SecuredServiceClient is the client API for SecuredService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecuredServiceClient interface {
	// Creates a valid JWT
	Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninResponse, error)
	// This action is secured with the jwt scheme
	Secure(ctx context.Context, in *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// query string.
	DoublySecure(ctx context.Context, in *DoublySecureRequest, opts ...grpc.CallOption) (*DoublySecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// header.
	AlsoDoublySecure(ctx context.Context, in *AlsoDoublySecureRequest, opts ...grpc.CallOption) (*AlsoDoublySecureResponse, error)
}

type securedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecuredServiceClient(cc grpc.ClientConnInterface) SecuredServiceClient {
	return &securedServiceClient{cc}
}

func (c *securedServiceClient) Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninResponse, error) {
	out := new(SigninResponse)
	err := c.cc.Invoke(ctx, "/secured_service.SecuredService/Signin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securedServiceClient) Secure(ctx context.Context, in *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error) {
	out := new(SecureResponse)
	err := c.cc.Invoke(ctx, "/secured_service.SecuredService/Secure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securedServiceClient) DoublySecure(ctx context.Context, in *DoublySecureRequest, opts ...grpc.CallOption) (*DoublySecureResponse, error) {
	out := new(DoublySecureResponse)
	err := c.cc.Invoke(ctx, "/secured_service.SecuredService/DoublySecure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securedServiceClient) AlsoDoublySecure(ctx context.Context, in *AlsoDoublySecureRequest, opts ...grpc.CallOption) (*AlsoDoublySecureResponse, error) {
	out := new(AlsoDoublySecureResponse)
	err := c.cc.Invoke(ctx, "/secured_service.SecuredService/AlsoDoublySecure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecuredServiceServer is the server API for SecuredService service.
// All implementations must embed UnimplementedSecuredServiceServer
// for forward compatibility
type SecuredServiceServer interface {
	// Creates a valid JWT
	Signin(context.Context, *SigninRequest) (*SigninResponse, error)
	// This action is secured with the jwt scheme
	Secure(context.Context, *SecureRequest) (*SecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// query string.
	DoublySecure(context.Context, *DoublySecureRequest) (*DoublySecureResponse, error)
	// This action is secured with the jwt scheme and also requires an API key
	// header.
	AlsoDoublySecure(context.Context, *AlsoDoublySecureRequest) (*AlsoDoublySecureResponse, error)
	mustEmbedUnimplementedSecuredServiceServer()
}

// UnimplementedSecuredServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecuredServiceServer struct {
}

func (UnimplementedSecuredServiceServer) Signin(context.Context, *SigninRequest) (*SigninResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signin not implemented")
}
func (UnimplementedSecuredServiceServer) Secure(context.Context, *SecureRequest) (*SecureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Secure not implemented")
}
func (UnimplementedSecuredServiceServer) DoublySecure(context.Context, *DoublySecureRequest) (*DoublySecureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoublySecure not implemented")
}
func (UnimplementedSecuredServiceServer) AlsoDoublySecure(context.Context, *AlsoDoublySecureRequest) (*AlsoDoublySecureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlsoDoublySecure not implemented")
}
func (UnimplementedSecuredServiceServer) mustEmbedUnimplementedSecuredServiceServer() {}

// UnsafeSecuredServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecuredServiceServer will
// result in compilation errors.
type UnsafeSecuredServiceServer interface {
	mustEmbedUnimplementedSecuredServiceServer()
}

func RegisterSecuredServiceServer(s grpc.ServiceRegistrar, srv SecuredServiceServer) {
	s.RegisterService(&SecuredService_ServiceDesc, srv)
}

func _SecuredService_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigninRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secured_service.SecuredService/Signin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).Signin(ctx, req.(*SigninRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuredService_Secure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).Secure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secured_service.SecuredService/Secure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).Secure(ctx, req.(*SecureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuredService_DoublySecure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoublySecureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).DoublySecure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secured_service.SecuredService/DoublySecure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).DoublySecure(ctx, req.(*DoublySecureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuredService_AlsoDoublySecure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlsoDoublySecureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuredServiceServer).AlsoDoublySecure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secured_service.SecuredService/AlsoDoublySecure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuredServiceServer).AlsoDoublySecure(ctx, req.(*AlsoDoublySecureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecuredService_ServiceDesc is the grpc.ServiceDesc for SecuredService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecuredService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "secured_service.SecuredService",
	HandlerType: (*SecuredServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signin",
			Handler:    _SecuredService_Signin_Handler,
		},
		{
			MethodName: "Secure",
			Handler:    _SecuredService_Secure_Handler,
		},
		{
			MethodName: "DoublySecure",
			Handler:    _SecuredService_DoublySecure_Handler,
		},
		{
			MethodName: "AlsoDoublySecure",
			Handler:    _SecuredService_AlsoDoublySecure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_multiauth_secured_service.proto",
}

// Code generated with goa v3.17.1, DO NOT EDIT.
//
// storage protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/cellar/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: goagen_cellar_storage.proto

package storagepb

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
	Storage_List_FullMethodName        = "/storage.Storage/List"
	Storage_Show_FullMethodName        = "/storage.Storage/Show"
	Storage_Add_FullMethodName         = "/storage.Storage/Add"
	Storage_Remove_FullMethodName      = "/storage.Storage/Remove"
	Storage_Rate_FullMethodName        = "/storage.Storage/Rate"
	Storage_MultiAdd_FullMethodName    = "/storage.Storage/MultiAdd"
	Storage_MultiUpdate_FullMethodName = "/storage.Storage/MultiUpdate"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The storage service makes it possible to view, add or remove wine bottles.
type StorageClient interface {
	// List all stored bottles
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error)
	// Show bottle by ID
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
	// Add new bottle and return its ID.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Remove bottle from storage
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	// Rate bottles by IDs
	Rate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	// Add n number of bottles and return their IDs. This is a multipart request
	// and each part has field name 'bottle' and contains the encoded bottle info
	// to be added.
	MultiAdd(ctx context.Context, in *MultiAddRequest, opts ...grpc.CallOption) (*MultiAddResponse, error)
	// Update bottles with the given IDs. This is a multipart request and each part
	// has field name 'bottle' and contains the encoded bottle info to be updated.
	// The IDs in the query parameter is mapped to each part in the request.
	MultiUpdate(ctx context.Context, in *MultiUpdateRequest, opts ...grpc.CallOption) (*MultiUpdateResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoredBottleCollection)
	err := c.cc.Invoke(ctx, Storage_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, Storage_Show_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, Storage_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, Storage_Remove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Rate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, Storage_Rate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) MultiAdd(ctx context.Context, in *MultiAddRequest, opts ...grpc.CallOption) (*MultiAddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultiAddResponse)
	err := c.cc.Invoke(ctx, Storage_MultiAdd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) MultiUpdate(ctx context.Context, in *MultiUpdateRequest, opts ...grpc.CallOption) (*MultiUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultiUpdateResponse)
	err := c.cc.Invoke(ctx, Storage_MultiUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
//
// The storage service makes it possible to view, add or remove wine bottles.
type StorageServer interface {
	// List all stored bottles
	List(context.Context, *ListRequest) (*StoredBottleCollection, error)
	// Show bottle by ID
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
	// Add new bottle and return its ID.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Remove bottle from storage
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	// Rate bottles by IDs
	Rate(context.Context, *RateRequest) (*RateResponse, error)
	// Add n number of bottles and return their IDs. This is a multipart request
	// and each part has field name 'bottle' and contains the encoded bottle info
	// to be added.
	MultiAdd(context.Context, *MultiAddRequest) (*MultiAddResponse, error)
	// Update bottles with the given IDs. This is a multipart request and each part
	// has field name 'bottle' and contains the encoded bottle info to be updated.
	// The IDs in the query parameter is mapped to each part in the request.
	MultiUpdate(context.Context, *MultiUpdateRequest) (*MultiUpdateResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) List(context.Context, *ListRequest) (*StoredBottleCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStorageServer) Show(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedStorageServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedStorageServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedStorageServer) Rate(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rate not implemented")
}
func (UnimplementedStorageServer) MultiAdd(context.Context, *MultiAddRequest) (*MultiAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiAdd not implemented")
}
func (UnimplementedStorageServer) MultiUpdate(context.Context, *MultiUpdateRequest) (*MultiUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiUpdate not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Rate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Rate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Rate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Rate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_MultiAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).MultiAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_MultiAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).MultiAdd(ctx, req.(*MultiAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_MultiUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).MultiUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_MultiUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).MultiUpdate(ctx, req.(*MultiUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Storage_List_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Storage_Show_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Storage_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Storage_Remove_Handler,
		},
		{
			MethodName: "Rate",
			Handler:    _Storage_Rate_Handler,
		},
		{
			MethodName: "MultiAdd",
			Handler:    _Storage_MultiAdd_Handler,
		},
		{
			MethodName: "MultiUpdate",
			Handler:    _Storage_MultiUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_cellar_storage.proto",
}

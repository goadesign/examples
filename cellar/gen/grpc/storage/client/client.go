// Code generated by goa v3.21.5, DO NOT EDIT.
//
// storage gRPC client
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package client

import (
	"context"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli storagepb.StorageClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the storage service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: storagepb.NewStorageClient(cc),
		opts:    opts,
	}
}

// List calls the "List" function in storagepb.StorageClient interface.
func (c *Client) List() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildListFunc(c.grpccli, c.opts...),
			nil,
			DecodeListResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault("%s", err.Error())
		}
		return res, nil
	}
}

// Show calls the "Show" function in storagepb.StorageClient interface.
func (c *Client) Show() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildShowFunc(c.grpccli, c.opts...),
			EncodeShowRequest,
			DecodeShowResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *storagepb.ShowNotFoundError:
				return nil, NewShowNotFoundError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault("%s", err.Error())
			}
		}
		return res, nil
	}
}

// Add calls the "Add" function in storagepb.StorageClient interface.
func (c *Client) Add() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildAddFunc(c.grpccli, c.opts...),
			EncodeAddRequest,
			DecodeAddResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault("%s", err.Error())
		}
		return res, nil
	}
}

// Remove calls the "Remove" function in storagepb.StorageClient interface.
func (c *Client) Remove() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildRemoveFunc(c.grpccli, c.opts...),
			EncodeRemoveRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault("%s", err.Error())
		}
		return res, nil
	}
}

// Rate calls the "Rate" function in storagepb.StorageClient interface.
func (c *Client) Rate() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildRateFunc(c.grpccli, c.opts...),
			EncodeRateRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault("%s", err.Error())
		}
		return res, nil
	}
}

// MultiAdd calls the "MultiAdd" function in storagepb.StorageClient interface.
func (c *Client) MultiAdd() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildMultiAddFunc(c.grpccli, c.opts...),
			EncodeMultiAddRequest,
			DecodeMultiAddResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault("%s", err.Error())
		}
		return res, nil
	}
}

// MultiUpdate calls the "MultiUpdate" function in storagepb.StorageClient
// interface.
func (c *Client) MultiUpdate() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildMultiUpdateFunc(c.grpccli, c.opts...),
			EncodeMultiUpdateRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault("%s", err.Error())
		}
		return res, nil
	}
}

// Code generated by goa v2.1.2, DO NOT EDIT.
//
// divider gRPC client
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"context"

	dividerpb "goa.design/examples/error/gen/grpc/divider/pb"
	goa "goa.design/goa"
	goagrpc "goa.design/goa/grpc"
	goapb "goa.design/goa/grpc/pb"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli dividerpb.DividerClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the divider service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: dividerpb.NewDividerClient(cc),
		opts:    opts,
	}
}

// IntegerDivide calls the "IntegerDivide" function in dividerpb.DividerClient
// interface.
func (c *Client) IntegerDivide() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildIntegerDivideFunc(c.grpccli, c.opts...),
			EncodeIntegerDivideRequest,
			DecodeIntegerDivideResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Divide calls the "Divide" function in dividerpb.DividerClient interface.
func (c *Client) Divide() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDivideFunc(c.grpccli, c.opts...),
			EncodeDivideRequest,
			DecodeDivideResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Code generated by goa v3.19.1, DO NOT EDIT.
//
// interceptors client
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package interceptors

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "interceptors" service client.
type Client struct {
	GetEndpoint    goa.Endpoint
	CreateEndpoint goa.Endpoint
	StreamEndpoint goa.Endpoint
}

// NewClient initializes a "interceptors" service client given the endpoints.
func NewClient(get, create, stream goa.Endpoint, ci ClientInterceptors) *Client {
	return &Client{
		GetEndpoint:    WrapGetClientEndpoint(get, ci),
		CreateEndpoint: WrapCreateClientEndpoint(create, ci),
		StreamEndpoint: WrapStreamClientEndpoint(stream, ci),
	}
}

// Get calls the "get" endpoint of the "interceptors" service.
// Get may return the following errors:
//   - "NotFound" (type *goa.ServiceError): Record not found
//   - "Unavailable" (type *goa.ServiceError): Temporary error
//   - error: internal error
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *GetResult, err error) {
	var ires any
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetResult), nil
}

// Create calls the "create" endpoint of the "interceptors" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res *CreateResult, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateResult), nil
}

// Stream calls the "stream" endpoint of the "interceptors" service.
func (c *Client) Stream(ctx context.Context, p *StreamPayload) (res StreamClientStream, err error) {
	var ires any
	ires, err = c.StreamEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(StreamClientStream), nil
}

// Code generated by goa v3.17.1, DO NOT EDIT.
//
// concat client
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design

package concat

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "concat" service client.
type Client struct {
	ConcatEndpoint goa.Endpoint
}

// NewClient initializes a "concat" service client given the endpoints.
func NewClient(concat goa.Endpoint) *Client {
	return &Client{
		ConcatEndpoint: concat,
	}
}

// Concat calls the "concat" endpoint of the "concat" service.
func (c *Client) Concat(ctx context.Context, p *ConcatPayload) (res string, err error) {
	var ires any
	ires, err = c.ConcatEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

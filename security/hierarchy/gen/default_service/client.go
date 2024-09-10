// Code generated by goa v3.19.0, DO NOT EDIT.
//
// default_service client
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package defaultservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "default_service" service client.
type Client struct {
	DefaultEndpoint goa.Endpoint
}

// NewClient initializes a "default_service" service client given the endpoints.
func NewClient(default_ goa.Endpoint) *Client {
	return &Client{
		DefaultEndpoint: default_,
	}
}

// Default calls the "default" endpoint of the "default_service" service.
func (c *Client) Default(ctx context.Context, p *DefaultPayload) (err error) {
	_, err = c.DefaultEndpoint(ctx, p)
	return
}

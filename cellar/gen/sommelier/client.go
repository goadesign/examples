// Code generated by goa v2.2.0, DO NOT EDIT.
//
// sommelier client
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package sommelier

import (
	"context"

	"goa.design/goa"
)

// Client is the "sommelier" service client.
type Client struct {
	PickEndpoint goa.Endpoint
}

// NewClient initializes a "sommelier" service client given the endpoints.
func NewClient(pick goa.Endpoint) *Client {
	return &Client{
		PickEndpoint: pick,
	}
}

// Pick calls the "pick" endpoint of the "sommelier" service.
// Pick may return the following errors:
//	- "no_criteria" (type NoCriteria)
//	- "no_match" (type NoMatch)
//	- error: internal error
func (c *Client) Pick(ctx context.Context, p *Criteria) (res StoredBottleCollection, err error) {
	var ires interface{}
	ires, err = c.PickEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(StoredBottleCollection), nil
}

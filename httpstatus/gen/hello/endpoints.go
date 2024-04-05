// Code generated by goa v3.15.2, DO NOT EDIT.
//
// hello endpoints
//
// Command:
// $ goa gen goa.design/examples/httpstatus/design

package hello

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "hello" service endpoints.
type Endpoints struct {
	HelloEndpoint goa.Endpoint
}

// NewEndpoints wraps the methods of the "hello" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		HelloEndpoint: NewHelloEndpointEndpoint(s),
	}
}

// Use applies the given middleware to all the "hello" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.HelloEndpoint = m(e.HelloEndpoint)
}

// NewHelloEndpointEndpoint returns an endpoint function that calls the method
// "hello" of service "hello".
func NewHelloEndpointEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*HelloPayload)
		res, err := s.HelloEndpoint(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedHello(res, "default")
		return vres, nil
	}
}

// Code generated by goa v3.21.5, DO NOT EDIT.
//
// sommelier endpoints
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package sommelier

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "sommelier" service endpoints.
type Endpoints struct {
	Pick goa.Endpoint
}

// NewEndpoints wraps the methods of the "sommelier" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Pick: NewPickEndpoint(s),
	}
}

// Use applies the given middleware to all the "sommelier" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Pick = m(e.Pick)
}

// NewPickEndpoint returns an endpoint function that calls the method "pick" of
// service "sommelier".
func NewPickEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*Criteria)
		res, err := s.Pick(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredBottleCollection(res, "default")
		return vres, nil
	}
}

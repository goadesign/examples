// Code generated by goa v3.15.2, DO NOT EDIT.
//
// resume endpoints
//
// Command:
// $ goa gen goa.design/examples/multipart/design

package resume

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "resume" service endpoints.
type Endpoints struct {
	List goa.Endpoint
	Add  goa.Endpoint
}

// NewEndpoints wraps the methods of the "resume" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List: NewListEndpoint(s),
		Add:  NewAddEndpoint(s),
	}
}

// Use applies the given middleware to all the "resume" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Add = m(e.Add)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "resume".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredResumeCollection(res, "default")
		return vres, nil
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "resume".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.([]*Resume)
		return s.Add(ctx, p)
	}
}

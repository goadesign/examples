// Code generated by goa v3.21.5, DO NOT EDIT.
//
// default_service endpoints
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package defaultservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "default_service" service endpoints.
type Endpoints struct {
	Default goa.Endpoint
}

// NewEndpoints wraps the methods of the "default_service" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Default: NewDefaultEndpoint(s, a.BasicAuth),
	}
}

// Use applies the given middleware to all the "default_service" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Default = m(e.Default)
}

// NewDefaultEndpoint returns an endpoint function that calls the method
// "default" of service "default_service".
func NewDefaultEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DefaultPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.Username, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Default(ctx, p)
	}
}

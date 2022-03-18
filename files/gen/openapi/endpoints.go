// Code generated by goa v3.7.0, DO NOT EDIT.
//
// openapi endpoints
//
// Command:
// $ goa gen goa.design/examples/files/design -o
// $(GOPATH)/src/goa.design/examples/files

package openapi

import (
	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "openapi" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "openapi" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "openapi" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}

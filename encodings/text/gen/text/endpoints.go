// Code generated by goa v3.17.1, DO NOT EDIT.
//
// text endpoints
//
// Command:
// $ goa gen goa.design/examples/encodings/text/design

package text

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "text" service endpoints.
type Endpoints struct {
	Concatstrings     goa.Endpoint
	Concatbytes       goa.Endpoint
	Concatstringfield goa.Endpoint
	Concatbytesfield  goa.Endpoint
}

// NewEndpoints wraps the methods of the "text" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Concatstrings:     NewConcatstringsEndpoint(s),
		Concatbytes:       NewConcatbytesEndpoint(s),
		Concatstringfield: NewConcatstringfieldEndpoint(s),
		Concatbytesfield:  NewConcatbytesfieldEndpoint(s),
	}
}

// Use applies the given middleware to all the "text" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Concatstrings = m(e.Concatstrings)
	e.Concatbytes = m(e.Concatbytes)
	e.Concatstringfield = m(e.Concatstringfield)
	e.Concatbytesfield = m(e.Concatbytesfield)
}

// NewConcatstringsEndpoint returns an endpoint function that calls the method
// "concatstrings" of service "text".
func NewConcatstringsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ConcatstringsPayload)
		return s.Concatstrings(ctx, p)
	}
}

// NewConcatbytesEndpoint returns an endpoint function that calls the method
// "concatbytes" of service "text".
func NewConcatbytesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ConcatbytesPayload)
		return s.Concatbytes(ctx, p)
	}
}

// NewConcatstringfieldEndpoint returns an endpoint function that calls the
// method "concatstringfield" of service "text".
func NewConcatstringfieldEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ConcatstringfieldPayload)
		return s.Concatstringfield(ctx, p)
	}
}

// NewConcatbytesfieldEndpoint returns an endpoint function that calls the
// method "concatbytesfield" of service "text".
func NewConcatbytesfieldEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ConcatbytesfieldPayload)
		return s.Concatbytesfield(ctx, p)
	}
}

// Code generated by goa v3.21.5, DO NOT EDIT.
//
// sommelier HTTP server
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package server

import (
	"context"
	"net/http"

	sommelier "goa.design/examples/cellar/gen/sommelier"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the sommelier service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Pick   http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the sommelier service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *sommelier.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Pick", "POST", "/sommelier"},
		},
		Pick: NewPickHandler(e.Pick, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "sommelier" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Pick = m(s.Pick)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return sommelier.MethodNames[:] }

// Mount configures the mux to serve the sommelier endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountPickHandler(mux, h.Pick)
}

// Mount configures the mux to serve the sommelier endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountPickHandler configures the mux to serve the "sommelier" service "pick"
// endpoint.
func MountPickHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/sommelier", f)
}

// NewPickHandler creates a HTTP handler which loads the HTTP request and calls
// the "sommelier" service "pick" endpoint.
func NewPickHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePickRequest(mux, decoder)
		encodeResponse = EncodePickResponse(encoder)
		encodeError    = EncodePickError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "pick")
		ctx = context.WithValue(ctx, goa.ServiceKey, "sommelier")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil && errhandler != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil && errhandler != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			if errhandler != nil {
				errhandler(ctx, w, err)
			}
		}
	})
}

// Code generated by goa v3.17.2, DO NOT EDIT.
//
// tus HTTP server
//
// Command:
// $ goa gen goa.design/examples/tus/design

package server

import (
	"context"
	"net/http"

	tus "goa.design/examples/tus/gen/tus"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the tus service endpoint HTTP handlers.
type Server struct {
	Mounts  []*MountPoint
	Head    http.Handler
	Patch   http.Handler
	Options http.Handler
	Post    http.Handler
	Delete  http.Handler
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

// New instantiates HTTP handlers for all the tus service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *tus.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Head", "HEAD", "/upload/{id}"},
			{"Patch", "PATCH", "/upload/{id}"},
			{"Options", "OPTIONS", "/upload"},
			{"Post", "POST", "/upload"},
			{"Delete", "DELETE", "/upload/{id}"},
		},
		Head:    NewHeadHandler(e.Head, mux, decoder, encoder, errhandler, formatter),
		Patch:   NewPatchHandler(e.Patch, mux, decoder, encoder, errhandler, formatter),
		Options: NewOptionsHandler(e.Options, mux, decoder, encoder, errhandler, formatter),
		Post:    NewPostHandler(e.Post, mux, decoder, encoder, errhandler, formatter),
		Delete:  NewDeleteHandler(e.Delete, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "tus" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Head = m(s.Head)
	s.Patch = m(s.Patch)
	s.Options = m(s.Options)
	s.Post = m(s.Post)
	s.Delete = m(s.Delete)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return tus.MethodNames[:] }

// Mount configures the mux to serve the tus endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountHeadHandler(mux, h.Head)
	MountPatchHandler(mux, h.Patch)
	MountOptionsHandler(mux, h.Options)
	MountPostHandler(mux, h.Post)
	MountDeleteHandler(mux, h.Delete)
}

// Mount configures the mux to serve the tus endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountHeadHandler configures the mux to serve the "tus" service "head"
// endpoint.
func MountHeadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("HEAD", "/upload/{id}", f)
}

// NewHeadHandler creates a HTTP handler which loads the HTTP request and calls
// the "tus" service "head" endpoint.
func NewHeadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeHeadRequest(mux, decoder)
		encodeResponse = EncodeHeadResponse(encoder)
		encodeError    = EncodeHeadError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "head")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tus")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountPatchHandler configures the mux to serve the "tus" service "patch"
// endpoint.
func MountPatchHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/upload/{id}", f)
}

// NewPatchHandler creates a HTTP handler which loads the HTTP request and
// calls the "tus" service "patch" endpoint.
func NewPatchHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePatchRequest(mux, decoder)
		encodeResponse = EncodePatchResponse(encoder)
		encodeError    = EncodePatchError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "patch")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tus")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		data := &tus.PatchRequestData{Payload: payload.(*tus.PatchPayload), Body: r.Body}
		res, err := endpoint(ctx, data)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountOptionsHandler configures the mux to serve the "tus" service "options"
// endpoint.
func MountOptionsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/upload", f)
}

// NewOptionsHandler creates a HTTP handler which loads the HTTP request and
// calls the "tus" service "options" endpoint.
func NewOptionsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeOptionsResponse(encoder)
		encodeError    = EncodeOptionsError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "options")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tus")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountPostHandler configures the mux to serve the "tus" service "post"
// endpoint.
func MountPostHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/upload", f)
}

// NewPostHandler creates a HTTP handler which loads the HTTP request and calls
// the "tus" service "post" endpoint.
func NewPostHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePostRequest(mux, decoder)
		encodeResponse = EncodePostResponse(encoder)
		encodeError    = EncodePostError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "post")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tus")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		data := &tus.PostRequestData{Payload: payload.(*tus.PostPayload), Body: r.Body}
		res, err := endpoint(ctx, data)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteHandler configures the mux to serve the "tus" service "delete"
// endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/upload/{id}", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "tus" service "delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, decoder)
		encodeResponse = EncodeDeleteResponse(encoder)
		encodeError    = EncodeDeleteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tus")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

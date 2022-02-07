// Code generated by goa v3.5.5, DO NOT EDIT.
//
// updown HTTP server
//
// Command:
// $ goa gen goa.design/examples/upload_download/design -o
// $(GOPATH)/src/goa.design/examples/upload_download

package server

import (
	"context"
	"io"
	"net/http"

	updown "goa.design/examples/upload_download/gen/updown"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the updown service endpoint HTTP handlers.
type Server struct {
	Mounts   []*MountPoint
	Upload   http.Handler
	Download http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
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

// New instantiates HTTP handlers for all the updown service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *updown.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Upload", "POST", "/upload/{*dir}"},
			{"Download", "GET", "/download/{*filename}"},
		},
		Upload:   NewUploadHandler(e.Upload, mux, decoder, encoder, errhandler, formatter),
		Download: NewDownloadHandler(e.Download, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "updown" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Upload = m(s.Upload)
	s.Download = m(s.Download)
}

// Mount configures the mux to serve the updown endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUploadHandler(mux, h.Upload)
	MountDownloadHandler(mux, h.Download)
}

// Mount configures the mux to serve the updown endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountUploadHandler configures the mux to serve the "updown" service "upload"
// endpoint.
func MountUploadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/upload/{*dir}", f)
}

// NewUploadHandler creates a HTTP handler which loads the HTTP request and
// calls the "updown" service "upload" endpoint.
func NewUploadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUploadRequest(mux, decoder)
		encodeResponse = EncodeUploadResponse(encoder)
		encodeError    = EncodeUploadError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "upload")
		ctx = context.WithValue(ctx, goa.ServiceKey, "updown")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		data := &updown.UploadRequestData{Payload: payload.(*updown.UploadPayload), Body: r.Body}
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

// MountDownloadHandler configures the mux to serve the "updown" service
// "download" endpoint.
func MountDownloadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/download/{*filename}", f)
}

// NewDownloadHandler creates a HTTP handler which loads the HTTP request and
// calls the "updown" service "download" endpoint.
func NewDownloadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDownloadRequest(mux, decoder)
		encodeResponse = EncodeDownloadResponse(encoder)
		encodeError    = EncodeDownloadError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "download")
		ctx = context.WithValue(ctx, goa.ServiceKey, "updown")
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
		o := res.(*updown.DownloadResponseData)
		defer o.Body.Close()
		if err := encodeResponse(ctx, w, o.Result); err != nil {
			errhandler(ctx, w, err)
			return
		}
		if _, err := io.Copy(w, o.Body); err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
		}
	})
}

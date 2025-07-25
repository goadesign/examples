// Code generated by goa v3.21.5, DO NOT EDIT.
//
// storage HTTP server
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package server

import (
	"context"
	"mime/multipart"
	"net/http"

	storage "goa.design/examples/cellar/gen/storage"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the storage service endpoint HTTP handlers.
type Server struct {
	Mounts      []*MountPoint
	List        http.Handler
	Show        http.Handler
	Add         http.Handler
	Remove      http.Handler
	Rate        http.Handler
	MultiAdd    http.Handler
	MultiUpdate http.Handler
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

// StorageMultiAddDecoderFunc is the type to decode multipart request for the
// "storage" service "multi_add" endpoint.
type StorageMultiAddDecoderFunc func(*multipart.Reader, *[]*storage.Bottle) error

// StorageMultiUpdateDecoderFunc is the type to decode multipart request for
// the "storage" service "multi_update" endpoint.
type StorageMultiUpdateDecoderFunc func(*multipart.Reader, **storage.MultiUpdatePayload) error

// New instantiates HTTP handlers for all the storage service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *storage.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	storageMultiAddDecoderFn StorageMultiAddDecoderFunc,
	storageMultiUpdateDecoderFn StorageMultiUpdateDecoderFunc,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"List", "GET", "/storage"},
			{"Show", "GET", "/storage/{id}"},
			{"Add", "POST", "/storage"},
			{"Remove", "DELETE", "/storage/{id}"},
			{"Rate", "POST", "/storage/rate"},
			{"MultiAdd", "POST", "/storage/multi_add"},
			{"MultiUpdate", "PUT", "/storage/multi_update"},
		},
		List:        NewListHandler(e.List, mux, decoder, encoder, errhandler, formatter),
		Show:        NewShowHandler(e.Show, mux, decoder, encoder, errhandler, formatter),
		Add:         NewAddHandler(e.Add, mux, decoder, encoder, errhandler, formatter),
		Remove:      NewRemoveHandler(e.Remove, mux, decoder, encoder, errhandler, formatter),
		Rate:        NewRateHandler(e.Rate, mux, decoder, encoder, errhandler, formatter),
		MultiAdd:    NewMultiAddHandler(e.MultiAdd, mux, NewStorageMultiAddDecoder(mux, storageMultiAddDecoderFn), encoder, errhandler, formatter),
		MultiUpdate: NewMultiUpdateHandler(e.MultiUpdate, mux, NewStorageMultiUpdateDecoder(mux, storageMultiUpdateDecoderFn), encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "storage" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.List = m(s.List)
	s.Show = m(s.Show)
	s.Add = m(s.Add)
	s.Remove = m(s.Remove)
	s.Rate = m(s.Rate)
	s.MultiAdd = m(s.MultiAdd)
	s.MultiUpdate = m(s.MultiUpdate)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return storage.MethodNames[:] }

// Mount configures the mux to serve the storage endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListHandler(mux, h.List)
	MountShowHandler(mux, h.Show)
	MountAddHandler(mux, h.Add)
	MountRemoveHandler(mux, h.Remove)
	MountRateHandler(mux, h.Rate)
	MountMultiAddHandler(mux, h.MultiAdd)
	MountMultiUpdateHandler(mux, h.MultiUpdate)
}

// Mount configures the mux to serve the storage endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListHandler configures the mux to serve the "storage" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		var err error
		res, err := endpoint(ctx, nil)
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

// MountShowHandler configures the mux to serve the "storage" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/{id}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "show" endpoint.
func NewShowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowRequest(mux, decoder)
		encodeResponse = EncodeShowResponse(encoder)
		encodeError    = EncodeShowError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
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

// MountAddHandler configures the mux to serve the "storage" service "add"
// endpoint.
func MountAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage", f)
}

// NewAddHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "add" endpoint.
func NewAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddRequest(mux, decoder)
		encodeResponse = EncodeAddResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
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

// MountRemoveHandler configures the mux to serve the "storage" service
// "remove" endpoint.
func MountRemoveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/storage/{id}", f)
}

// NewRemoveHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "remove" endpoint.
func NewRemoveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveRequest(mux, decoder)
		encodeResponse = EncodeRemoveResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
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

// MountRateHandler configures the mux to serve the "storage" service "rate"
// endpoint.
func MountRateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/rate", f)
}

// NewRateHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "rate" endpoint.
func NewRateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRateRequest(mux, decoder)
		encodeResponse = EncodeRateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "rate")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
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

// MountMultiAddHandler configures the mux to serve the "storage" service
// "multi_add" endpoint.
func MountMultiAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/multi_add", f)
}

// NewMultiAddHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "multi_add" endpoint.
func NewMultiAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMultiAddRequest(mux, decoder)
		encodeResponse = EncodeMultiAddResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "multi_add")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
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

// MountMultiUpdateHandler configures the mux to serve the "storage" service
// "multi_update" endpoint.
func MountMultiUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/storage/multi_update", f)
}

// NewMultiUpdateHandler creates a HTTP handler which loads the HTTP request
// and calls the "storage" service "multi_update" endpoint.
func NewMultiUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMultiUpdateRequest(mux, decoder)
		encodeResponse = EncodeMultiUpdateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "multi_update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
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

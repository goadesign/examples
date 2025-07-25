// Code generated by goa v3.21.5, DO NOT EDIT.
//
// updown endpoints
//
// Command:
// $ goa gen goa.design/examples/upload_download/design

package updown

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "updown" service endpoints.
type Endpoints struct {
	Upload   goa.Endpoint
	Download goa.Endpoint
}

// UploadRequestData holds both the payload and the HTTP request body reader of
// the "upload" method.
type UploadRequestData struct {
	// Payload is the method payload.
	Payload *UploadPayload
	// Body streams the HTTP request body.
	Body io.ReadCloser
}

// DownloadResponseData holds both the result and the HTTP response body reader
// of the "download" method.
type DownloadResponseData struct {
	// Result is the method result.
	Result *DownloadResult
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "updown" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Upload:   NewUploadEndpoint(s),
		Download: NewDownloadEndpoint(s),
	}
}

// Use applies the given middleware to all the "updown" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Upload = m(e.Upload)
	e.Download = m(e.Download)
}

// NewUploadEndpoint returns an endpoint function that calls the method
// "upload" of service "updown".
func NewUploadEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		ep := req.(*UploadRequestData)
		return nil, s.Upload(ctx, ep.Payload, ep.Body)
	}
}

// NewDownloadEndpoint returns an endpoint function that calls the method
// "download" of service "updown".
func NewDownloadEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(string)
		res, body, err := s.Download(ctx, p)
		if err != nil {
			return nil, err
		}
		return &DownloadResponseData{Result: res, Body: body}, nil
	}
}

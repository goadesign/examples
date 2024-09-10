// Code generated by goa v3.15.2, DO NOT EDIT.
//
// hello HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/httpstatus/design

package server

import (
	"context"
	"net/http"

	helloviews "goa.design/examples/httpstatus/gen/hello/views"
	goahttp "goa.design/goa/v3/http"
)

// EncodeHelloEndpointResponse returns an encoder for responses returned by the
// hello hello endpoint.
func EncodeHelloEndpointResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*helloviews.Hello)
		if res.Projected.Outcome != nil && *res.Projected.Outcome == "created" {
			enc := encoder(ctx, w)
			body := NewHelloCreatedResponseBody(res.Projected)
			w.WriteHeader(http.StatusCreated)
			return enc.Encode(body)
		}
		if res.Projected.Outcome != nil && *res.Projected.Outcome == "accepted" {
			enc := encoder(ctx, w)
			body := NewHelloAcceptedResponseBody(res.Projected)
			w.WriteHeader(http.StatusAccepted)
			return enc.Encode(body)
		}
		enc := encoder(ctx, w)
		body := NewHelloOKResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeHelloEndpointRequest returns a decoder for requests sent to the hello
// hello endpoint.
func DecodeHelloEndpointRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			greeting string

			params = mux.Vars(r)
		)
		greeting = params["greeting"]
		payload := NewHelloPayload(greeting)

		return payload, nil
	}
}
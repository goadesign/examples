// Code generated by goa v3.21.5, DO NOT EDIT.
//
// interceptors HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	interceptors "goa.design/examples/interceptors/gen/interceptors"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetResponse returns an encoder for responses returned by the
// interceptors get endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*interceptors.GetResult)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the interceptors get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body GetRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateGetRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			tenantID string
			recordID string
			auth     string

			params = mux.Vars(r)
		)
		tenantID = params["tenantID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("tenantID", tenantID, goa.FormatUUID))
		recordID = params["recordID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("recordID", recordID, goa.FormatUUID))
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("auth", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetPayload(&body, tenantID, recordID, auth)

		return payload, nil
	}
}

// EncodeGetError returns an encoder for errors returned by the get
// interceptors endpoint.
func EncodeGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "Unavailable":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetUnavailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusServiceUnavailable)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateResponse returns an encoder for responses returned by the
// interceptors create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*interceptors.CreateResult)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the interceptors
// create endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			tenantID string
			auth     string

			params = mux.Vars(r)
		)
		tenantID = params["tenantID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("tenantID", tenantID, goa.FormatUUID))
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("auth", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreatePayload(&body, tenantID, auth)

		return payload, nil
	}
}

// DecodeStreamRequest returns a decoder for requests sent to the interceptors
// stream endpoint.
func DecodeStreamRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			tenantID string
			auth     string
			err      error

			params = mux.Vars(r)
		)
		tenantID = params["tenantID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("tenantID", tenantID, goa.FormatUUID))
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("auth", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewStreamPayload(tenantID, auth)

		return payload, nil
	}
}

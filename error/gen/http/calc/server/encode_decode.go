// Code generated by goa v3.14.0, DO NOT EDIT.
//
// calc HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	calc "goa.design/examples/error/gen/calc"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeDivideResponse returns an encoder for responses returned by the calc
// divide endpoint.
func EncodeDivideResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*calc.DivideResult)
		enc := encoder(ctx, w)
		body := NewDivideResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDivideRequest returns a decoder for requests sent to the calc divide
// endpoint.
func DecodeDivideRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body DivideRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateDivideRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewDividePayload(&body)

		return payload, nil
	}
}

// EncodeDivideError returns an encoder for errors returned by the divide calc
// endpoint.
func EncodeDivideError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "div_by_zero":
			var res *calc.DivByZero
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDivideDivByZeroResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "timeout":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDivideTimeoutResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusGatewayTimeout)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// Code generated by goa v3.0.10, DO NOT EDIT.
//
// updown HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/upload_download/design

package server

import (
	"context"
	"net/http"
	"strconv"

	updown "goa.design/examples/upload_download/gen/updown"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUploadResponse returns an encoder for responses returned by the updown
// upload endpoint.
func EncodeUploadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUploadRequest returns a decoder for requests sent to the updown upload
// endpoint.
func DecodeUploadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			name   string
			length uint
			err    error

			params = mux.Vars(r)
		)
		name = params["name"]
		{
			lengthRaw := r.Header.Get("Content-Length")
			if lengthRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Content-Length", "header"))
			}
			v, err2 := strconv.ParseUint(lengthRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("length", lengthRaw, "unsigned integer"))
			}
			length = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewUploadPayload(name, length)

		return payload, nil
	}
}

// EncodeDownloadResponse returns an encoder for responses returned by the
// updown download endpoint.
func EncodeDownloadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*updown.DownloadResult)
		val := res.Length
		lengths := strconv.FormatUint(uint64(val), 10)
		w.Header().Set("Content-Length", lengths)
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDownloadRequest returns a decoder for requests sent to the updown
// download endpoint.
func DecodeDownloadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			name string

			params = mux.Vars(r)
		)
		name = params["name"]
		payload := name

		return payload, nil
	}
}
// Code generated by goa v3.17.0, DO NOT EDIT.
//
// updown HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/upload_download/design

package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	updown "goa.design/examples/upload_download/gen/updown"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUploadResponse returns an encoder for responses returned by the updown
// upload endpoint.
func EncodeUploadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeUploadRequest returns a decoder for requests sent to the updown upload
// endpoint.
func DecodeUploadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			dir         string
			contentType string
			err         error

			params = mux.Vars(r)
		)
		dir = params["dir"]
		contentTypeRaw := r.Header.Get("Content-Type")
		if contentTypeRaw != "" {
			contentType = contentTypeRaw
		} else {
			contentType = "multipart/form-data; boundary=goa"
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("content_type", contentType, "multipart/[^;]+; boundary=.+"))
		if err != nil {
			return nil, err
		}
		payload := NewUploadPayload(dir, contentType)

		return payload, nil
	}
}

// EncodeUploadError returns an encoder for errors returned by the upload
// updown endpoint.
func EncodeUploadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_media_type":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadInvalidMediaTypeResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "invalid_multipart_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadInvalidMultipartRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "internal_error":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDownloadResponse returns an encoder for responses returned by the
// updown download endpoint.
func EncodeDownloadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*updown.DownloadResult)
		{
			val := res.Length
			lengths := strconv.FormatInt(val, 10)
			w.Header().Set("Content-Length", lengths)
		}
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDownloadRequest returns a decoder for requests sent to the updown
// download endpoint.
func DecodeDownloadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			filename string

			params = mux.Vars(r)
		)
		filename = params["filename"]
		payload := filename

		return payload, nil
	}
}

// EncodeDownloadError returns an encoder for errors returned by the download
// updown endpoint.
func EncodeDownloadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_file_path":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDownloadInvalidFilePathResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "internal_error":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDownloadInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

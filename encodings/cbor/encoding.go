package concatapi

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/fxamacker/cbor/v2"
	goahttp "goa.design/goa/v3/http"
)

// RequestDecoder returns a HTTP request body decoder. The resulting decoder
// uses the CBOR binary serialization standard if no content type is provided in
// the request or if the content type explicitly specifies "application/cbor".
// The implementation also supports decoding JSON, XML and GOB encoded requests
// when the content type header is set accordingly.
func RequestDecoder(r *http.Request) goahttp.Decoder {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" || contentType == "application/cbor" {
		// default to CBOR if not Content-Type provided
		return cbor.NewDecoder(r.Body)
	}
	// Degrade gracefully to support JSON/XML/GOB.
	return goahttp.RequestDecoder(r)
}

// ResponseEncoder always encodes the response with CBOR. See the default Goa
// ResponseEncoder for implementing content type negotiation if your service
// needs to support other types of response encodings.
func ResponseEncoder(ctx context.Context, w http.ResponseWriter) goahttp.Encoder {
	w.Header().Set("Content-Type", "application/cbor")
	return cbor.NewEncoder(w)
}

// RequestEncoder returns a HTTP request encoder that uses CBOR.
func RequestEncoder(r *http.Request) goahttp.Encoder {
	var buf bytes.Buffer
	r.Body = io.NopCloser(&buf)
	return cbor.NewEncoder(&buf)
}

// ResponseDecoder returns a HTTP response decoder that uses CBOR.
func ResponseDecoder(resp *http.Response) goahttp.Decoder {
	return cbor.NewDecoder(resp.Body)
}

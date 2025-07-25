// Code generated by goa v3.21.5, DO NOT EDIT.
//
// concat HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	concat "goa.design/examples/encodings/cbor/gen/concat"
	goahttp "goa.design/goa/v3/http"
)

// BuildConcatRequest instantiates a HTTP request object with method and path
// set to call the "concat" service "concat" endpoint
func (c *Client) BuildConcatRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		a string
		b string
	)
	{
		p, ok := v.(*concat.ConcatPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("concat", "concat", "*concat.ConcatPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConcatConcatPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("concat", "concat", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeConcatResponse returns a decoder for responses returned by the concat
// concat endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeConcatResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("concat", "concat", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("concat", "concat", resp.StatusCode, string(body))
		}
	}
}

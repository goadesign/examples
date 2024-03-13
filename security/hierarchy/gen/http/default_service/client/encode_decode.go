// Code generated by goa v3.15.2, DO NOT EDIT.
//
// default_service HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	defaultservice "goa.design/examples/security/hierarchy/gen/default_service"
	goahttp "goa.design/goa/v3/http"
)

// BuildDefaultRequest instantiates a HTTP request object with method and path
// set to call the "default_service" service "default" endpoint
func (c *Client) BuildDefaultRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DefaultDefaultServicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("default_service", "default", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDefaultRequest returns an encoder for requests sent to the
// default_service default server.
func EncodeDefaultRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*defaultservice.DefaultPayload)
		if !ok {
			return goahttp.ErrInvalidType("default_service", "default", "*defaultservice.DefaultPayload", v)
		}
		req.SetBasicAuth(p.Username, p.Password)
		return nil
	}
}

// DecodeDefaultResponse returns a decoder for responses returned by the
// default_service default endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeDefaultResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("default_service", "default", resp.StatusCode, string(body))
		}
	}
}

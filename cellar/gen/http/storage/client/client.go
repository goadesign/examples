// Code generated by goa v3.21.5, DO NOT EDIT.
//
// storage client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package client

import (
	"context"
	"mime/multipart"
	"net/http"

	storage "goa.design/examples/cellar/gen/storage"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the storage service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Remove Doer is the HTTP client used to make requests to the remove endpoint.
	RemoveDoer goahttp.Doer

	// Rate Doer is the HTTP client used to make requests to the rate endpoint.
	RateDoer goahttp.Doer

	// MultiAdd Doer is the HTTP client used to make requests to the multi_add
	// endpoint.
	MultiAddDoer goahttp.Doer

	// MultiUpdate Doer is the HTTP client used to make requests to the
	// multi_update endpoint.
	MultiUpdateDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// StorageMultiAddEncoderFunc is the type to encode multipart request for the
// "storage" service "multi_add" endpoint.
type StorageMultiAddEncoderFunc func(*multipart.Writer, []*storage.Bottle) error

// StorageMultiUpdateEncoderFunc is the type to encode multipart request for
// the "storage" service "multi_update" endpoint.
type StorageMultiUpdateEncoderFunc func(*multipart.Writer, *storage.MultiUpdatePayload) error

// NewClient instantiates HTTP clients for all the storage service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListDoer:            doer,
		ShowDoer:            doer,
		AddDoer:             doer,
		RemoveDoer:          doer,
		RateDoer:            doer,
		MultiAddDoer:        doer,
		MultiUpdateDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// List returns an endpoint that makes HTTP requests to the storage service
// list server.
func (c *Client) List() goa.Endpoint {
	var (
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the storage service
// show server.
func (c *Client) Show() goa.Endpoint {
	var (
		encodeRequest  = EncodeShowRequest(c.encoder)
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "show", err)
		}
		return decodeResponse(resp)
	}
}

// Add returns an endpoint that makes HTTP requests to the storage service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddRequest(c.encoder)
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Remove returns an endpoint that makes HTTP requests to the storage service
// remove server.
func (c *Client) Remove() goa.Endpoint {
	var (
		decodeResponse = DecodeRemoveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRemoveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "remove", err)
		}
		return decodeResponse(resp)
	}
}

// Rate returns an endpoint that makes HTTP requests to the storage service
// rate server.
func (c *Client) Rate() goa.Endpoint {
	var (
		encodeRequest  = EncodeRateRequest(c.encoder)
		decodeResponse = DecodeRateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "rate", err)
		}
		return decodeResponse(resp)
	}
}

// MultiAdd returns an endpoint that makes HTTP requests to the storage service
// multi_add server.
func (c *Client) MultiAdd(storageMultiAddEncoderFn StorageMultiAddEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeMultiAddRequest(NewStorageMultiAddEncoder(storageMultiAddEncoderFn))
		decodeResponse = DecodeMultiAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMultiAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MultiAddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "multi_add", err)
		}
		return decodeResponse(resp)
	}
}

// MultiUpdate returns an endpoint that makes HTTP requests to the storage
// service multi_update server.
func (c *Client) MultiUpdate(storageMultiUpdateEncoderFn StorageMultiUpdateEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeMultiUpdateRequest(NewStorageMultiUpdateEncoder(storageMultiUpdateEncoderFn))
		decodeResponse = DecodeMultiUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMultiUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MultiUpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "multi_update", err)
		}
		return decodeResponse(resp)
	}
}

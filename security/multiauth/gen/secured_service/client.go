// Code generated by goa v3.14.0, DO NOT EDIT.
//
// secured_service client
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design

package securedservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "secured_service" service client.
type Client struct {
	SigninEndpoint           goa.Endpoint
	SecureEndpoint           goa.Endpoint
	DoublySecureEndpoint     goa.Endpoint
	AlsoDoublySecureEndpoint goa.Endpoint
}

// NewClient initializes a "secured_service" service client given the endpoints.
func NewClient(signin, secure, doublySecure, alsoDoublySecure goa.Endpoint) *Client {
	return &Client{
		SigninEndpoint:           signin,
		SecureEndpoint:           secure,
		DoublySecureEndpoint:     doublySecure,
		AlsoDoublySecureEndpoint: alsoDoublySecure,
	}
}

// Signin calls the "signin" endpoint of the "secured_service" service.
// Signin may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - error: internal error
func (c *Client) Signin(ctx context.Context, p *SigninPayload) (res *Creds, err error) {
	var ires any
	ires, err = c.SigninEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Creds), nil
}

// Secure calls the "secure" endpoint of the "secured_service" service.
// Secure may return the following errors:
//   - "invalid-scopes" (type InvalidScopes)
//   - "unauthorized" (type Unauthorized)
//   - error: internal error
func (c *Client) Secure(ctx context.Context, p *SecurePayload) (res string, err error) {
	var ires any
	ires, err = c.SecureEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// DoublySecure calls the "doubly_secure" endpoint of the "secured_service"
// service.
// DoublySecure may return the following errors:
//   - "invalid-scopes" (type InvalidScopes)
//   - "unauthorized" (type Unauthorized)
//   - error: internal error
func (c *Client) DoublySecure(ctx context.Context, p *DoublySecurePayload) (res string, err error) {
	var ires any
	ires, err = c.DoublySecureEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// AlsoDoublySecure calls the "also_doubly_secure" endpoint of the
// "secured_service" service.
// AlsoDoublySecure may return the following errors:
//   - "invalid-scopes" (type InvalidScopes)
//   - "unauthorized" (type Unauthorized)
//   - error: internal error
func (c *Client) AlsoDoublySecure(ctx context.Context, p *AlsoDoublySecurePayload) (res string, err error) {
	var ires any
	ires, err = c.AlsoDoublySecureEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

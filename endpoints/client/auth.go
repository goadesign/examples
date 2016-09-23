package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// BasicAuthPath computes a request path to the basic action of auth.
func BasicAuthPath() string {
	return fmt.Sprintf("/auth/basic")
}

// BasicAuth makes a request to the basic action endpoint of the auth resource
func (c *Client) BasicAuth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewBasicAuthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewBasicAuthRequest create the request corresponding to the basic action endpoint of the auth resource.
func (c *Client) NewBasicAuthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.APIKeySigner != nil {
		c.APIKeySigner.Sign(req)
	}
	return req, nil
}

// JWTAuthPath computes a request path to the jwt action of auth.
func JWTAuthPath() string {
	return fmt.Sprintf("/auth/jwt")
}

// JWTAuth makes a request to the jwt action endpoint of the auth resource
func (c *Client) JWTAuth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewJWTAuthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewJWTAuthRequest create the request corresponding to the jwt action endpoint of the auth resource.
func (c *Client) NewJWTAuthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.GoogleJWTSigner != nil {
		c.GoogleJWTSigner.Sign(req)
	}
	return req, nil
}

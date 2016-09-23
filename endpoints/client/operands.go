package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AddOperandsPath computes a request path to the add action of operands.
func AddOperandsPath(left int, right int) string {
	return fmt.Sprintf("/add/%v/%v", left, right)
}

// add returns the sum of the left and right parameters in the response body
func (c *Client) AddOperands(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAddOperandsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddOperandsRequest create the request corresponding to the add action endpoint of the operands resource.
func (c *Client) NewAddOperandsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
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

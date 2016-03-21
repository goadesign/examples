package client

import (
	goaclient "github.com/goadesign/goa/client"
	"net/http"
)

// Client is the adder service client.
type Client struct {
	*goaclient.Client
}

// New instantiates the client.
func New(c *http.Client) *Client {
	return &Client{Client: goaclient.New(c)}
}

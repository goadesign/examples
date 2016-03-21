package client

import (
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"net/url"
)

// echo websocket server
func (c *Client) ConnectEcho(ctx context.Context, path string, initial string) (*websocket.Conn, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "ws"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("initial", initial)
	u.RawQuery = values.Encode()
	return websocket.Dial(u.String(), "", u.String())
}

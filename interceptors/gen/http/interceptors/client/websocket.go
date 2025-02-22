// Code generated by goa v3.19.1, DO NOT EDIT.
//
// interceptors WebSocket client streaming
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package client

import (
	"context"
	"io"

	"github.com/gorilla/websocket"
	interceptors "goa.design/examples/interceptors/gen/interceptors"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "interceptors" service.
type ConnConfigurer struct {
	StreamFn goahttp.ConnConfigureFunc
}

// StreamClientStream implements the interceptors.StreamClientStream interface.
type StreamClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "interceptors" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		StreamFn: fn,
	}
}

// Recv reads instances of "interceptors.StreamResult" from the "stream"
// endpoint websocket connection.
func (s *StreamClientStream) Recv() (*interceptors.StreamResult, error) {
	var (
		rv   *interceptors.StreamResult
		body StreamResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	err = ValidateStreamResponseBody(&body)
	if err != nil {
		return rv, err
	}
	res := NewStreamResultOK(&body)
	return res, nil
}

// RecvWithContext reads instances of "interceptors.StreamResult" from the
// "stream" endpoint websocket connection with context.
func (s *StreamClientStream) RecvWithContext(ctx context.Context) (*interceptors.StreamResult, error) {
	return s.Recv()
}

// Send streams instances of "interceptors.StreamStreamingPayload" to the
// "stream" endpoint websocket connection.
func (s *StreamClientStream) Send(v *interceptors.StreamStreamingPayload) error {
	body := NewStreamStreamingBody(v)
	return s.conn.WriteJSON(body)
}

// SendWithContext streams instances of "interceptors.StreamStreamingPayload"
// to the "stream" endpoint websocket connection with context.
func (s *StreamClientStream) SendWithContext(ctx context.Context, v *interceptors.StreamStreamingPayload) error {
	return s.Send(v)
}

// Close closes the "stream" endpoint websocket connection.
func (s *StreamClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}

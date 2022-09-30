// Code generated by goa v3.8.7, DO NOT EDIT.
//
// chatter WebSocket client streaming
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package client

import (
	"io"

	"github.com/gorilla/websocket"
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "chatter" service.
type ConnConfigurer struct {
	EchoerFn    goahttp.ConnConfigureFunc
	ListenerFn  goahttp.ConnConfigureFunc
	SummaryFn   goahttp.ConnConfigureFunc
	SubscribeFn goahttp.ConnConfigureFunc
	HistoryFn   goahttp.ConnConfigureFunc
}

// EchoerClientStream implements the chatter.EchoerClientStream interface.
type EchoerClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// ListenerClientStream implements the chatter.ListenerClientStream interface.
type ListenerClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SummaryClientStream implements the chatter.SummaryClientStream interface.
type SummaryClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SubscribeClientStream implements the chatter.SubscribeClientStream interface.
type SubscribeClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// HistoryClientStream implements the chatter.HistoryClientStream interface.
type HistoryClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
	// view is the view to render  result type before sending to the websocket
	// connection.
	view string
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "chatter" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		EchoerFn:    fn,
		ListenerFn:  fn,
		SummaryFn:   fn,
		SubscribeFn: fn,
		HistoryFn:   fn,
	}
}

// Recv reads instances of "string" from the "echoer" endpoint websocket
// connection.
func (s *EchoerClientStream) Recv() (string, error) {
	var (
		rv   string
		body string
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}

// Send streams instances of "string" to the "echoer" endpoint websocket
// connection.
func (s *EchoerClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Close closes the "echoer" endpoint websocket connection.
func (s *EchoerClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}

// Send streams instances of "string" to the "listener" endpoint websocket
// connection.
func (s *ListenerClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Close closes the "listener" endpoint websocket connection.
func (s *ListenerClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}

// CloseAndRecv stops sending messages to the "summary" endpoint websocket
// connection and reads instances of "chatter.ChatSummaryCollection" from the
// connection.
func (s *SummaryClientStream) CloseAndRecv() (chatter.ChatSummaryCollection, error) {
	var (
		rv   chatter.ChatSummaryCollection
		body SummaryResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewSummaryChatSummaryCollectionOK(body)
	vres := chatterviews.ChatSummaryCollection{res, "default"}
	if err := chatterviews.ValidateChatSummaryCollection(vres); err != nil {
		return rv, goahttp.ErrValidationError("chatter", "summary", err)
	}
	return chatter.NewChatSummaryCollection(vres), nil
}

// Send streams instances of "string" to the "summary" endpoint websocket
// connection.
func (s *SummaryClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Recv reads instances of "chatter.Event" from the "subscribe" endpoint
// websocket connection.
func (s *SubscribeClientStream) Recv() (*chatter.Event, error) {
	var (
		rv   *chatter.Event
		body SubscribeResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	err = ValidateSubscribeResponseBody(&body)
	if err != nil {
		return rv, err
	}
	res := NewSubscribeEventOK(&body)
	return res, nil
}

// Recv reads instances of "chatter.ChatSummary" from the "history" endpoint
// websocket connection.
func (s *HistoryClientStream) Recv() (*chatter.ChatSummary, error) {
	var (
		rv   *chatter.ChatSummary
		body HistoryResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewHistoryChatSummaryOK(&body)
	vres := &chatterviews.ChatSummary{res, s.view}
	if err := chatterviews.ValidateChatSummary(vres); err != nil {
		return rv, goahttp.ErrValidationError("chatter", "history", err)
	}
	return chatter.NewChatSummary(vres), nil
}

// SetView sets the view to render the  type before sending to the "history"
// endpoint websocket connection.
func (s *HistoryClientStream) SetView(view string) {
	s.view = view
}

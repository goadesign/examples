// Code generated by goa v3.14.2, DO NOT EDIT.
//
// chatter WebSocket server streaming
//
// Command:
// $ goa gen goa.design/examples/streaming/design

package server

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	chatter "goa.design/examples/streaming/gen/chatter"
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

// EchoerServerStream implements the chatter.EchoerServerStream interface.
type EchoerServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// ListenerServerStream implements the chatter.ListenerServerStream interface.
type ListenerServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SummaryServerStream implements the chatter.SummaryServerStream interface.
type SummaryServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SubscribeServerStream implements the chatter.SubscribeServerStream interface.
type SubscribeServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// HistoryServerStream implements the chatter.HistoryServerStream interface.
type HistoryServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
	// view is the view to render chatter.ChatSummary result type before sending to
	// the websocket connection.
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

// Send streams instances of "string" to the "echoer" endpoint websocket
// connection.
func (s *EchoerServerStream) Send(v string) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}

// Recv reads instances of "string" from the "echoer" endpoint websocket
// connection.
func (s *EchoerServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}

// Close closes the "echoer" endpoint websocket connection.
func (s *EchoerServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// Recv reads instances of "string" from the "listener" endpoint websocket
// connection.
func (s *ListenerServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}

// Close closes the "listener" endpoint websocket connection.
func (s *ListenerServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// SendAndClose streams instances of "chatter.ChatSummaryCollection" to the
// "summary" endpoint websocket connection and closes the connection.
func (s *SummaryServerStream) SendAndClose(v chatter.ChatSummaryCollection) error {
	defer s.conn.Close()
	res := chatter.NewViewedChatSummaryCollection(v, "default")
	body := NewChatSummaryResponseCollection(res.Projected)
	return s.conn.WriteJSON(body)
}

// Recv reads instances of "string" from the "summary" endpoint websocket
// connection.
func (s *SummaryServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}

// Send streams instances of "chatter.Event" to the "subscribe" endpoint
// websocket connection.
func (s *SubscribeServerStream) Send(v *chatter.Event) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewSubscribeResponseBody(res)
	return s.conn.WriteJSON(body)
}

// Close closes the "subscribe" endpoint websocket connection.
func (s *SubscribeServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// Send streams instances of "chatter.ChatSummary" to the "history" endpoint
// websocket connection.
func (s *HistoryServerStream) Send(v *chatter.ChatSummary) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		respHdr := make(http.Header)
		respHdr.Add("goa-view", s.view)
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := chatter.NewViewedChatSummary(v, s.view)
	var body any
	switch s.view {
	case "tiny":
		body = NewHistoryResponseBodyTiny(res.Projected)
	case "default", "":
		body = NewHistoryResponseBody(res.Projected)
	}
	return s.conn.WriteJSON(body)
}

// Close closes the "history" endpoint websocket connection.
func (s *HistoryServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// SetView sets the view to render the chatter.ChatSummary type before sending
// to the "history" endpoint websocket connection.
func (s *HistoryServerStream) SetView(view string) {
	s.view = view
}

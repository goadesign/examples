package events

import (
	"context"
	"io"
	"net/http"
	"time"
)

// SSEReadCloser implements io.ReadCloser and io.WriterTo for SSE streaming
type SSEReadCloser struct {
	broker   *EventBroker
	clientID string
	channel  chan Event
	ctx      context.Context
	cancel   context.CancelFunc
	buffer   string
}

// NewSSEReadCloser creates a new SSE reader
func NewSSEReadCloser(ctx context.Context, broker *EventBroker) *SSEReadCloser {
	ctx, cancel := context.WithCancel(ctx)
	clientID, channel := broker.Subscribe()

	return &SSEReadCloser{
		broker:   broker,
		clientID: clientID,
		channel:  channel,
		ctx:      ctx,
		cancel:   cancel,
		buffer:   "",
	}
}

// Read implements io.Reader for SSE
func (s *SSEReadCloser) Read(p []byte) (n int, err error) {
	// If we have data in the buffer, return it
	if len(s.buffer) > 0 {
		n = copy(p, s.buffer)
		s.buffer = s.buffer[n:]
		return n, nil
	}

	// Wait for the next event or context cancellation
	select {
	case event, ok := <-s.channel:
		if !ok {
			return 0, io.EOF
		}

		// Format the event and put it in our buffer
		s.buffer = event.String()

		// Copy as much as we can to p
		n = copy(p, s.buffer)
		s.buffer = s.buffer[n:]
		return n, nil

	case <-s.ctx.Done():
		return 0, io.EOF

	case <-time.After(15 * time.Second):
		// Send a comment as a keepalive
		s.buffer = ": keepalive \n\n"
		n = copy(p, s.buffer)
		s.buffer = s.buffer[n:]
		return n, nil
	}
}

// Close implements io.Closer for SSE
func (s *SSEReadCloser) Close() error {
	s.cancel()
	s.broker.Unsubscribe(s.clientID)
	return nil
}

// WriteTo implements io.WriterTo for SSE
// This allows us to use the HTTP Flusher interface for immediate event delivery
func (s *SSEReadCloser) WriteTo(w io.Writer) (n int64, err error) {
	// Check if the writer supports flushing
	flusher, canFlush := w.(http.Flusher)

	// Send initial response headers if this is an HTTP response writer
	if httpW, ok := w.(http.ResponseWriter); ok {
		httpW.Header().Set("Content-Type", "text/event-stream")
		httpW.Header().Set("Cache-Control", "no-cache")
		httpW.Header().Set("Connection", "keep-alive")
		httpW.Header().Set("Access-Control-Allow-Origin", "*")

		// Flush headers immediately
		if canFlush {
			flusher.Flush()
		}
	}

	// Buffer for reading events
	buf := make([]byte, 4096)
	var total int64

	for {
		// Check if context is done
		select {
		case <-s.ctx.Done():
			return total, nil
		default:
			// Continue processing
		}

		// Read event data
		count, err := s.Read(buf)
		if err != nil && err != io.EOF {
			return total, err
		}

		// Write to output
		if count > 0 {
			written, err := w.Write(buf[:count])
			if err != nil {
				return total, err
			}
			total += int64(written)

			// Flush immediately if possible
			if canFlush {
				flusher.Flush()
			}
		}

		// Exit if we've reached EOF
		if err == io.EOF {
			return total, nil
		}
	}
}

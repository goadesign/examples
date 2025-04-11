package events

import (
	"context"
	"embed"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	genevents "goa.design/examples/sse/gen/events"
)

// events service implementation with SSE support
type Service struct {
	broker *EventBroker
}

//go:embed public
var StaticFS embed.FS

// New returns the events service implementation
func New() genevents.Service {
	return &Service{
		broker: NewEventBroker(100), // Keep last 100 events
	}
}

// Stream implements the SSE endpoint
func (s *Service) Stream(ctx context.Context) (resp io.ReadCloser, err error) {
	// Create a new SSE stream
	resp = NewSSEReadCloser(ctx, s.broker)

	// Send a welcome event
	welcomeData := fmt.Sprintf("{\"message\":\"Connected to SSE stream\",\"clients\":%d}", s.broker.ClientCount())

	s.broker.Broadcast(Event{
		ID:        uuid.New().String(),
		Event:     "connected",
		Data:      welcomeData,
		Retry:     3000, // Reconnect after 3 seconds if connection is lost
		Timestamp: time.Now(),
	})

	return
}

// Trigger broadcasts a new event to all connected clients
func (s *Service) Trigger(ctx context.Context, p *genevents.TriggerPayload) (res *genevents.TriggerResult, err error) {
	// Create and broadcast the event
	s.broker.Broadcast(Event{
		ID:        uuid.New().String(),
		Event:     "message",
		Data:      fmt.Sprintf("{\"message\":\"%s\"}", p.Message),
		Timestamp: time.Now(),
	})

	res = &genevents.TriggerResult{
		Status: fmt.Sprintf("Event broadcast to %d clients", s.broker.ClientCount()),
	}

	return
}

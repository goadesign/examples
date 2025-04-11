package events

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// EventBroker manages SSE connections and broadcasts events to clients
type EventBroker struct {
	mutex      sync.RWMutex
	clients    map[string]chan Event
	events     []Event // Store recent events for replay
	maxHistory int
}

// Event represents a single SSE event
type Event struct {
	ID        string
	Event     string
	Data      string
	Retry     int
	Timestamp time.Time
}

// NewEventBroker creates a new event broker
func NewEventBroker(maxHistory int) *EventBroker {
	return &EventBroker{
		clients:    make(map[string]chan Event),
		events:     make([]Event, 0, maxHistory),
		maxHistory: maxHistory,
	}
}

// Subscribe adds a new client and returns a channel for events
func (b *EventBroker) Subscribe() (string, chan Event) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	clientID := uuid.New().String()
	channel := make(chan Event, 100) // Buffer to prevent blocking

	// Add client to the broker
	b.clients[clientID] = channel

	// Send recent events to the new client
	go func() {
		for _, event := range b.events {
			channel <- event
		}
	}()

	return clientID, channel
}

// Unsubscribe removes a client
func (b *EventBroker) Unsubscribe(clientID string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if channel, exists := b.clients[clientID]; exists {
		close(channel)
		delete(b.clients, clientID)
	}
}

// Broadcast sends an event to all connected clients
func (b *EventBroker) Broadcast(event Event) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// Store the event in history
	b.events = append(b.events, event)
	if len(b.events) > b.maxHistory {
		b.events = b.events[1:]
	}

	// Send to all clients
	successCount := 0
	dropCount := 0
	for _, channel := range b.clients {
		select {
		case channel <- event:
			// Event sent successfully
			successCount++
		default:
			// Channel buffer is full, we could handle this differently
			// For now, we'll just drop the event for this client
			dropCount++
		}
	}
}

// ClientCount returns the number of connected clients
func (b *EventBroker) ClientCount() int {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return len(b.clients)
}

// String formats the event according to the SSE spec
func (e Event) String() string {
	var sb strings.Builder

	if e.ID != "" {
		sb.WriteString(fmt.Sprintf("id: %s\n", e.ID))
	}

	if e.Event != "" {
		sb.WriteString(fmt.Sprintf("event: %s\n", e.Event))
	}

	// For multiline data, each line needs to be prefixed with "data: "
	// For now, we're just sending a single line
	sb.WriteString(fmt.Sprintf("data: %s\n", e.Data))

	if e.Retry > 0 {
		sb.WriteString(fmt.Sprintf("retry: %d\n", e.Retry))
	}

	sb.WriteString("\n") // End event with an extra newline
	return sb.String()
}

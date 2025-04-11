# Server-Sent Events (SSE) Example

This example demonstrates how to implement Server-Sent Events (SSE) in a Goa
service. SSE is a server push technology enabling a client to receive automatic
updates from a server via an HTTP connection. Unlike WebSockets, SSE is a
one-way communication channel (server to client only).

## Overview

This example includes:

1. A Goa service with two endpoints:
   - `stream`: An SSE endpoint that clients connect to for receiving real-time events
   - `trigger`: An endpoint to manually broadcast messages to all connected clients

2. A web-based client interface for testing the SSE functionality

3. An event broker that manages client connections and broadcasts events

## Running the Example

### Build and Start the Server

```bash
# Navigate to the example directory
cd /Users/raphael/go/src/goa.design/examples/sse

# Build the server
go build goa.design/examples/sse/cmd/events -o ./cmd/events/events

# Run the server
./cmd/events/events --http-port 8080
```

### Open the Web Client

Once the server is running, open your browser and navigate to:

```
http://localhost:8080
```

You'll see a web interface with the following features:
- A "Connect" button to establish an SSE connection
- A text input to send custom messages
- An event log showing all received events

### Testing the SSE Connection

1. Click the "Connect" button to establish an SSE connection
2. You should see a "Connection established" message in the events panel
3. Enter a message in the text field and click "Send" to broadcast it
4. The message will appear in the events panel of all connected clients

### Using the CLI to Trigger Events

You can also use the provided CLI tool to trigger events:

```bash
# Build the CLI tool
go build goa.design/examples/sse/cmd/events-cli -o ./cmd/events-cli/events-cli

# Trigger an event
./cmd/events-cli/events-cli events trigger --message "Hello from CLI!"
```

## Implementation Details

### Service Design

The service is defined in `design/design.go` with two main endpoints:

```go
Method("stream", func() {
    Description("Stream events using Server-Sent Events")

    HTTP(func() {
        GET("/events/stream")
        // This is crucial for SSE - skip response body encode/decode
        SkipResponseBodyEncodeDecode()
        Response(func() {
            ContentType("text/event-stream")
        })
    })
})

Method("trigger", func() {
    Description("Trigger a new event to be sent to all connected clients")

    Payload(func() {
        Attribute("message", String, "Message to broadcast")
        Required("message")
    })

    Result(func() {
        Attribute("status", String, "Status of the operation")
        Required("status")
    })

    HTTP(func() {
        POST("/events/trigger")
    })
})
```

### Key Components

1. **EventBroker**: Manages client connections and broadcasts events to all connected clients
   - Keeps track of connected clients
   - Maintains a history of recent events
   - Handles subscribing and unsubscribing clients

2. **SSEReadCloser**: Implements `io.ReadCloser` for SSE streaming
   - Formats events according to the SSE specification
   - Handles keepalive messages
   - Manages the connection lifecycle

3. **Service Implementation**: Connects the Goa endpoints with the SSE functionality
   - `Stream`: Establishes an SSE connection and sends events to the client
   - `Trigger`: Broadcasts new events to all connected clients

### SSE Format

The implementation follows the [HTML specification for Server-Sent Events](https://html.spec.whatwg.org/multipage/server-sent-events.html), including:

```
id: <event-id>
event: <event-type>
data: <event-data>
retry: <reconnection-time-in-milliseconds>

```

### Client-Side Implementation

The web client uses the standard `EventSource` API to connect to the SSE stream:

```javascript
eventSource = new EventSource('/events/stream');

// Listen for specific event types
eventSource.addEventListener('connected', handleConnectedEvent);
eventSource.addEventListener('message', handleMessageEvent);

// Handle errors
eventSource.onerror = handleError;
```

## Benefits of SSE

- Simpler than WebSockets for one-way communication
- Automatic reconnection
- Event IDs allow resuming from where the client left off
- Works over standard HTTP/HTTPS
- Native browser support via the EventSource API

package design

import (
	_ "embed"

	. "goa.design/goa/v3/dsl"
)

var _ = Service("events", func() {
	Description("The events service provides real-time updates via SSE")

	// Define the events endpoint that will stream data
	Method("stream", func() {
		Description("Stream events using Server-Sent Events")

		// No request payload needed
		HTTP(func() {
			GET("/events/stream")
			// This is crucial for SSE - skip response body encode/decode
			SkipResponseBodyEncodeDecode()
			Response(func() {
				ContentType("text/event-stream")
			})
		})
	})

	// Add a simple endpoint to trigger events for demonstration
	Method("trigger", func() {
		Description("Trigger a new event to be sent to all connected clients")

		Payload(func() {
			Attribute("message", String, "Message to broadcast", func() {
				Example("Something happened!")
			})
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

	// Serve static files
	Files("/{*filepath}", "public")
})

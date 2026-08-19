// Package design defines a service whose generated clients safely retry a
// temporary failure over HTTP and gRPC.
package design

import . "goa.design/goa/v3/dsl"

var _ = API("retry", func() {
	Title("Retry Service")
	Description("Demonstrates transport-independent retries for an idempotent method.")

	Server("retry", func() {
		Description("retry hosts the retry example service.")
		Services("retry")
		Host("development", func() {
			URI("http://localhost:18080")
			URI("grpc://localhost:18081")
		})
	})
})

var _ = Service("retry", func() {
	Description("Returns a message after demonstrating one safe retry.")

	Method("get_message", func() {
		Description("Returns a stable message after a temporary first-attempt failure.")
		Idempotent()
		Payload(func() {
			Field(1, "id", String, "Identifies one demonstration request.")
			Required("id")
		})
		Result(func() {
			Field(1, "message", String, "Message returned after the retry succeeds.")
			Required("message")
		})
		Error("unavailable", func() {
			Description("The message is temporarily unavailable.")
			Temporary()
		})

		HTTP(func() {
			GET("/messages/{id}")
			Response(StatusOK)
			Response("unavailable", StatusServiceUnavailable)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("unavailable", CodeUnavailable)
		})
	})
})

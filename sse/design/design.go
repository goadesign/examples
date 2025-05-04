package design

import (
	_ "embed"

	. "goa.design/goa/v3/dsl"
)

var _ = Service("monitor", func() {
	Description("The monitor service provides real-time system monitoring via SSE")

	Method("monitor", func() {
		Description("Monitor CPU and memory usage via Server-Sent Events")
		StreamingResult(Usage)
		HTTP(func() {
			GET("/monitor")
			ServerSentEvents()
		})
	})

	Files("/{*filepath}", "public")
})

var Usage = Type("Usage", func() {
	Attribute("timestamp", String, "Timestamp of the usage report", func() {
		Format(FormatDateTime)
	})
	Attribute("cpu", Int, "CPU usage percentage", func() {
		Example(50)
	})
	Attribute("memory", Int, "Memory usage percentage", func() {
		Example(30)
	})
	Required("timestamp", "cpu", "memory")
})

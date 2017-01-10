package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("archiver", func() {
	Title("The goa AWS X-Ray example downstream service")
	Description("Archiver is a service that manages the content of HTTP responses")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("archiver", func() {
	CanonicalActionName("read")
	BasePath("/archive")
	Action("archive", func() {
		Routing(POST("/"))
		Description("Archive HTTP response")
		Payload(ArchivePayload)
		Response(OK, ArchiveMedia)
	})
	Action("read", func() {
		Routing(GET("/:id"))
		Description("Read HTTP response from archive")
		Params(func() {
			Param("id", Integer, "ID of archive", func() {
				Minimum(0)
			})
		})
		Response(OK, ArchiveMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("health", func() {
	BasePath("/health")
	Action("show", func() {
		Routing(GET("/"))
		Description("Health check endpoint")
		Response(OK, "text/plain")
	})
})

var ArchivePayload = Type("ArchivePayload", func() {
	Attribute("status", Integer, "HTTP status", func() {
		Minimum(0)
		Example(200)
	})
	Attribute("body", String, "HTTP response body content")
	Required("status", "body")
})

var ArchiveMedia = MediaType("application/vnd.goa.archive", func() {
	Description("Archive is an HTTP response archive")
	TypeName("ArchiveMedia")
	Reference(ArchivePayload)
	Attributes(func() {
		Attribute("href", String, "The archive resouce href", func() {
			Pattern("^/archive/[0-9]+$")
			Example("/archive1/")
		})
		Attribute("status")
		Attribute("body")
		Required("href", "status", "body")
	})
	View("default", func() {
		Attribute("href")
		Attribute("status")
		Attribute("body")
	})
})

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("fetcher", func() {
	Title("The goa AWS X-Ray example upstream service")
	Description("Fetcher is a service that makes GET requests to arbitrary URLs and stores the results in the downstream 'archiver' service.")
	Host("localhost:8081")
	Scheme("http")
})

var _ = Resource("fetcher", func() {
	CanonicalActionName("fetch")
	Action("fetch", func() {
		Description("Fetch makes a GET request to the given URL and stores the results in the archiver service which must be running or the request fails")
		Routing(GET("fetch/*url"))
		Params(func() {
			Param("url", String, "URL to be fetched", func() {
				Format("uri")
			})
		})
		Response(OK, FetchMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var FetchMedia = MediaType("application/vnd.goa.fetch", func() {
	Description("FetchResponse contains the HTTP status code returned by the fetched service and the href to the archived results in the archiver service")
	TypeName("FetchMedia")
	Attributes(func() {
		Attribute("status", Integer, "HTTP status code returned by fetched service", func() {
			Minimum(0)
			Example(200)
		})
		Attribute("archive_href", String, "The href to the corresponding archive in the archiver service", func() {
			Pattern("^/archive/[0-9]+$")
			Example("/archive/1")
		})
		Required("status", "archive_href")
	})
	View("default", func() {
		Attribute("status")
		Attribute("archive_href")
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

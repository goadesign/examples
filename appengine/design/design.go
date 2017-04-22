package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("appengine", func() {
	Title("The appengine example")
	Description("A simple appengine example")
	Contact(func() {
		Name("goa team")
		Email("admin@goa.design")
		URL("http://goa.design")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("goa guide")
		URL("http://goa.design/getting-started.html")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/")

	Origin("*", func() {
		Methods("GET")
		MaxAge(600)
		Credentials()
	})
})

var ExampleMediaType = MediaType("application/vnd.example+json", func() {
	Description("example MediaType")
	Attributes(func() {
		Attribute("message", String, func() {
			Example("Hello")
		})
	})
	View("default", func() {
		Attribute("message")
	})
})

var _ = Resource("hello", func() {
	BasePath("/hello")
	Action("show", func() {
		Routing(
			GET("/"),
		)
	})
	Response(OK, ExampleMediaType)
	Response(BadRequest, ErrorMedia)
})

package design

import (
	_ "github.com/goadesign/examples/design_division/design/resources"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("design_division", func() {
	Title("The design division example")
	Description("A simple design division sample")
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

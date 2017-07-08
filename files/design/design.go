package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("files", func() {
	Title("The files example")
	Description("An example to serve static assets")
})

var _ = Resource("swagger", func() {
	Files("/swagger/*filepath", "public/swagger/")
})

var _ = Resource("schema", func() {
	Files("/schema/*filepath", "public/schema/")
})

package media

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Hello = MediaType("application/vnd.hello+json", func() {
	Description("hello")
	Attributes(func() {
		Attribute("message", String, func() {
			Example("Hello")
		})
	})
	View("default", func() {
		Attribute("message")
	})
})

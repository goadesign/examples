package media

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Sample = MediaType("application/vnd.sample+json", func() {
	Description("sample")
	Attributes(func() {
		Attribute("sample", String, func() {
			Example("sample")
		})
	})
	View("default", func() {
		Attribute("sample")
	})
})

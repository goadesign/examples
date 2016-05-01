package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("types", func() {
	Description("Fake API used to showcase the types DSL")
	Host("localhost")
})

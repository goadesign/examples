package resources

import (
	"github.com/goadesign/examples/design_division/design/media"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("hello", func() {
	BasePath("/hello")
	Action("show", func() {
		Routing(GET(""))
	})
	Response(OK, media.Hello)
	Response(BadRequest, ErrorMedia)
})

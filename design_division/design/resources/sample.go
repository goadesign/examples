package resources

import (
	"github.com/goadesign/examples/design_division/design/media"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("sample", func() {
	BasePath("/sample")
	Action("show", func() {
		Routing(GET(""))
	})
	Response(OK, media.Sample)
	Response(BadRequest, ErrorMedia)
})

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BasicAuth defines a security scheme using basic authentication.
var BasicAuth = BasicAuthSecurity("basic_auth")

var _ = Resource("BasicAuthEndpoints", func() {
	Description("This resource uses basic auth to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(BasicAuth)

	Action("secured", func() {
		Description("This action is secured with the basic_auth scheme")
		Routing(GET("/basic"))
		Response(OK)
		Response(Unauthorized)
	})

	Action("unsecured", func() {
		Description("This action does not require auth")
		Routing(GET("/basic/unsecured"))
		NoSecurity()
		Response(OK)
	})
})

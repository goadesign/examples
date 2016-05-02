package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines a required scope "api".
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api", "API access")
})

var _ = Resource("JWTEndpoints", func() {
	Description("This resource uses JWT to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(JWT)

	Action("secured", func() {
		Description("This action is secured with the jwt scheme")
		Routing(GET("/jwt"))
		Response(OK)
		Response(Unauthorized)
	})

	Action("unsecured", func() {
		Description("This action does not require auth")
		Routing(GET("/jwt/unsecured"))
		NoSecurity()
		Response(OK)
	})
})

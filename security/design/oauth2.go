package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// OAuth defines a security scheme using OAuth2. The scheme uses the standard "access code" flow
// where the application must first retrieve an authorization code to request for the refresn and
// access tokens.
//
var OAuth2 = OAuth2Security("oauth2", func() {
	// /oauth2/code is the path to the action that produces access codes
	// /oauth2/token is the path to the action that produces refresh and access tokens from
	// access codes and existing refresh tokens.
	AccessCodeFlow("/oauth2/code", "/oauth2/token")
	// All secured API requests must carry the "api" scope.
	Scope("api")
})

var _ = Resource("OAuth2Endpoints", func() {
	Description("This resource uses OAuth2 to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(OAuth2)

	Action("secured", func() {
		Description("This action is secured with the oauth2 scheme")
		Routing(GET("/oauth2"))
		Response(OK)
		Response(Unauthorized)
	})

	Action("extra_scope", func() {
		Description(`This action requires an additional scope on top of "api"`)
		Routing(GET("/oauth2/extra_scope"))
		Security(OAuth2, func() {
			Scope("extra")
		})
		Response(OK)
		Response(Unauthorized)
	})

	Action("unsecured", func() {
		Description("This action does not require auth")
		Routing(GET("/oauth2/unsecured"))
		NoSecurity()
		Response(OK)
	})
})

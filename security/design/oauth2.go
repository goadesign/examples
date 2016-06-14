package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/goadesign/oauth2/design" // import oauth2 provider design
)

// Define OAuth2 provider endpoints and OAuth2 scopes.
// This defines the OAuth2Sec security scheme used by the API endpoints below.
// See the github.com/goadesign/oauth2 package for details.
var OAuth2Sec = OAuth2("/oauth2/authorize", "/oauth2/token", func() {
	// List all available scopes
	Scope("api:read")
	Scope("api:write")
})

// oauth2 defined the API endpoints secured via OAuth2.
var _ = Resource("oauth2", func() {
	Description("This resource uses OAuth2 to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(OAuth2Sec, func() {
		Scope("api:read") // Require "api:read" scope in all actions by default
	})

	Action("secured", func() {
		Description("This action is secured with the oauth2 scheme")
		Routing(GET("/oauth2/read"))
		Response(OK)
		Response(Unauthorized)
	})

	Action("write", func() {
		Description(`This action requires an additional scope on top of "api:read"`)
		Routing(POST("/oauth2/write"))
		Security(OAuth2Sec, func() {
			Scope("api:write") // Require 'api:write' scope on top of 'api:read'
		})
		Response(OK)
		Response(Unauthorized)
	})

	Action("unsecured", func() {
		Description("This action does not require auth")
		Routing(GET("/oauth2/unsecured"))
		NoSecurity() // Override the need to auth
		Response(OK)
	})
})

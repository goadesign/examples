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

	Action("secure", func() {
		Description("This action is secure with the oauth2 scheme")
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

	Action("unsecure", func() {
		Description("This action does not require auth")
		Routing(GET("/oauth2/unsecure"))
		NoSecurity() // Override the need to auth
		Response(OK)
	})

	Action("handle_redirect", func() {
		Description("Handles the OAuth2 authorization code redirect by logging the retrieved code. This action would be implemented by a 3rd party service in a real application")
		Routing(GET("/oauth2/handle_redirect"))
		NoSecurity() // Override the need to auth
		Params(func() {
			Param("code", String, "The OAuth2 authorization code returned by the OAuth2 provider")
		})
		Response(NoContent)
	})
})

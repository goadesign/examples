package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const (
	// AuthCodePath is the request path to the OAuth2 authorize code endpoint.
	// This endpoint should prompt the resource owner for authorization and upon approval
	// redirect to a given URL with the authorization code.
	AuthCodePath = "/oauth2/auth"

	// AccessTokenPath is the request path to the OAuth2 access token refresh endpoint.
	// This endpoint should validate the given authorization code or refresh token and provide a
	// new access token.
	AccessTokenPath = "/oauth2/token"
)

// OAuth2 defines a security scheme using OAuth2. The scheme uses the standard "access code" flow
// where the client application must first retrieve an authorization code by requesting access to
// the user then use the code to request for the access token.
//
var OAuth2 = OAuth2Security("oauth2", func() {
	// AccessCodeFlow defines a "authorization code" OAuth2 flow
	// see https://tools.ietf.org/html/rfc6749#section-1.3.1
	AccessCodeFlow(AuthCodePath, AccessTokenPath)
	// All secured API requests must carry the "api:read" scope.
	// This is an example of scope usage, the OAuth2 scope construct does not define the
	// semantic attached to the scope values.
	Scope("api:read", "API access")
	// Some actions may also require the "api:write" scope.
	Scope("api:write", "API write access")
})

// AccessTokenBasicAuth defines the basic auth used to make requests to retrieve the initial
// authorization code and the refresh token as well as to refresh access tokens.
// The username and password must correspond to the client id and secret and be encoded using form
// encoding as described in https://tools.ietf.org/html/rfc6749#section-2.3.1
var AccessTokenBasicAuth = BasicAuthSecurity("oauth2_client_basic_auth", func() {
	Description("Basic auth used by client to make the requests needed to retrieve and refresh access tokens")
})

// TokenMedia describes the response sent in case of successful access token request.
// See https://tools.ietf.org/html/rfc6749#section-5.1
var TokenMedia = MediaType("application/vnd.goa.example.oauth2.token+json", func() {
	Description("OAuth2 access token request successful response, see https://tools.ietf.org/html/rfc6749#section-5.1")
	Attributes(func() {
		Attribute("access_token", String, "The access token issued by the authorization server")
		Attribute("token_type", String, `The type of the token issued, e.g. "bearer" or "mac"`)
		Attribute("expires_in", Integer, "The lifetime in seconds of the access token")
		Attribute("refresh_token", String, "The refresh token")
		Attribute("scope", String, "The scope of the access token")
		Required("access_token", "token_type")
	})
	View("default", func() {
		Attribute("access_token")
		Attribute("token_type")
		Attribute("expires_in")
		Attribute("refresh_token")
		Attribute("scope")
	})
})

// OAuth2ErrorMedia describes responses sent in case of invalid request to the provider endpoints.
// See https://tools.ietf.org/html/rfc6749#section-4.1.2.1
var OAuth2ErrorMedia = MediaType("application/vnd.goa.example.oauth2.error+json", func() {
	Description("OAuth2 error response, see https://tools.ietf.org/html/rfc6749#section-5.2")
	Attributes(func() {
		Attribute("error", String, "Error returned by authorization server", func() {
			Enum("invalid_request", "invalid_client", "invalid_grant", "unauthorized_client", "unsupported_grant_type")
		})
		Attribute("error_description", String, "Human readable ASCII text providing additional information")
		Attribute("error_uri", String, "A URI identifying a human-readable web page with information about the error")
		Required("error")
	})
	View("default", func() {
		Attribute("error")
		Attribute("error_description")
		Attribute("error_uri")
	})
})

// RefreshTokenPayload describes the body sent by the client to refresh a token.
// See https://tools.ietf.org/html/rfc6749#section-6
var RefreshTokenPayload = Type("RefreshTokenPayload", func() {
	Description("Payload sent by client to refresh a token, see https://tools.ietf.org/html/rfc6749#section-6")
	Attribute("grant_type", String, `Value MUST be set to "refresh_token".`, func() {
		Enum("refresh_token")
	})
	Attribute("refresh_token", String, "The refresh token issued to the client")
	Attribute("scope", String, "The scope of the access request")
	Required("grant_type", "refresh_token")
})

// The resource that implements the OAuth2 standard defined by RFC6749.
// See https://tools.ietf.org/html/rfc6749
var _ = Resource("OAuth2Provider", func() {
	Description("This resource implements the OAuth2 authorization code flow")

	Action("request_auth", func() {
		Description("Request coming from client requesting authorization code")
		Routing(GET(AuthCodePath))
		Params(func() {
			Param("response_type", String, `Value MUST be set to "code"`, func() {
				Enum("code")
			})
			Param("client_id", String, "The client identifier")
			Param("redirect_uri", String, "Redirection endpoint")
			Param("scope", String, "The scope of the access request")
			Param("state", String, "An opaque value used by the client to maintain state between the request and callback")
			Required("response_type", "client_id")
		})
		Response(Found, func() {
			Headers(func() {
				Header("Location", String, "Redirect URL containing the authorization code and state param")
			})
		})
		Response(BadRequest, OAuth2ErrorMedia)
	})

	Action("exchange_token", func() {
		Description("Request coming from client to retrieve access token from authorization code")
		Routing(POST(AccessTokenPath))
		Security(AccessTokenBasicAuth)
		Response(OK, TokenMedia)
		Response(BadRequest, OAuth2ErrorMedia)
	})

	Action("refresh_token", func() {
		Description("Request coming from client to refresh expired access token")
		Routing(POST(AccessTokenPath + "/refresh"))
		Security(AccessTokenBasicAuth)
		Payload(RefreshTokenPayload)
		Response(OK, TokenMedia)
		Response(BadRequest, OAuth2ErrorMedia)
	})
})

// OAuth2Endpoints defined the API endpoints secured via OAuth2.
var _ = Resource("OAuth2Endpoints", func() {
	Description("This resource uses OAuth2 to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(OAuth2, func() {
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
		Security(OAuth2, func() {
			Scope("api:write") // Require 'api:write' scope
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

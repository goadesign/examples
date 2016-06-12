package main

import (
	"net/url"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
)

// OAuth2ProviderController contains the endpoints required to implement OAuth2.
type OAuth2ProviderController struct {
	*goa.Controller
}

// NewOAuth2ProviderController creates a OAuth2Provider controller.
func NewOAuth2ProviderController(service *goa.Service) *OAuth2ProviderController {
	return &OAuth2ProviderController{Controller: service.NewController("OAuth2ProviderController")}
}

// RequestAuth returns an authorization code after validating the client should have access with the
// resource owner.
func (c *OAuth2ProviderController) RequestAuth(ctx *RequestAuthOAuth2ProviderContext) error {
	ctx.ParseForm()

	// Extract authorization code request elements from incoming request URL.
	var (
		clientID     string // ID of client requesting authorization code
		redirectURI  string // Redirect URI used to build URL that receives authorization code
		state        string // Data that should be passed through back to the requesting app
		scope        string // Requested oauth2 scope
		responseType string // Must be "code" for the authorization code flow
	)
	{
		uri, err := url.QueryUnescape(ctx.Form.Get("redirect_uri"))
		if err != nil {
			return ctx.BadRequest("invalid redirect URI")
		}
		redirectURI, err = url.Parse(uri)
		if err != nil {
			// Note: passing the redirect URI in the authorize request is not required
			// by the standard.
			return ctx.BadRequest("invalid redirect URI")
		}
		state = ctx.Form.Get("state")
		scope = ctx.Form.Get("scope")
		clientID = ctx.Form.Get("client_id")
		responseType = ctx.Form.Get("response_type")
	}

	// Now validate authorization request

	// Check responseType, must be "code" for this oauth2 flow.
	if responseType != "code" {
		return ctx.Unauthorized("bad response type")
	}

	// Check client ID, this is the client ID of the service requesting the authorization code.
	// For example if we implemented Facebook this would be the ID of the application requesting
	// the authorization code.
	if clientID != "42" {
		return ctx.Unauthorized("invalid client")
	}

	// Validate redirect URI.
	// This URI is used to form the redirect URL receiving the generated authorization code.
	// In general it should be associated with the client. To keep with the example - if we
	// implemented Facebook this would be the redirect URL configured for the application.
	if redirect_uri != "https://iredirect.baz" {
		return ctx.Unauthorized("invalid redirect")
	}

	// This is where the user (resource owner) should be prompted to authorize the client with
	// the requested scope. If we implemented Facebook we'd display the "Request For Permission"
	// dialog.

	// The code below would thus run in a different action handler, the one that receives the
	// authorization grant from the resource owner (i.e. the post back from the form rendered
	// above if we implemented Facebook). We keep things simple here and redirect right away.

	// Compute code
	code := "43"

	// And finally redirect
	redirectURI.Query.Set("state", state)
	redirectURI.Query.Set("code", code)
	redirectURI.RawQuery = redirectURI.Query.Encode()
	ctx.Request.Header.Set("Location", redirectURI)

	return ctx.Found()
}

// AccessToken creates an access token after validating the authorization code or refresh token.
func (c *OAuth2ProviderController) AccessToken(ctx *app.AccessTokenOAuth2ProviderContext) error {
	ctx.ParseForm()

	// Extract access token request elements from incoming request URL.
	var (
		grantType   string // Must be "authorization_code", a real service could also support "refresh_token"
		redirectURI string // Redirect URI used to build URL that receives access token
		code        string // The authorization code previously obtained via AuthorizationCide.
	)
	{
		uri, err := url.QueryUnescape(ctx.Form.Get("redirect_uri"))
		if err != nil {
			return ctx.BadRequest("invalid redirect URI")
		}
		redirectURI, err = url.Parse(uri)
		if err != nil {
			// Note: passing the redirect URI in the authorize request is not required
			// by the standard.
			return ctx.BadRequest("invalid redirect URI")
		}
		state = ctx.Form.Get("state")
		scope = ctx.Form.Get("scope")
		clientID = ctx.Form.Get("client_id")
		responseType = ctx.Form.Get("response_type")
	}

}

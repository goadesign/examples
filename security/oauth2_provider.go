package main

import (
	"fmt"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/oauth2"
)

// Hard-coded credentials to keep example simple.
const (
	// TheAuthorizationCode is a hard coded authorization code returned by the dummy provider.
	TheAuthorizationCode = "authcode"
	// TheRefreshToken is a hard coded refresh token returned by the dummy provider.
	TheRefreshToken = "refreshtoken"
	// TheAccessToken is a hard coded access token returned by the dummy provider.
	TheAccessToken = "accesstoken"
	// OAuth2ClientID is the only authorized client ID
	OAuth2ClientID = "client"
	// OAuth2ClientSecret is the only authorized client secret
	OAuth2ClientSecret = "secret"
)

// OAuth2ProviderController implements the OAuth2Provider resource.
type OAuth2ProviderController struct {
	*oauth2.ProviderController
}

// NewOAuth2ProviderController creates a OAuth2Provider controller.
func NewOAuth2ProviderController(service *goa.Service, provider oauth2.Provider) *OAuth2ProviderController {
	return &OAuth2ProviderController{
		ProviderController: oauth2.NewProviderController(service, provider),
	}
}

// Authorize runs the authorize action.
func (c *OAuth2ProviderController) Authorize(ctx *app.AuthorizeOauth2ProviderContext) error {
	return c.ProviderController.Authorize(ctx.Context, ctx.ResponseWriter, ctx.Request)
}

// GetToken runs the get_token action.
func (c *OAuth2ProviderController) GetToken(ctx *app.GetTokenOauth2ProviderContext) error {
	p := ctx.Payload
	return c.ProviderController.GetToken(ctx.Context, ctx.ResponseWriter, p.GrantType,
		p.Code, p.RedirectURI, p.RefreshToken, p.Scope)
}

// NewOAuth2Provider returns the object implementing the OAuth2 provider.
// Typically the implementation would require access to a persistent store to keep track of
// clients, authorization codes and tokens.
func NewOAuth2Provider() oauth2.Provider {
	return &DummyProvider{
		authorizationCode: TheAuthorizationCode,
		refreshToken:      TheRefreshToken,
		accessToken:       TheAccessToken,
		clientID:          OAuth2ClientID,
		clientSecret:      OAuth2ClientSecret,
	}
}

// DummyProvider implements a OAuth2 provider that always returns the same authorization code and
// tokens.
type DummyProvider struct {
	authorizationCode string
	refreshToken      string
	accessToken       string
	clientID          string
	clientSecret      string
}

// Authorize is a dummy implementation of oauth2.Provider.Authorize that simply returns the hard
// coded authorization code.
func (p *DummyProvider) Authorize(clientID, scope, redirectURI string) (code string, err error) {
	return p.authorizationCode, nil
}

// Exchange is a dummy implementation of oauth2.Provider.Exchange that simply returns the hard coded
// refresh and access tokens.
func (p *DummyProvider) Exchange(clientID, code, redirectURI string) (refreshToken, accessToken string, expiresIn int, err error) {
	return p.refreshToken, p.accessToken, 3600, nil
}

// Refresh is a dummy implementation of oauth2.Provider.Refresh that simply returns the hard coded
// refresh and access tokens.
func (p *DummyProvider) Refresh(refreshToken, scope string) (newRefreshToken, accessToken string, expiresIn int, err error) {
	return p.refreshToken, p.accessToken, 3600, nil
}

// Authenticate is a dummy implementation of oauth2.Provider.Authenticate that simply validates the
// client credentials against hard-coded values.
func (p *DummyProvider) Authenticate(clientID, clientSecret string) error {
	if clientID != p.clientID || clientSecret != p.clientSecret {
		return fmt.Errorf("invalid client credentials")
	}
	return nil
}

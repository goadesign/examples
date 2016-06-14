package main

import (
	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
)

// OAuth2ProviderController implements the OAuth2Provider resource.
type OAuth2ProviderController struct {
	*goa.Controller
}

// NewOAuth2ProviderController creates a OAuth2Provider controller.
func NewOAuth2ProviderController(service *goa.Service) *OAuth2ProviderController {
	return &OAuth2ProviderController{Controller: service.NewController("OAuth2ProviderController")}
}

// Authorize runs the authorize action.
func (c *OAuth2ProviderController) Authorize(ctx *app.AuthorizeOAuth2ProviderContext) error {
	// TBD: implement
	return nil
}

// ExchangeToken runs the exchange_token action.
func (c *OAuth2ProviderController) GetToken(ctx *app.GetTokenOAuth2ProviderContext) error {
	// TBD: implement
	res := &app.TokenMedia{}
	return ctx.OK(res)
}

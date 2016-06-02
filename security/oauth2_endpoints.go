package main

import (
	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
)

// OAuth2EndpointsController implements the OAuth2Endpoints resource.
type OAuth2EndpointsController struct {
	*goa.Controller
}

// NewOAuth2EndpointsController creates a OAuth2Endpoints controller.
func NewOAuth2EndpointsController(service *goa.Service) *OAuth2EndpointsController {
	return &OAuth2EndpointsController{Controller: service.NewController("OAuth2EndpointsController")}
}

// ExtraScope runs the extra_scope action.
func (c *OAuth2EndpointsController) ExtraScope(ctx *app.ExtraScopeOAuth2EndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Secured runs the secured action.
func (c *OAuth2EndpointsController) Secured(ctx *app.SecuredOAuth2EndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *OAuth2EndpointsController) Unsecured(ctx *app.UnsecuredOAuth2EndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

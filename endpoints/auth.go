package main

import (
	"github.com/goadesign/examples/endpoints/app"
	"github.com/goadesign/examples/endpoints/middleware"
	"github.com/goadesign/goa"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Basic runs the basic action.
func (c *AuthController) Basic(ctx *app.BasicAuthContext) error {
	res := app.Auth(*middleware.UserInfo(ctx.Context))
	return ctx.OK(&res)
}

// JWT runs the jwt action.
func (c *AuthController) JWT(ctx *app.JWTAuthContext) error {
	res := app.Auth(*middleware.UserInfo(ctx.Context))
	return ctx.OK(&res)
}

package main

import (
	"net/http"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			// A real app would do something more interesting here
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}

			// Proceed
			goa.LogInfo(ctx, "auth", "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}

// BasicAuthEndpointsController implements the BasicAuthEndpoints resource.
type BasicAuthEndpointsController struct {
	*goa.Controller
}

// NewBasicAuthEndpointsController creates a BasicAuthEndpoints controller.
func NewBasicAuthEndpointsController(service *goa.Service) *BasicAuthEndpointsController {
	return &BasicAuthEndpointsController{Controller: service.NewController("BasicAuthEndpointsController")}
}

// Secured runs the secured action.
func (c *BasicAuthEndpointsController) Secured(ctx *app.SecuredBasicAuthEndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *BasicAuthEndpointsController) Unsecured(ctx *app.UnsecuredBasicAuthEndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

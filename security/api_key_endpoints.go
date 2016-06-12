package main

import (
	"net/http"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// NewAPIKeyMiddleware creates a middleware that checks for the presence of an authorization header
// and validates its content.
func NewAPIKeyMiddleware() goa.Middleware {
	// Instantiate API Key security scheme details generated from design
	scheme := app.NewAPIKeySecurity()

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			key := req.Header.Get(scheme.Name)
			// A real app would do something more interesting here
			if key == "" {
				goa.LogInfo(ctx, "failed api key auth")
				return ErrUnauthorized("missing auth")
			}
			// Proceed.
			goa.LogInfo(ctx, "auth", "apikey", "key", key)
			return h(ctx, rw, req)
		}
	}
}

// APIKeyEndpointsController implements the APIKeyEndpoints resource.
type APIKeyEndpointsController struct {
	*goa.Controller
}

// NewAPIKeyEndpointsController creates a APIKeyEndpoints controller.
func NewAPIKeyEndpointsController(service *goa.Service) *APIKeyEndpointsController {
	return &APIKeyEndpointsController{Controller: service.NewController("APIKeyEndpointsController")}
}

// Secured runs the secured action.
func (c *APIKeyEndpointsController) Secured(ctx *app.SecuredAPIKeyEndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *APIKeyEndpointsController) Unsecured(ctx *app.UnsecuredAPIKeyEndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

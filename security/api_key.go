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

// APIKeyController implements the APIKey resource.
type APIKeyController struct {
	*goa.Controller
}

// NewAPIKeyController creates a APIKey controller.
func NewAPIKeyController(service *goa.Service) *APIKeyController {
	return &APIKeyController{Controller: service.NewController("APIKeyController")}
}

// Secured runs the secured action.
func (c *APIKeyController) Secured(ctx *app.SecuredAPIKeyContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *APIKeyController) Unsecured(ctx *app.UnsecuredAPIKeyContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

package main

import (
	"context"
	"net/http"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
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
			if len(key) == 0 || key == "Bearer" {
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

// Secure runs the secure action.
func (c *APIKeyController) Secure(ctx *app.SecureAPIKeyContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *APIKeyController) Unsecure(ctx *app.UnsecureAPIKeyContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

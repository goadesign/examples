package main

import (
	"net/http"
	"strings"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

const (
	// TheAccessToken is a hard coded value for the only valid access token accepted by this
	// controller. This is done this way to keep the example simple. In a real service tokens
	// would be created dynamically and kept in a persistent store.
	TheAccessToken = "authorizeme"
)

// OAuth2EndpointsController implements the OAuth2Endpoints resource.
type OAuth2EndpointsController struct {
	*goa.Controller
}

// NewOAuth2Middleware creates a middleware that checks for the presence of an authorization header
// and validates its content.
func NewOAuth2Middleware() goa.Middleware {
	// Instantiate scheme described in design to retrieve
	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			token := req.Header["Authorization"]
			if token == nil {
				return ErrUnauthorized("missing auth header")
			}
			if len(token) < 10 || !strings.HasPrefix(token, "Bearer: ") {
				return ErrUnauthorized("invalid auth header")
			}
			token := token[8:]

			// Validate token here against value stored in DB for example
			if token != TheAccessToken {
				return ErrUnauthorized("invalid token")
			}
			return h(ctx, rw, req)
		}
	}
}

// NewOAuth2EndpointsController creates a OAuth2Endpoints controller.
func NewOAuth2EndpointsController(service *goa.Service) *OAuth2EndpointsController {
	// Mount authorization middleware
	app.UseOAuth2Middleware(NewOAuth2Middleware())

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

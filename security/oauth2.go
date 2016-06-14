package main

import (
	"net/http"
	"strings"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// OAuth2Controller implements the OAuth2 resource.
type OAuth2Controller struct {
	*goa.Controller
}

const (
	// TheAccessToken is the only authorized bearer value for incoming request authorization
	// headers. A real service would generate random tokens and persist them in a backend.
	TheAccessToken = "acessme"
)

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
			tok := token[0]
			if len(tok) < 10 || !strings.HasPrefix(tok, "Bearer: ") {
				return ErrUnauthorized("invalid auth header")
			}
			tok = tok[8:]

			// Validate token here against value stored in DB for example
			if tok != TheAccessToken {
				return ErrUnauthorized("invalid token")
			}
			return h(ctx, rw, req)
		}
	}
}

// NewOAuth2Controller creates a OAuth2 controller.
func NewOAuth2Controller(service *goa.Service) *OAuth2Controller {
	return &OAuth2Controller{Controller: service.NewController("OAuth2Controller")}
}

// Secure runs the secure action.
func (c *OAuth2Controller) Secure(ctx *app.SecureOauth2Context) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *OAuth2Controller) Unsecure(ctx *app.UnsecureOauth2Context) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Write runs the write action.
func (c *OAuth2Controller) Write(ctx *app.WriteOauth2Context) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

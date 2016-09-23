//************************************************************************//
// API "adder": Application Security
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/examples/endpoints/design
// --out=$(GOPATH)/src/github.com/goadesign/examples/endpoints
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

type (
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseAPIKeyMiddleware mounts the api_key auth middleware onto the service.
func UseAPIKeyMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("api_key"), middleware)
}

// NewAPIKeySecurity creates a api_key security definition.
func NewAPIKeySecurity() *goa.APIKeySecurity {
	def := goa.APIKeySecurity{
		In:   goa.LocQuery,
		Name: "key",
	}
	return &def
}

// UseGoogleJWTMiddleware mounts the google_jwt auth middleware onto the service.
func UseGoogleJWTMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("google_jwt"), middleware)
}

// NewGoogleJWTSecurity creates a google_jwt security definition.
func NewGoogleJWTSecurity() *goa.OAuth2Security {
	def := goa.OAuth2Security{
		Flow:             "implicit",
		TokenURL:         "",
		AuthorizationURL: "https://goa-endpoints.appspot.com/auth"}
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}

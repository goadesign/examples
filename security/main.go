package main

import (
	"net/http"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/net/context"
)

var (
	// Unauthorized is the error returned for unauthorized requests.
	Unauthorized = goa.NewErrorClass("unauthorized", 401)
)

func main() {
	// Create service
	service := goa.New("Secured API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "APIKeyEndpoints" controller
	c := NewAPIKeyEndpointsController(service)
	app.MountAPIKeyEndpointsController(service, c)
	// Mount "BasicAuthEndpoints" controller
	c2 := NewBasicAuthEndpointsController(service)
	app.MountBasicAuthEndpointsController(service, c2)
	// Mount "JWTEndpoints" controller
	c3 := NewJWTEndpointsController(service)
	app.MountJWTEndpointsController(service, c3)
	// Mount "OAuth2Endpoints" controller
	c4 := NewOAuth2EndpointsController(service)
	app.MountOAuth2EndpointsController(service, c4)

	// Mount security middlewares
	app.UseBasicAuthMiddleware(service, NewBasicAuthMiddleware())
	app.UseApiKeyMiddleware(service, NewApiKeyMiddleware())
	app.UseJWTMiddleware(service, NewJWTMiddleware())
	app.UseOauth2Middleware(service, NewOAuth2Middleware())

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

// NewOAuth2Middleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
func NewJWTMiddleware() goa.Middleware {
	// Instantiate JWT security scheme details generated from design
	scheme := app.NewJWTSecurity()

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			key := req.Header.Get(scheme.Name)
			authorized(ctx, "auth", "JWT", "key", key) {
			// Proceed.
			return h(ctx, rw, req)
		}
	}
}

// Custom authorization logic, log and always return true for the purpose of the example.
func authorized(ctx context.Context, keyvals ...interface{}) bool {
	goa.LogInfo(ctx, "Authorizing", keyvals...)
	return true
}

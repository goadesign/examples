package main

import (
	"fmt"
	"os"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
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
	app.UseAPIKeyMiddleware(service, NewAPIKeyMiddleware())
	jwtMiddleware, err := NewJWTMiddleware()
	exitOnFailure(err)
	app.UseJWTMiddleware(service, jwtMiddleware)
	app.UseOauth2Middleware(service, NewOAuth2Middleware())

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

// exitOnFailure prints a fatal error message and exits the process with status 1.
func exitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
	os.Exit(1)
}

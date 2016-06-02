package main

import (
	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
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

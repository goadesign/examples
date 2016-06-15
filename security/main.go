package main

import (
	"fmt"
	"os"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/oauth2"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
)

func main() {
	// Create service
	service := goa.New("Secure API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middlewares
	jwtMiddleware, err := NewJWTMiddleware()
	exitOnFailure(err)
	app.UseBasicAuthMiddleware(service, NewBasicAuthMiddleware())
	app.UseAPIKeyMiddleware(service, NewAPIKeyMiddleware())
	app.UseJWTMiddleware(service, jwtMiddleware)
	app.UseOAuth2Middleware(service, NewOAuth2Middleware())

	// Security middleware used to secure the creation of JWT tokens.
	app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	// Security middleware used by the OAuth2 provider to authorize the client to obtain access
	// tokens.
	provider := NewOAuth2Provider()
	oauth2ClientAuth := oauth2.NewOAuth2ClientBasicAuthMiddleware(provider)
	app.UseOauth2ClientBasicAuthMiddleware(service, oauth2ClientAuth)

	// Mount "APIKey" controller
	c := NewAPIKeyController(service)
	app.MountAPIKeyController(service, c)
	// Mount "BasicAuth" controller
	c2 := NewBasicController(service)
	app.MountBasicController(service, c2)
	// Mount "JWT" controller
	c3, err := NewJWTController(service)
	exitOnFailure(err)
	app.MountJWTController(service, c3)
	// Mount "OAuth2" controller
	c4 := NewOAuth2Controller(service)
	app.MountOauth2Controller(service, c4)
	// Mount "OAuth2Provider" controller
	c5 := NewOAuth2ProviderController(service, provider)
	app.MountOauth2ProviderController(service, c5)

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

//go:generate goagen bootstrap -d github.com/goadesign/examples/endpoints/design
package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/goadesign/examples/endpoints/app"
	"github.com/goadesign/examples/endpoints/middleware"
	"github.com/goadesign/goa"
	goamiddleware "github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("Adder")

	// Setup middleware
	service.Use(goamiddleware.RequestID())
	service.Use(goamiddleware.LogRequest(true))
	service.Use(goamiddleware.ErrorHandler(service, true))
	service.Use(goamiddleware.Recover())

	// Setup security middleware which loads user info
	app.UseAPIKeyMiddleware(service, middleware.Endpoints())
	app.UseGoogleJWTMiddleware(service, middleware.Endpoints())

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)

	// Mount "auth" controller
	a := NewAuthController(service)
	app.MountAuthController(service, a)

	// Mount heathcheck endpoint.
	service.Mux.Handle("GET", "/_ah/health", func(rw http.ResponseWriter, _ *http.Request, _ url.Values) {
		rw.Write([]byte("ok"))
	})

	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}
	if err := service.ListenAndServe(":" + port); err != nil {
		service.LogError(err.Error())
	}
}

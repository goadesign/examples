//go:generate goagen bootstrap -d github.com/goadesign/examples/graceful/design
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/goadesign/examples/graceful/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tylerb/graceful"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)

	// Setup graceful server
	server := &graceful.Server{
		Timeout: time.Duration(15) * time.Second,
		Server:  &http.Server{Addr: ":8080", Handler: service.Mux},
	}

	// And run it
	log.Fatal(server.ListenAndServe())
}

//go:generate goagen bootstrap -d github.com/goadesign/examples/websocket/design
package main

import (
	"log"

	"github.com/goadesign/examples/websocket/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())

	// Mount "echo" controller
	c := NewEchoController(service)
	app.MountEchoController(service, c)

	log.Fatal(service.ListenAndServe(":8080"))
}

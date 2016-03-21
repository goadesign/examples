package main

import (
	"github.com/goadesign/examples/adder/app"
	"github.com/goadesign/examples/adder/swagger"
	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
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
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	service.ListenAndServe(":8080")
}

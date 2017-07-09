//go:generate goagen bootstrap -d github.com/goadesign/examples/design_division/design

package main

import (
	"github.com/goadesign/examples/design_division/app"
	"github.com/goadesign/examples/design_division/controllers"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("design_division")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "hello" controller
	c := controllers.NewHelloController(service)
	app.MountHelloController(service, c)
	// Mount "sample" controller
	c2 := controllers.NewSampleController(service)
	app.MountSampleController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

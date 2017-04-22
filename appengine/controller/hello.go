package controller

import (
	"github.com/goadesign/examples/appengine/app"
	"github.com/goadesign/goa"
)

// HelloController implements the hello resource.
type HelloController struct {
	*goa.Controller
}

// NewHelloController creates a hello controller.
func NewHelloController(service *goa.Service) *HelloController {
	return &HelloController{Controller: service.NewController("HelloController")}
}

// Show runs the show action.
func (c *HelloController) Show(ctx *app.ShowHelloContext) error {
	// HelloController_Show: start_implement

	// Put your logic here

	// HelloController_Show: end_implement
	msg := "Hello World"
	res := &app.Example{&msg}
	return ctx.OK(res)
}

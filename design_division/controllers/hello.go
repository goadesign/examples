package controllers

import (
	"github.com/goadesign/examples/design_division/app"
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
	h := "hello"
	res := &app.Hello{}
	res.Message = &h
	return ctx.OK(res)
}

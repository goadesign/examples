package controller

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

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
	appCtx := appengine.NewContext(ctx.Request)
	log.Debugf(appCtx, "Debugging\n")

	// HelloController_Show: end_implement
	msg := "Hello World"
	res := &app.Example{&msg}
	return ctx.OK(res)
}

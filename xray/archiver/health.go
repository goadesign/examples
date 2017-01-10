package main

import (
	"github.com/goadesign/examples/xray/archiver/app"
	"github.com/goadesign/goa"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Show runs the show action.
func (c *HealthController) Show(ctx *app.ShowHealthContext) error {
	return ctx.OK([]byte("OK"))
}

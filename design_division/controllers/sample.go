package controllers

import (
	"github.com/goadesign/examples/design_division/app"
	"github.com/goadesign/goa"
)

// SampleController implements the sample resource.
type SampleController struct {
	*goa.Controller
}

// NewSampleController creates a sample controller.
func NewSampleController(service *goa.Service) *SampleController {
	return &SampleController{Controller: service.NewController("SampleController")}
}

// Show runs the show action.
func (c *SampleController) Show(ctx *app.ShowSampleContext) error {
	// SampleController_Show: start_implement

	// Put your logic here

	// SampleController_Show: end_implement
	s := "Sample"
	res := &app.Sample{}
	res.Sample = &s
	return ctx.OK(res)
}

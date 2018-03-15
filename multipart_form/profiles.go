//go:generate goagen bootstrap -d github.com/goadesign/examples/multipart_form/design
package main

import (
	"github.com/goadesign/examples/multipart_form/app"
	"github.com/goadesign/goa"
)

// ProfilesController implements the profiles resource.
type ProfilesController struct {
	*goa.Controller
}

// NewProfilesController creates a profiles controller.
func NewProfilesController(service *goa.Service) *ProfilesController {
	return &ProfilesController{Controller: service.NewController("ProfilesController")}
}

// Submit runs the submit action.
func (c *ProfilesController) Submit(ctx *app.SubmitProfilesContext) error {
	res := &app.ResultMedia{
		Name:     ctx.Payload.Name,
		Birthday: ctx.Payload.Birthday,
	}
	return ctx.OK(res)
}

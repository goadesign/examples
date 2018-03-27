//go:generate goagen bootstrap -d github.com/goadesign/examples/multipart_form/design
package main

import (
	"fmt"
	"io"
	"os"

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
	file, err := ctx.Payload.Icon.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	f, err := os.OpenFile("./icons/"+ctx.Payload.Icon.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to save file: %s", err) // causes a 500 response
	}
	defer f.Close()
	io.Copy(f, file)
	return ctx.OK(res)
}

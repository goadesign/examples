//go:generate goagen bootstrap -d github.com/goadesign/examples/form_encoding/design
package main

import (
	"github.com/goadesign/examples/form_encoding/app"
	"github.com/goadesign/goa"
)

// SurveyFormController implements the survey_form resource.
type SurveyFormController struct {
	*goa.Controller
}

// NewSurveyFormController creates a survey_form controller.
func NewSurveyFormController(service *goa.Service) *SurveyFormController {
	return &SurveyFormController{Controller: service.NewController("SurveyFormController")}
}

// Submit runs the submit action.
func (c *SurveyFormController) Submit(ctx *app.SubmitSurveyFormContext) error {
	vote := ctx.Payload.Vote
	var message string
	if vote == "yes" {
		message = "yay :)"
	} else {
		message = "nay :("
	}
	res := &app.ResultMedia{Message: message, Name: ctx.Payload.Name}
	return ctx.OK(res)
}

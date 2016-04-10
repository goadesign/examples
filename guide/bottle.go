//go:generate goagen bootstrap -d github.com/goadesign/examples/guide/design
package main

import (
	"fmt"

	"github.com/goadesign/examples/guide/app"
	"github.com/goadesign/goa"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	if ctx.BottleID == 0 {
		// Emulate a missing record with ID 0
		return ctx.NotFound()
	}
	// Build the resource using the generated data structure
	bottle := app.GoaExampleBottle{
		ID:   ctx.BottleID,
		Name: fmt.Sprintf("Bottle #%d", ctx.BottleID),
		Href: app.BottleHref(ctx.BottleID),
	}

	// Let the generated code produce the HTTP response using the
	// media type described in the design (BottleMedia).
	return ctx.OK(&bottle)
}

package main

import (
	"strconv"
	"time"

	"github.com/goadesign/examples/graceful/app"
	"github.com/goadesign/goa"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("operands")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	time.Sleep(time.Duration(10) * time.Second)
	return ctx.OK([]byte(strconv.Itoa(ctx.Left + ctx.Right)))
}

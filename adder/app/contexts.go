//************************************************************************//
// API "adder": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/examples/adder
// --design=github.com/goadesign/examples/adder/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(ctx context.Context) (*AddOperandsContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := AddOperandsContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawLeft := req.Params.Get("left")
	if rawLeft != "" {
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	rawRight := req.Params.Get("right")
	if rawRight != "" {
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.StackErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

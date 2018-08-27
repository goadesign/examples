// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/goadesign/examples/appengine/design
// --out=$(GOPATH)/src/github.com/goadesign/examples/appengine
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// ShowHelloContext provides the hello show action context.
type ShowHelloContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowHelloContext parses the incoming request URL and body, performs validations and creates the
// context used by the hello controller show action.
func NewShowHelloContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowHelloContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowHelloContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowHelloContext) OK(r *Example) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.example+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowHelloContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

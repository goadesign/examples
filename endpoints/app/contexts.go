//************************************************************************//
// API "adder": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/examples/endpoints/design
// --out=$(GOPATH)/src/github.com/goadesign/examples/endpoints
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// BasicAuthContext provides the auth basic action context.
type BasicAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewBasicAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller basic action.
func NewBasicAuthContext(ctx context.Context, service *goa.Service) (*BasicAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := BasicAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *BasicAuthContext) OK(r *Auth) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa-cellar.auth+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// JWTAuthContext provides the auth jwt action context.
type JWTAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewJWTAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller jwt action.
func NewJWTAuthContext(ctx context.Context, service *goa.Service) (*JWTAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := JWTAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *JWTAuthContext) OK(r *Auth) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa-cellar.auth+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

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
func NewAddOperandsContext(ctx context.Context, service *goa.Service) (*AddOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := AddOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

//************************************************************************//
// API "echo": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/examples/websocket
// --design=github.com/goadesign/examples/websocket/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// ConnectEchoContext provides the echo connect action context.
type ConnectEchoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Initial *string
}

// NewConnectEchoContext parses the incoming request URL and body, performs validations and creates the
// context used by the echo controller connect action.
func NewConnectEchoContext(ctx context.Context) (*ConnectEchoContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ConnectEchoContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawInitial := req.Params.Get("initial")
	if rawInitial != "" {
		rctx.Initial = &rawInitial
	}
	return &rctx, err
}

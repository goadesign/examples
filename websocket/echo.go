package main

import (
	"io"

	"github.com/goadesign/examples/websocket/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
)

// EchoController implements the echo resource.
type EchoController struct {
	*goa.Controller
}

// NewEchoController creates a echo controller.
func NewEchoController(service *goa.Service) *EchoController {
	return &EchoController{Controller: service.NewController("echo")}
}

// Connect runs the connect action.
func (c *EchoController) Connect(ctx *app.ConnectEchoContext) error {
	c.ConnectWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// ConnectWSHandler establishes a websocket connection to run the connect action.
func (c *EchoController) ConnectWSHandler(ctx *app.ConnectEchoContext) websocket.Handler {
	var initial string
	if ctx.Initial != nil {
		initial = *ctx.Initial
	}
	return func(ws *websocket.Conn) {
		ws.Write([]byte(initial))
		// echo websocket server
		io.Copy(ws, ws)
	}
}

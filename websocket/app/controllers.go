//************************************************************************//
// API "echo": Application Controllers
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
	"net/http"
)

// inited is true if initService has been called
var inited = false

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	if inited {
		return
	}
	inited = true

	// Setup encoders and decoders
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Encoder(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder(goa.NewXMLEncoder, "application/xml")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// EchoController is the controller interface for the Echo actions.
type EchoController interface {
	goa.Muxer
	Connect(*ConnectEchoContext) error
}

// MountEchoController "mounts" a Echo resource controller on the given service.
func MountEchoController(service *goa.Service, ctrl EchoController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewConnectEchoContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Connect(rctx)
	}
	service.Mux.Handle("GET", "/echo", ctrl.MuxHandler("Connect", h, nil))
	service.Info("mount", "ctrl", "Echo", "action", "Connect", "route", "GET /echo")
}

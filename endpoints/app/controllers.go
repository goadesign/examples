//************************************************************************//
// API "adder": Application Controllers
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
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	Basic(*BasicAuthContext) error
	JWT(*JWTAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/auth/basic", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/auth/jwt", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewBasicAuthContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Basic(rctx)
	}
	h = handleAuthOrigin(h)
	h = handleSecurity("api_key", h)
	service.Mux.Handle("GET", "/auth/basic", ctrl.MuxHandler("Basic", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Basic", "route", "GET /auth/basic", "security", "api_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewJWTAuthContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.JWT(rctx)
	}
	h = handleAuthOrigin(h)
	h = handleSecurity("google_jwt", h)
	service.Mux.Handle("GET", "/auth/jwt", ctrl.MuxHandler("JWT", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "JWT", "route", "GET /auth/jwt", "security", "google_jwt")
}

// handleAuthOrigin applies the CORS response headers corresponding to the origin.
func handleAuthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://swagger.goa.design") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Muxer
	Add(*AddOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service *goa.Service, ctrl OperandsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/add/:left/:right", ctrl.MuxHandler("preflight", handleOperandsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddOperandsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	h = handleOperandsOrigin(h)
	h = handleSecurity("api_key", h)
	service.Mux.Handle("GET", "/add/:left/:right", ctrl.MuxHandler("Add", h, nil))
	service.LogInfo("mount", "ctrl", "Operands", "action", "Add", "route", "GET /add/:left/:right", "security", "api_key")
}

// handleOperandsOrigin applies the CORS response headers corresponding to the origin.
func handleOperandsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://swagger.goa.design") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

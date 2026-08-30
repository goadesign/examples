// This file configures the calc HTTP server, including the response body used
// when a request is missing a required field.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	gencalc "goa.design/examples/error/gen/calc"
	gencalcsvr "goa.design/examples/error/gen/http/calc/server"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	goa "goa.design/goa/v3/pkg"
)

type (
	// missingFieldError is written as a JSON string when a request omits a
	// required field.
	missingFieldError string

	// divByZeroError preserves the response declared for division by zero when
	// the custom formatter handles all HTTP errors.
	divByZeroError struct {
		Message string `form:"message" json:"message" xml:"message"`
	}
)

// StatusCode reports that a missing required field makes the request invalid.
func (missingFieldError) StatusCode() int {
	return http.StatusBadRequest
}

// StatusCode reports that division by zero makes the request invalid.
func (*divByZeroError) StatusCode() int {
	return http.StatusBadRequest
}

// handleHTTPServer starts an HTTP server on the given URL. It shuts down the
// server when the error channel receives a value.
func handleHTTPServer(ctx context.Context, u *url.URL, calcEndpoints *gencalc.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		calcServer *gencalcsvr.Server
	)
	{
		eh := errorHandler(logger)
		calcServer = gencalcsvr.New(calcEndpoints, mux, dec, enc, eh, formatError)
		if debug {
			servers := goahttp.Servers{
				calcServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	gencalcsvr.Mount(mux, calcServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler, ReadHeaderTimeout: time.Second * 60}
	for _, m := range calcServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			logger.Printf("failed to shutdown: %v", err)
		}
	}()
}

// formatError writes missing-field validation errors as JSON strings and
// preserves the response declared for division by zero. Every other error uses
// Goa's standard HTTP error response.
func formatError(ctx context.Context, err error) goahttp.Statuser {
	var serviceError *goa.ServiceError
	if errors.As(err, &serviceError) && serviceError.Name == "missing_field" {
		return missingFieldError(serviceError.Message)
	}

	var divideError *gencalc.DivByZero
	if errors.As(err, &divideError) {
		return &divByZeroError{Message: divideError.Message}
	}

	return goahttp.NewErrorResponse(ctx, err)
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

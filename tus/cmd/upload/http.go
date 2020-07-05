package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	tussvr "goa.design/examples/tus/gen/http/tus/server"
	tus "goa.design/examples/tus/gen/tus"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, tusEndpoints *tus.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {
	mux := goahttp.NewMuxer()
	tusServer := tussvr.New(tusEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, errorHandler(logger), nil)
	if debug {
		tusServer.Use(httpmdlwr.Debug(mux, os.Stdout))
	}
	tussvr.Mount(mux, tusServer)
	handler := httpmdlwr.Log(middleware.NewLogger(logger))(mux)
	handler = httpmdlwr.RequestID()(handler)
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range tusServer.Mounts {
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

		// Shutdown gracefully with a 300s timeout to allow for ongoing downloads to finish.
		ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

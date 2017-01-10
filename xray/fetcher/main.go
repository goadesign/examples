//go:generate goagen bootstrap -d github.com/goadesign/examples/xray/fetcher/design

package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/goadesign/examples/xray/fetcher/app"
	"github.com/goadesign/examples/xray/fetcher/services"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/xray"
)

func main() {
	var (
		listen   = flag.String("l", "127.0.0.1:8081", "The service listen `address`")
		daemon   = flag.String("d", "127.0.0.1:2000", "The daemon `UDP address`")
		archiver = flag.String("a", "127.0.0.1:8080", "The archiver service `host:port`")
	)
	flag.Parse()

	// Create service
	service := goa.New("fetcher")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Setup AWS X-Ray middleware
	m, err := xray.New("fetcher", *daemon)
	if err != nil {
		service.LogError("xray", "err", err)
		os.Exit(1)
	}
	service.Use(middleware.Tracer(100, xray.NewID, xray.NewTraceID))
	service.Use(m)

	// Setup downstream service clients
	a := services.NewArchiver(*archiver, http.DefaultClient)
	ok, err := a.WaitUntilHealthy(service.Context, time.Minute)
	if err != nil {
		service.LogError("failed to connect to archiver service", "err", err)
		os.Exit(1)
	}
	if !ok {
		service.LogError("downstream archiver service unhealthy after a minute, giving up.")
		os.Exit(1)
	}

	// Mount "fetcher" controller
	c := NewFetcherController(service, a)
	app.MountFetcherController(service, c)

	// Mount "health" controller
	c2 := NewHealthController(service)
	app.MountHealthController(service, c2)

	// Start service
	if err := service.ListenAndServe(*listen); err != nil {
		service.LogError("startup", "err", err)
	}
}

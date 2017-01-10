//go:generate goagen bootstrap -d github.com/goadesign/examples/xray/archiver/design

package main

import (
	"flag"
	"os"

	"github.com/goadesign/examples/xray/archiver/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/xray"
)

func main() {
	var (
		listen = flag.String("l", "127.0.0.1:8080", "The service listen `address`")
		daemon = flag.String("d", "127.0.0.1:2000", "The daemon `UDP address`")
	)
	flag.Parse()

	// Create service
	service := goa.New("archiver")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(false))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Setup AWS X-Ray middleware
	m, err := xray.New("archiver", *daemon)
	if err != nil {
		service.LogError("xray", "err", err)
		os.Exit(1)
	}
	service.Use(middleware.Tracer(100, xray.NewID, xray.NewTraceID))
	service.Use(m)

	// Mount "archiver" controller
	c := NewArchiverController(service)
	app.MountArchiverController(service, c)

	// Mount "health" controller
	c2 := NewHealthController(service)
	app.MountHealthController(service, c2)

	// Start service
	if err := service.ListenAndServe(*listen); err != nil {
		service.LogError("startup", "err", err)
	}
}

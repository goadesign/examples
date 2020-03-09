package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	uploaddownload "goa.design/examples/upload_download"
	updown "goa.design/examples/upload_download/gen/updown"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		dirF      = flag.String("dir", "public", "Relative path to directory containing files to download and storing uploaded files.")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[upload-download] ", log.Ltime)
	}

	// Initialize the services.
	updownSvc, err := uploaddownload.NewUpdown(*dirF, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		updownEndpoints *updown.Endpoints
	)
	{
		updownEndpoints = updown.NewEndpoints(updownSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	addr := "http://localhost:8080"
	u, _ := url.Parse(addr)
	if *secureF {
		u.Scheme = "https"
	}
	if *domainF != "" {
		u.Host = *domainF
	}
	if *httpPortF != "" {
		h := strings.Split(u.Host, ":")[0]
		u.Host = h + ":" + *httpPortF
	}
	handleHTTPServer(ctx, u, updownEndpoints, &wg, errc, logger, *dbgF)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

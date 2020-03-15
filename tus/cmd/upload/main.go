package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"time"

	tussvc "goa.design/examples/tus"
	"goa.design/examples/tus/gen/tus"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		uploadDir = flag.String("dir", ".", "Upload target directory")
		maxSize   = flag.Int64("max-size", 1024*1024*20, "Maximum allowed upload size in bytes")
		timeout   = flag.Int64("timeout", 0, "Maximum number of seconds before an upload is aborted (default is 0 which never aborts)")
		retention = flag.Int64("retention", 0, "Maximum number of seconds after upload completes that results should be kept in memory")
		hostF     = flag.String("host", "development", "Server host (valid values: development)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	logger := log.New(os.Stderr, "[tus] ", log.Ltime)
	svc := tussvc.New(
		newFileWriterFunc(*uploadDir),
		*maxSize,
		(time.Duration(*timeout))*time.Second,
		(time.Duration(*retention))*time.Second,
		logger,
	)
	endpoints := tus.NewEndpoints(svc)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "development":
		{
			addr := "http://localhost:8000/upload"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, endpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: development)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

// newFileWriterFunc returns a function that creates a io.Writer each time it is
// invoked that writes to the given file path relative to uploadDir.
func newFileWriterFunc(uploadDir string) func(string, *int64) (io.WriteCloser, error) {
	return func(fname string, _ *int64) (io.WriteCloser, error) {
		return os.Open(filepath.Join(uploadDir, fname))
	}
}

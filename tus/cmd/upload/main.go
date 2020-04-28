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
	"sync"
	"time"

	tussvc "goa.design/examples/tus"
	"goa.design/examples/tus/gen/tus"
	"goa.design/examples/tus/persist"
)

func main() {
	var (
		uploadDir = flag.String("dir", ".", "Upload target directory")
		maxSize   = flag.Int64("max-size", 1024*1024*20, "Maximum allowed upload size in bytes")
		timeout   = flag.Int64("timeout", 0, "Maximum number of seconds before an upload is aborted (default is 0 which never aborts)")
		httpPortF = flag.String("port", "8080", "HTTP listen port")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	logger := log.New(os.Stderr, "[tus] ", log.Ltime)
	svc := tussvc.New(
		persist.NewInMemory(),
		newFileWriterFunc(*uploadDir),
		*maxSize,
		(time.Duration(*timeout))*time.Second,
		logger,
	)
	endpoints := tus.NewEndpoints(svc)

	// SIGINT and SIGTERM signals cause the services to stop gracefully.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	u, err := url.Parse(fmt.Sprintf("http://localhost:%s/upload", *httpPortF))
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid port %#v: %s\n", *httpPortF, err)
		os.Exit(1)
	}
	handleHTTPServer(ctx, u, endpoints, &wg, errc, logger, *dbgF)

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
		return os.Create(filepath.Join(uploadDir, fname))
	}
}

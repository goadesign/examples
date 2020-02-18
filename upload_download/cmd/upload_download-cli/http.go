package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	cli "goa.design/examples/upload_download/gen/http/cli/upload_download"
	"goa.design/examples/upload_download/gen/updown"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

func doHTTP(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	var (
		doer goahttp.Doer
	)
	{
		doer = &http.Client{Timeout: time.Duration(timeout) * time.Second}
		if debug {
			doer = goahttp.NewDebugDoer(doer)
		}
	}

	ep, p, err := cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
	)

	// Wrap the file reader created by the flag parser into a reader that
	// produces a multipart request body.
	if pup, ok := p.(*updown.UploadRequestData); ok {
		pup.Body = newMultipartFileReader(pup.Body.(*os.File))
	}

	return ep, p, err
}

func httpUsageCommands() string {
	return cli.UsageCommands()
}

func httpUsageExamples() string {
	return cli.UsageExamples()
}

// multipartFileReader wraps a file io.Reader into a reader that returns a
// multipart form-data enveloppe around the content of the file.
type multipartFileReader struct {
	io.Reader
	f *os.File
}

// newMultipartFileReader creates a reader that streams a multipart request
// header, the content of the given file and the multipart request trailer.
func newMultipartFileReader(f *os.File) io.ReadCloser {
	fname := "dummy"
	if fi, err := f.Stat(); err == nil {
		fname = fi.Name()
	}
	header := bytes.NewBufferString(fmt.Sprintf("--goa\r\nContent-Disposition: form-data; name=%q; filename=%q\r\nContent-Type:application/octet-stream\r\n\r\n", fname, fname))
	trailer := bytes.NewBufferString("\r\n--goa--\r\n")

	return &multipartFileReader{Reader: io.MultiReader(header, f, trailer), f: f}
}

// Close the file upon closing the reader.
func (mr *multipartFileReader) Close() error {
	return mr.f.Close()
}

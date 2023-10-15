package main

import (
	"net/http"
	"time"

	cli "goa.design/examples/basic/gen/http/cli/calc"
	goahttp "goa.design/goa/v3/http"
	"goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/http/middleware/xray"
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
		// Wrap doer with X-Ray and trace client middleware. Order is very important.
		doer = xray.WrapDoer(doer)
		doer = middleware.WrapDoer(doer)
	}

	return cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
	)
}
func httpUsageCommands() string {
	return cli.UsageCommands()
}

func httpUsageExamples() string {
	return cli.UsageExamples()
}

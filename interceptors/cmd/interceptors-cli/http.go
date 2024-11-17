package main

import (
	"net/http"
	"time"

	cli "goa.design/examples/interceptors/gen/http/cli/interceptors"
	"goa.design/examples/interceptors/gen/interceptors"
	interceptorsex "goa.design/examples/interceptors/interceptors"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

func doHTTP(scheme, host string, timeout int, debug bool) (goa.Endpoint, any, error) {
	var (
		doer                     goahttp.Doer
		interceptorsInterceptors interceptors.ClientInterceptors
	)
	{
		doer = &http.Client{Timeout: time.Duration(timeout) * time.Second}
		if debug {
			doer = goahttp.NewDebugDoer(doer)
		}
		interceptorsInterceptors = interceptorsex.NewInterceptorsClientInterceptors()
	}

	return cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
		interceptorsInterceptors,
	)
}

func httpUsageCommands() string {
	return cli.UsageCommands()
}

func httpUsageExamples() string {
	return cli.UsageExamples()
}

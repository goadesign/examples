package main

import (
	"net/http"
	"time"

	cellar "goa.design/examples/cellar"
	cli "goa.design/examples/cellar/gen/http/cli/cellar"
	"goa.design/goa"
	goahttp "goa.design/goa/http"
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

	return cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
		cellar.StorageMultiAddEncoderFunc,
		cellar.StorageMultiUpdateEncoderFunc,
	)
}

func httpUsageCommands() string {
	return cli.UsageCommands()
}

func httpUsageExamples() string {
	return cli.UsageExamples()
}

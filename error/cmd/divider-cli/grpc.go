package main

import (
	"fmt"
	"os"

	cli "goa.design/examples/error/gen/grpc/cli/divider"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

func doGRPC(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to gRPC server at %s: %v\n", host, err)
	}
	return cli.ParseEndpoint(conn)
}

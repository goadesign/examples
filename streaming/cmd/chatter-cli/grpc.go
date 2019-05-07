package main

import (
	"fmt"
	"os"

	cli "goa.design/examples/streaming/gen/grpc/cli/chatter"
	"goa.design/goa/v3"
	"google.golang.org/grpc"
)

func doGRPC(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("could not connect to gRPC server at %s: %v", host, err))
	}
	return cli.ParseEndpoint(conn)
}

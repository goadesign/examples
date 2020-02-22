package main

import (
	"fmt"
	"os"

	cli "goa.design/examples/streaming/gen/grpc/cli/chatter"
	"goa.design/goa"
	"google.golang.org/grpc"
)

func doGRPC(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to gRPC server at %s: %v\n", host, err)
	}
	return cli.ParseEndpoint(conn)
}

package main

import (
	"fmt"
	"os"

	cli "goa.design/examples/error/gen/grpc/cli/calc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// See: https://github.com/googleapis/google-api-go-client/issues/2559.
	_ "google.golang.org/genproto/googleapis/type/datetime"
)

func doGRPC(_, host string, _ int, _ bool) (goa.Endpoint, any, error) {
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to gRPC server at %s: %v\n", host, err)
	}
	return cli.ParseEndpoint(conn)
}

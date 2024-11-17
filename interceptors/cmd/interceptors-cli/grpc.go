package main

import (
	"fmt"
	"os"

	cli "goa.design/examples/interceptors/gen/grpc/cli/interceptors"
	"goa.design/examples/interceptors/interceptors"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doGRPC(_, host string, _ int, _ bool) (goa.Endpoint, any, error) {
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to gRPC server at %s: %v\n", host, err)
	}
	interceptorsInterceptors := interceptors.NewInterceptorsClientInterceptors()
	return cli.ParseEndpoint(
		conn,
		interceptorsInterceptors,
	)
}

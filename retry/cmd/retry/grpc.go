package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"sync"

	"goa.design/clue/debug"
	"goa.design/clue/log"
	retrypb "goa.design/examples/retry/gen/grpc/retry/pb"
	retrysvr "goa.design/examples/retry/gen/grpc/retry/server"
	retry "goa.design/examples/retry/gen/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, retryEndpoints *retry.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		retryServer *retrysvr.Server
	)
	{
		retryServer = retrysvr.New(retryEndpoints, nil)
	}

	// Create interceptor which sets up the logger in each request context.
	chain := grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx))
	if dbg {
		// Log request and response content if debug logs are enabled.
		chain = grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx), debug.UnaryServerInterceptor())
	}

	// Initialize gRPC server
	srv := grpc.NewServer(chain)

	// Register the servers.
	retrypb.RegisterRetryServer(srv, retryServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			log.Printf(ctx, "serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			if lis == nil {
				errc <- fmt.Errorf("failed to listen on %q", u.Host)
			}
			log.Printf(ctx, "gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		log.Printf(ctx, "shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}

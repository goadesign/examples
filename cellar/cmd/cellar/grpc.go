package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"sync"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommeliersvr "goa.design/examples/cellar/gen/grpc/sommelier/server"
	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storagesvr "goa.design/examples/cellar/gen/grpc/storage/server"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	storage "goa.design/examples/cellar/gen/storage"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, sommelierEndpoints *sommelier.Endpoints, storageEndpoints *storage.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		sommelierServer *sommeliersvr.Server
		storageServer   *storagesvr.Server
	)
	{
		sommelierServer = sommeliersvr.New(sommelierEndpoints, nil)
		storageServer = storagesvr.New(storageEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	sommelierpb.RegisterSommelierServer(srv, sommelierServer)
	storagepb.RegisterStorageServer(srv, storageServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Printf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Printf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Printf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}

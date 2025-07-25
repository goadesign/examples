// Code generated by goa v3.21.5, DO NOT EDIT.
//
// interceptors gRPC server
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package server

import (
	"context"
	"errors"

	interceptorspb "goa.design/examples/interceptors/gen/grpc/interceptors/pb"
	interceptors "goa.design/examples/interceptors/gen/interceptors"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the interceptorspb.InterceptorsServer interface.
type Server struct {
	GetH goagrpc.UnaryHandler
	interceptorspb.UnimplementedInterceptorsServer
}

// New instantiates the server struct with the interceptors service endpoints.
func New(e *interceptors.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GetH: NewGetHandler(e.Get, uh),
	}
}

// NewGetHandler creates a gRPC handler which serves the "interceptors" service
// "get" endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in interceptorspb.InterceptorsServer
// interface.
func (s *Server) Get(ctx context.Context, message *interceptorspb.GetRequest) (*interceptorspb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "interceptors")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "NotFound":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			case "Unavailable":
				return nil, goagrpc.NewStatusError(codes.Unavailable, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*interceptorspb.GetResponse), nil
}

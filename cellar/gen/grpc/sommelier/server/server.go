// Code generated by goa v3.19.1, DO NOT EDIT.
//
// sommelier gRPC server
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package server

import (
	"context"
	"errors"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the sommelierpb.SommelierServer interface.
type Server struct {
	PickH goagrpc.UnaryHandler
	sommelierpb.UnimplementedSommelierServer
}

// New instantiates the server struct with the sommelier service endpoints.
func New(e *sommelier.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		PickH: NewPickHandler(e.Pick, uh),
	}
}

// NewPickHandler creates a gRPC handler which serves the "sommelier" service
// "pick" endpoint.
func NewPickHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodePickRequest, EncodePickResponse)
	}
	return h
}

// Pick implements the "Pick" method in sommelierpb.SommelierServer interface.
func (s *Server) Pick(ctx context.Context, message *sommelierpb.PickRequest) (*sommelierpb.StoredBottleCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "pick")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sommelier")
	resp, err := s.PickH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "no_criteria":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "no_match":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sommelierpb.StoredBottleCollection), nil
}

// Code generated by goa v3.21.5, DO NOT EDIT.
//
// storage gRPC server
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package server

import (
	"context"
	"errors"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storage "goa.design/examples/cellar/gen/storage"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the storagepb.StorageServer interface.
type Server struct {
	ListH        goagrpc.UnaryHandler
	ShowH        goagrpc.UnaryHandler
	AddH         goagrpc.UnaryHandler
	RemoveH      goagrpc.UnaryHandler
	RateH        goagrpc.UnaryHandler
	MultiAddH    goagrpc.UnaryHandler
	MultiUpdateH goagrpc.UnaryHandler
	storagepb.UnimplementedStorageServer
}

// New instantiates the server struct with the storage service endpoints.
func New(e *storage.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ListH:        NewListHandler(e.List, uh),
		ShowH:        NewShowHandler(e.Show, uh),
		AddH:         NewAddHandler(e.Add, uh),
		RemoveH:      NewRemoveHandler(e.Remove, uh),
		RateH:        NewRateHandler(e.Rate, uh),
		MultiAddH:    NewMultiAddHandler(e.MultiAdd, uh),
		MultiUpdateH: NewMultiUpdateHandler(e.MultiUpdate, uh),
	}
}

// NewListHandler creates a gRPC handler which serves the "storage" service
// "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in storagepb.StorageServer interface.
func (s *Server) List(ctx context.Context, message *storagepb.ListRequest) (*storagepb.StoredBottleCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*storagepb.StoredBottleCollection), nil
}

// NewShowHandler creates a gRPC handler which serves the "storage" service
// "show" endpoint.
func NewShowHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeShowRequest, EncodeShowResponse)
	}
	return h
}

// Show implements the "Show" method in storagepb.StorageServer interface.
func (s *Server) Show(ctx context.Context, message *storagepb.ShowRequest) (*storagepb.ShowResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "show")
	ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
	resp, err := s.ShowH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "not_found":
				var er *storage.NotFound
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewShowNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*storagepb.ShowResponse), nil
}

// NewAddHandler creates a gRPC handler which serves the "storage" service
// "add" endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in storagepb.StorageServer interface.
func (s *Server) Add(ctx context.Context, message *storagepb.AddRequest) (*storagepb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*storagepb.AddResponse), nil
}

// NewRemoveHandler creates a gRPC handler which serves the "storage" service
// "remove" endpoint.
func NewRemoveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveRequest, EncodeRemoveResponse)
	}
	return h
}

// Remove implements the "Remove" method in storagepb.StorageServer interface.
func (s *Server) Remove(ctx context.Context, message *storagepb.RemoveRequest) (*storagepb.RemoveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "remove")
	ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
	resp, err := s.RemoveH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*storagepb.RemoveResponse), nil
}

// NewRateHandler creates a gRPC handler which serves the "storage" service
// "rate" endpoint.
func NewRateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRateRequest, EncodeRateResponse)
	}
	return h
}

// Rate implements the "Rate" method in storagepb.StorageServer interface.
func (s *Server) Rate(ctx context.Context, message *storagepb.RateRequest) (*storagepb.RateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "rate")
	ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
	resp, err := s.RateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*storagepb.RateResponse), nil
}

// NewMultiAddHandler creates a gRPC handler which serves the "storage" service
// "multi_add" endpoint.
func NewMultiAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeMultiAddRequest, EncodeMultiAddResponse)
	}
	return h
}

// MultiAdd implements the "MultiAdd" method in storagepb.StorageServer
// interface.
func (s *Server) MultiAdd(ctx context.Context, message *storagepb.MultiAddRequest) (*storagepb.MultiAddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "multi_add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
	resp, err := s.MultiAddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*storagepb.MultiAddResponse), nil
}

// NewMultiUpdateHandler creates a gRPC handler which serves the "storage"
// service "multi_update" endpoint.
func NewMultiUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeMultiUpdateRequest, EncodeMultiUpdateResponse)
	}
	return h
}

// MultiUpdate implements the "MultiUpdate" method in storagepb.StorageServer
// interface.
func (s *Server) MultiUpdate(ctx context.Context, message *storagepb.MultiUpdateRequest) (*storagepb.MultiUpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "multi_update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
	resp, err := s.MultiUpdateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*storagepb.MultiUpdateResponse), nil
}

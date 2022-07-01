// Code generated by goa v3.7.8, DO NOT EDIT.
//
// storage gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"context"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildListFunc builds the remote method to invoke for "storage" service
// "list" endpoint.
func BuildListFunc(grpccli storagepb.StorageClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.List(ctx, reqpb.(*storagepb.ListRequest), opts...)
		}
		return grpccli.List(ctx, &storagepb.ListRequest{}, opts...)
	}
}

// DecodeListResponse decodes responses from the storage list endpoint.
func DecodeListResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*storagepb.StoredBottleCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "list", "*storagepb.StoredBottleCollection", v)
	}
	res := NewListResult(message)
	vres := storageviews.StoredBottleCollection{Projected: res, View: view}
	if err := storageviews.ValidateStoredBottleCollection(vres); err != nil {
		return nil, err
	}
	return storage.NewStoredBottleCollection(vres), nil
}

// BuildShowFunc builds the remote method to invoke for "storage" service
// "show" endpoint.
func BuildShowFunc(grpccli storagepb.StorageClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Show(ctx, reqpb.(*storagepb.ShowRequest), opts...)
		}
		return grpccli.Show(ctx, &storagepb.ShowRequest{}, opts...)
	}
}

// EncodeShowRequest encodes requests sent to storage show endpoint.
func EncodeShowRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*storage.ShowPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "show", "*storage.ShowPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewProtoShowRequest(payload), nil
}

// DecodeShowResponse decodes responses from the storage show endpoint.
func DecodeShowResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*storagepb.ShowResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "show", "*storagepb.ShowResponse", v)
	}
	res := NewShowResult(message)
	vres := &storageviews.StoredBottle{Projected: res, View: view}
	if err := storageviews.ValidateStoredBottle(vres); err != nil {
		return nil, err
	}
	return storage.NewStoredBottle(vres), nil
}

// BuildAddFunc builds the remote method to invoke for "storage" service "add"
// endpoint.
func BuildAddFunc(grpccli storagepb.StorageClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*storagepb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &storagepb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to storage add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*storage.Bottle)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "add", "*storage.Bottle", v)
	}
	return NewProtoAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the storage add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*storagepb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "add", "*storagepb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}

// BuildRemoveFunc builds the remote method to invoke for "storage" service
// "remove" endpoint.
func BuildRemoveFunc(grpccli storagepb.StorageClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Remove(ctx, reqpb.(*storagepb.RemoveRequest), opts...)
		}
		return grpccli.Remove(ctx, &storagepb.RemoveRequest{}, opts...)
	}
}

// EncodeRemoveRequest encodes requests sent to storage remove endpoint.
func EncodeRemoveRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*storage.RemovePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "remove", "*storage.RemovePayload", v)
	}
	return NewProtoRemoveRequest(payload), nil
}

// BuildRateFunc builds the remote method to invoke for "storage" service
// "rate" endpoint.
func BuildRateFunc(grpccli storagepb.StorageClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Rate(ctx, reqpb.(*storagepb.RateRequest), opts...)
		}
		return grpccli.Rate(ctx, &storagepb.RateRequest{}, opts...)
	}
}

// EncodeRateRequest encodes requests sent to storage rate endpoint.
func EncodeRateRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(map[uint32][]string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "rate", "map[uint32][]string", v)
	}
	return NewProtoRateRequest(payload), nil
}

// BuildMultiAddFunc builds the remote method to invoke for "storage" service
// "multi_add" endpoint.
func BuildMultiAddFunc(grpccli storagepb.StorageClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.MultiAdd(ctx, reqpb.(*storagepb.MultiAddRequest), opts...)
		}
		return grpccli.MultiAdd(ctx, &storagepb.MultiAddRequest{}, opts...)
	}
}

// EncodeMultiAddRequest encodes requests sent to storage multi_add endpoint.
func EncodeMultiAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.([]*storage.Bottle)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "multi_add", "[]*storage.Bottle", v)
	}
	return NewProtoMultiAddRequest(payload), nil
}

// DecodeMultiAddResponse decodes responses from the storage multi_add endpoint.
func DecodeMultiAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*storagepb.MultiAddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "multi_add", "*storagepb.MultiAddResponse", v)
	}
	res := NewMultiAddResult(message)
	return res, nil
}

// BuildMultiUpdateFunc builds the remote method to invoke for "storage"
// service "multi_update" endpoint.
func BuildMultiUpdateFunc(grpccli storagepb.StorageClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.MultiUpdate(ctx, reqpb.(*storagepb.MultiUpdateRequest), opts...)
		}
		return grpccli.MultiUpdate(ctx, &storagepb.MultiUpdateRequest{}, opts...)
	}
}

// EncodeMultiUpdateRequest encodes requests sent to storage multi_update
// endpoint.
func EncodeMultiUpdateRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*storage.MultiUpdatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "multi_update", "*storage.MultiUpdatePayload", v)
	}
	return NewProtoMultiUpdateRequest(payload), nil
}

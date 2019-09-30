// Code generated by goa v2.0.5, DO NOT EDIT.
//
// divider gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package server

import (
	"context"

	divider "goa.design/examples/error/gen/divider"
	dividerpb "goa.design/examples/error/gen/grpc/divider/pb"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeIntegerDivideResponse encodes responses from the "divider" service
// "integer_divide" endpoint.
func EncodeIntegerDivideResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "integer_divide", "int", v)
	}
	resp := NewIntegerDivideResponse(result)
	return resp, nil
}

// DecodeIntegerDivideRequest decodes requests sent to "divider" service
// "integer_divide" endpoint.
func DecodeIntegerDivideRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *dividerpb.IntegerDivideRequest
		ok      bool
	)
	{
		if message, ok = v.(*dividerpb.IntegerDivideRequest); !ok {
			return nil, goagrpc.ErrInvalidType("divider", "integer_divide", "*dividerpb.IntegerDivideRequest", v)
		}
	}
	var payload *divider.IntOperands
	{
		payload = NewIntegerDividePayload(message)
	}
	return payload, nil
}

// EncodeDivideResponse encodes responses from the "divider" service "divide"
// endpoint.
func EncodeDivideResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(float64)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "divide", "float64", v)
	}
	resp := NewDivideResponse(result)
	return resp, nil
}

// DecodeDivideRequest decodes requests sent to "divider" service "divide"
// endpoint.
func DecodeDivideRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *dividerpb.DivideRequest
		ok      bool
	)
	{
		if message, ok = v.(*dividerpb.DivideRequest); !ok {
			return nil, goagrpc.ErrInvalidType("divider", "divide", "*dividerpb.DivideRequest", v)
		}
	}
	var payload *divider.FloatOperands
	{
		payload = NewDividePayload(message)
	}
	return payload, nil
}

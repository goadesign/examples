// Code generated by goa v3.7.5, DO NOT EDIT.
//
// calc gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package client

import (
	"context"

	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildMultiplyFunc builds the remote method to invoke for "calc" service
// "multiply" endpoint.
func BuildMultiplyFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Multiply(ctx, reqpb.(*calcpb.MultiplyRequest), opts...)
		}
		return grpccli.Multiply(ctx, &calcpb.MultiplyRequest{}, opts...)
	}
}

// EncodeMultiplyRequest encodes requests sent to calc multiply endpoint.
func EncodeMultiplyRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.MultiplyPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calc.MultiplyPayload", v)
	}
	return NewProtoMultiplyRequest(payload), nil
}

// DecodeMultiplyResponse decodes responses from the calc multiply endpoint.
func DecodeMultiplyResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.MultiplyResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calcpb.MultiplyResponse", v)
	}
	res := NewMultiplyResult(message)
	return res, nil
}

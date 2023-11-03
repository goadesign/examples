// Code generated by goa v3.14.0, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen goa.design/examples/basic/design

package server

import (
	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
)

// NewMultiplyPayload builds the payload of the "multiply" endpoint of the
// "calc" service from the gRPC request type.
func NewMultiplyPayload(message *calcpb.MultiplyRequest) *calc.MultiplyPayload {
	v := &calc.MultiplyPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewProtoMultiplyResponse builds the gRPC response type from the result of
// the "multiply" endpoint of the "calc" service.
func NewProtoMultiplyResponse(result int) *calcpb.MultiplyResponse {
	message := &calcpb.MultiplyResponse{}
	message.Field = int32(result)
	return message
}

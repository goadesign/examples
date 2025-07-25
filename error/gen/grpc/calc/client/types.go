// Code generated by goa v3.21.5, DO NOT EDIT.
//
// calc gRPC client types
//
// Command:
// $ goa gen goa.design/examples/error/design

package client

import (
	calc "goa.design/examples/error/gen/calc"
	calcpb "goa.design/examples/error/gen/grpc/calc/pb"
)

// NewProtoDivideRequest builds the gRPC request type from the payload of the
// "divide" endpoint of the "calc" service.
func NewProtoDivideRequest(payload *calc.DividePayload) *calcpb.DivideRequest {
	message := &calcpb.DivideRequest{
		Dividend: int32(payload.Dividend),
		Divisor:  int32(payload.Divisor),
	}
	return message
}

// NewDivideResult builds the result type of the "divide" endpoint of the
// "calc" service from the gRPC response type.
func NewDivideResult(message *calcpb.DivideResponse) *calc.DivideResult {
	result := &calc.DivideResult{
		Quotient: int(message.Quotient),
		Reminder: int(message.Reminder),
	}
	return result
}

// NewDivideDivByZeroError builds the error type of the "divide" endpoint of
// the "calc" service from the gRPC error response type.
func NewDivideDivByZeroError(message *calcpb.DivideDivByZeroError) *calc.DivByZero {
	er := &calc.DivByZero{
		Message: message.Message_,
	}
	return er
}

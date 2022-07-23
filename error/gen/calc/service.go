// Code generated by goa v3.7.13, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen goa.design/examples/error/design -o error

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the calc service interface.
type Service interface {
	// Divide implements divide.
	Divide(context.Context, *DividePayload) (res *DivideResult, err error)
	// Add implements add.
	Add(context.Context, *AddPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"divide", "add"}

// AddPayload is the payload type of the calc service add method.
type AddPayload struct {
	A int
	B int
}

// DivByZero is the error returned when using value 0 as divisor.
type DivByZero struct {
	// division by zero leads to infinity.
	Message string
}

// DividePayload is the payload type of the calc service divide method.
type DividePayload struct {
	Dividend int
	Divisor  int
}

// DivideResult is the result type of the calc service divide method.
type DivideResult struct {
	Quotient int
	Reminder int
}

// Error returns an error description.
func (e *DivByZero) Error() string {
	return "DivByZero is the error returned when using value 0 as divisor."
}

// ErrorName returns "DivByZero".
func (e *DivByZero) ErrorName() string {
	return "div_by_zero"
}

// MakeTimeout builds a goa.ServiceError from an error.
func MakeTimeout(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "timeout", true, false, false)
}

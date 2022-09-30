// Code generated by goa v3.8.7, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/examples/error/design -o error

package server

import (
	calc "goa.design/examples/error/gen/calc"
	goa "goa.design/goa/v3/pkg"
)

// DivideRequestBody is the type of the "calc" service "divide" endpoint HTTP
// request body.
type DivideRequestBody struct {
	Dividend *int `form:"dividend,omitempty" json:"dividend,omitempty" xml:"dividend,omitempty"`
	Divisor  *int `form:"divisor,omitempty" json:"divisor,omitempty" xml:"divisor,omitempty"`
}

// DivideResponseBody is the type of the "calc" service "divide" endpoint HTTP
// response body.
type DivideResponseBody struct {
	Quotient int `form:"quotient" json:"quotient" xml:"quotient"`
	Reminder int `form:"reminder" json:"reminder" xml:"reminder"`
}

// DivideDivByZeroResponseBody is the type of the "calc" service "divide"
// endpoint HTTP response body for the "div_by_zero" error.
type DivideDivByZeroResponseBody struct {
	// division by zero leads to infinity.
	Message string `form:"message" json:"message" xml:"message"`
}

// DivideTimeoutResponseBody is the type of the "calc" service "divide"
// endpoint HTTP response body for the "timeout" error.
type DivideTimeoutResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewDivideResponseBody builds the HTTP response body from the result of the
// "divide" endpoint of the "calc" service.
func NewDivideResponseBody(res *calc.DivideResult) *DivideResponseBody {
	body := &DivideResponseBody{
		Quotient: res.Quotient,
		Reminder: res.Reminder,
	}
	return body
}

// NewDivideDivByZeroResponseBody builds the HTTP response body from the result
// of the "divide" endpoint of the "calc" service.
func NewDivideDivByZeroResponseBody(res *calc.DivByZero) *DivideDivByZeroResponseBody {
	body := &DivideDivByZeroResponseBody{
		Message: res.Message,
	}
	return body
}

// NewDivideTimeoutResponseBody builds the HTTP response body from the result
// of the "divide" endpoint of the "calc" service.
func NewDivideTimeoutResponseBody(res *goa.ServiceError) *DivideTimeoutResponseBody {
	body := &DivideTimeoutResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDividePayload builds a calc service divide endpoint payload.
func NewDividePayload(body *DivideRequestBody) *calc.DividePayload {
	v := &calc.DividePayload{
		Dividend: *body.Dividend,
		Divisor:  *body.Divisor,
	}

	return v
}

// ValidateDivideRequestBody runs the validations defined on DivideRequestBody
func ValidateDivideRequestBody(body *DivideRequestBody) (err error) {
	if body.Dividend == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("dividend", "body"))
	}
	if body.Divisor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("divisor", "body"))
	}
	return
}

package design

import . "goa.design/goa/v3/dsl"

var _ = Service("calc", func() {

	// The timeout error can be returned by both the "divide" and "add" methods
	Error("timeout", func() { // Use default error type
		Timeout()
	})
	// Equivalent to:
	// Error("timeout", ErrorResult, func() { ...

	Method("divide", func() {
		Payload(func() {
			Field(1, "dividend", Int)
			Field(2, "divisor", Int)
			Required("dividend", "divisor")
		})
		Result(func() {
			Field(1, "quotient", Int)
			Field(2, "reminder", Int)
			Required("quotient", "reminder")
		})

		// The "div_by_zero" error is specific to the "divide" method.
		Error("div_by_zero", DivByZero, "division by 0") // Use custom error type

		HTTP(func() {
			POST("/")
			Response("div_by_zero", StatusBadRequest, func() {
				// Use HTTP status code 400 (BadRequest) to write "div_by_zero" errors
				Description("Response used for division by zero errors")
			})
			Response("timeout", StatusGatewayTimeout, func() {
				Description("Operation timed out")
			})
		})
		GRPC(func() {
			Response("div_by_zero", CodeInvalidArgument, func() {
				// Use gRPC status code 3 (InvalidArgument) to write "div_by_zero" errors
				Description("Response used for division by zero errors")
			})
			Response("timeout", CodeDeadlineExceeded, func() {
				Description("Operation timed out")
			})
		})
	})

	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int)
			Field(2, "b", Int)
			Required("a", "b")
		})
		Result(Int)
	})
})

var DivByZero = Type("DivByZero", func() {
	Description("DivByZero is the error returned when using value 0 as divisor.")
	Field(1, "message", String, "division by zero leads to infinity.")
	Required("message")
})

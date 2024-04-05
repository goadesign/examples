package helloapi

import (
	"context"
	"log"

	"goa.design/examples/httpstatus/gen/hello"
)

// hello service example implementation.
// The example methods log the requests and return zero values.
type hellosrvc struct {
	logger *log.Logger
}

// NewHello returns the hello service implementation.
func NewHello(logger *log.Logger) hello.Service {
	return &hellosrvc{logger}
}

// Hello implements hello.
func (s *hellosrvc) HelloEndpoint(ctx context.Context, p *hello.HelloPayload) (*hello.Hello, error) {
	var res hello.Hello
	res.Greeting = p.Greeting
	switch p.Greeting {
	case "hello":
		res.Outcome = "created"
	case "bye":
		res.Outcome = "accepted"
	default:
		res.Outcome = "defaultStatus"
	}
	return &res, nil
}

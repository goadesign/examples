package calc

import (
	"context"
	"log"

	calcsvc "goa.design/examples/basic/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcSvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calcsvc.Service {
	return &calcSvc{logger}
}

// Multiply implements multiply.
func (s *calcSvc) Multiply(ctx context.Context, p *calcsvc.MultiplyPayload) (int, error) {
	return p.A * p.B, nil
}

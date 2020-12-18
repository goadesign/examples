package calcapi

import (
	"context"
	"log"
	"time"

	calc "goa.design/examples/error/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Divide implements divide. It uses a convoluted implementation to illustrate returning timeout errors.
func (s *calcsrvc) Divide(ctx context.Context, p *calc.DividePayload) (res *calc.DivideResult, err error) {
	// First make sure divisor is not 0, return DivByZero error if it is.
	if p.Divisor == 0 {
		return nil, &calc.DivByZero{Message: "divide by zero"}
	}

	// Set a timeout in the context.
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	// Use a channel to wait for either timeout or completion of division
	c := make(chan *calc.DivideResult, 1)

	// Call division operation in separate Go routine
	go func() { c <- divide(p.Dividend, p.Divisor) }()

	// Wait for results or timeout
	select {
	case <-ctx.Done():
		// Timeout triggered, return timeout error
		return nil, calc.MakeTimeout(ctx.Err())
	case res := <-c:
		// Division completed successfully
		return res, nil
	}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	return p.A + p.B, nil
}

// divide implements the divide operation, it sleeps for 100 milliseconds if
// dividend is 42 to emulate timeouts.
func divide(dividend, divisor int) *calc.DivideResult {
	if dividend == 42 {
		time.Sleep(100 * time.Millisecond)
	}
	return &calc.DivideResult{
		Quotient: dividend / divisor,
		Reminder: dividend % divisor,
	}
}

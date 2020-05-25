package concatapi

import (
	"context"
	"log"

	concat "goa.design/examples/encodings/cbor/gen/concat"
)

// concat service example implementation.
// The example methods log the requests and return zero values.
type concatsrvc struct {
	logger *log.Logger
}

// NewConcat returns the concat service implementation.
func NewConcat(logger *log.Logger) concat.Service {
	return &concatsrvc{logger}
}

// Concat implements concat.
func (s *concatsrvc) Concat(ctx context.Context, p *concat.ConcatPayload) (res string, err error) {
	s.logger.Print("concat.concat")
	return p.A + p.B, nil
}

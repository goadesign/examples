package concatapi

import (
	"context"

	"goa.design/clue/log"

	concat "goa.design/examples/encodings/cbor/gen/concat"
)

// concat service example implementation.
// The example methods log the requests and return zero values.
type concatsrvc struct {
}

// NewConcat returns the concat service implementation.
func NewConcat() concat.Service {
	return &concatsrvc{}
}

// Concat implements concat.
func (s *concatsrvc) Concat(ctx context.Context, p *concat.ConcatPayload) (res string, err error) {
	log.Print(ctx, log.KV{K: "concat.concat", V: p})
	return p.A + p.B, nil
}

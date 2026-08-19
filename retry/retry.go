// Package retryapi implements the retry example service. The generated HTTP
// and gRPC clients call this service twice when its first response reports a
// temporary failure for an idempotent request.
package retryapi

import (
	"context"
	"errors"
	"sync"

	"goa.design/clue/log"
	genretry "goa.design/examples/retry/gen/retry"
)

// retrysrvc records whether each demonstration request has already received
// its intentional first-attempt failure.
type retrysrvc struct {
	mu       sync.Mutex
	attempts map[string]int
}

// NewRetry returns the retry service implementation.
func NewRetry() genretry.Service {
	return &retrysrvc{attempts: make(map[string]int)}
}

// GetMessage returns a temporary error for a request's first attempt and the
// same successful result for every later invocation with that request ID.
func (s *retrysrvc) GetMessage(ctx context.Context, p *genretry.GetMessagePayload) (*genretry.GetMessageResult, error) {
	s.mu.Lock()
	attempt := s.attempts[p.ID]
	s.attempts[p.ID] = attempt + 1
	s.mu.Unlock()

	if attempt == 0 {
		log.Printf(ctx, "retry.get_message first attempt unavailable id=%s", p.ID)
		return nil, genretry.MakeUnavailable(errors.New("intentional first-attempt failure"))
	}

	log.Printf(ctx, "retry.get_message succeeded id=%s", p.ID)
	return &genretry.GetMessageResult{
		Message: "request " + p.ID + " succeeded after an automatic retry",
	}, nil
}

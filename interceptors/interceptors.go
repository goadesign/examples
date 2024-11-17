package interceptorsapi

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"goa.design/clue/log"

	interceptors "goa.design/examples/interceptors/gen/interceptors"
)

// interceptors service example implementation.
// The example methods log the requests and return zero values.
type interceptorssrvc struct {
	records sync.Map
}

// ReturnUnavailable is a hard-coded UUID used to demonstrate the client-side Retry
// interceptor by simulating a temporary service unavailability
const ReturnUnavailable interceptors.UUID = "00000000-0000-0000-0000-000000000000"

// Hang is a hard-coded UUID used to demonstrate the Deadline interceptor
// by simulating a slow response that should trigger timeouts
const Hang interceptors.UUID = "00000000-0000-0000-0000-000000000001"

// NewInterceptors returns the interceptors service implementation.
func NewInterceptors() interceptors.Service {
	return &interceptorssrvc{}
}

// Get retrieves a record by ID demonstrating various interceptor scenarios:
// - ReturnUnavailable: Triggers the retry interceptor with a 503 response
// - Hang: Triggers the deadline interceptor by delaying the response
// - Other IDs: Normal flow demonstrating tenant and logging interceptors
func (s *interceptorssrvc) Get(ctx context.Context, p *interceptors.GetPayload) (*interceptors.GetResult, error) {
	log.Printf(ctx, "[Get] ID: %s", p.RecordID)

	switch p.RecordID {
	case ReturnUnavailable:
		return nil, interceptors.MakeUnavailable(fmt.Errorf("service temporarily unavailable - retry should handle this"))
	case Hang:
		log.Printf(ctx, "[Get] Simulating slow response for deadline demonstration")
		time.Sleep(10 * time.Second)
	}

	record, ok := s.records.Load(p.RecordID)
	if !ok {
		return nil, interceptors.MakeNotFound(fmt.Errorf("record with ID %s not found", p.RecordID))
	}
	result := record.(*interceptors.CreateResult)
	return &interceptors.GetResult{
		ID:     result.ID,
		Value:  result.Value,
		Tenant: result.Tenant,
		Status: 200,
	}, nil
}

// Create generates a new record demonstrating the logging and tenant interceptors.
// It shows how interceptors can access and utilize context values (like tenant ID)
// throughout the request lifecycle.
func (s *interceptorssrvc) Create(ctx context.Context, p *interceptors.CreatePayload) (*interceptors.CreateResult, error) {
	log.Printf(ctx, "[Create] Creating record for tenant: %s with value: %s", p.TenantID, p.Value)
	id := interceptors.UUID(uuid.New().String())
	result := &interceptors.CreateResult{
		ID:     id,
		Value:  fmt.Sprintf("Created record with value: %s", p.Value),
		Tenant: string(p.TenantID),
		Status: 201,
	}

	s.records.Store(id, result)
	log.Printf(ctx, "[Create] Successfully created record with ID: %s", id)

	return result, nil
}

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

// Stream demonstrates the streaming interceptor by sending a stream of records and accepting new records to
// be created. It shows how interceptors can access and utilize context values (like trace ID and span ID)
// throughout the request lifecycle.
func (s *interceptorssrvc) Stream(ctx context.Context, p *interceptors.StreamPayload, stream interceptors.StreamServerStream) error {
	log.Printf(ctx, "[Stream] Streaming records for tenant: %s", p.TenantID)
	defer stream.Close()

	var err error
	s.records.Range(func(key, value any) bool {
		result := value.(*interceptors.CreateResult)
		log.Printf(ctx, "[Stream] Sending stream result: %s", result.ID)
		err = stream.SendWithContext(ctx, &interceptors.StreamResult{
			ID:     result.ID,
			Value:  result.Value,
			Tenant: result.Tenant,
			Status: 200,
		})
		return err == nil
	})
	if err != nil {
		log.Printf(ctx, "[Stream] Error sending stream result: %s", err)
		return err
	}

	errCh := make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				payload, err := stream.RecvWithContext(ctx)
				if err != nil {
					log.Printf(ctx, "[Stream] Error receiving stream payload: %s", err)
					errCh <- err
					return
				}
				log.Printf(ctx, "[Stream] Received stream payload: %s", payload.ID)

				result := &interceptors.CreateResult{
					ID:     payload.ID,
					Value:  payload.Value,
					Tenant: string(p.TenantID),
					Status: 201,
				}

				s.records.Store(result.ID, result)
				log.Printf(ctx, "[Stream] Sending stream result: %s", result.ID)
				err = stream.SendWithContext(ctx, &interceptors.StreamResult{
					ID:     result.ID,
					Value:  result.Value,
					Tenant: result.Tenant,
					Status: 200,
				})
				if err != nil {
					log.Printf(ctx, "[Stream] Error sending stream result: %s", err)
					errCh <- err
					return
				}
			}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			log.Printf(ctx, "[Stream] Context done")
			return nil
		case err := <-errCh:
			return err
		}
	}
}

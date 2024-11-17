// Code generated by goa v3.19.1, DO NOT EDIT.
//
// interceptors service
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package interceptors

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The interceptors service demonstrates a comprehensive use of interceptors
// combining
// authentication, tenant validation, caching, audit logging, and retry
// mechanisms. It showcases
// both client-side and server-side interceptors working together to provide a
// robust service.
type Service interface {
	// Get retrieves a record by ID with all interceptors in action
	Get(context.Context, *GetPayload) (res *GetResult, err error)
	// Create a new record with all interceptors in action
	Create(context.Context, *CreatePayload) (res *CreateResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "interceptors"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "interceptors"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"get", "create"}

// CreatePayload is the payload type of the interceptors service create method.
type CreatePayload struct {
	// Tenant ID for the request
	TenantID UUID
	// Value to store in record
	Value string
	// JWT auth token
	Auth string
	// Unique trace ID for request, initialized by the TraceRequest interceptor
	TraceID *UUID
	// Unique span ID for request, initialized by the TraceRequest interceptor
	SpanID *UUID
}

// CreateResult is the result type of the interceptors service create method.
type CreateResult struct {
	// ID of the created record
	ID UUID
	// Value of the record
	Value string
	// Tenant the record belongs to
	Tenant string
	// Response status code
	Status int
	// Timestamp when processed
	ProcessedAt string
	// Processing duration in ms
	Duration int
	// Number of retry attempts made
	RetryCount *int
	// Total time spent retrying
	RetryDuration *int
}

// GetPayload is the payload type of the interceptors service get method.
type GetPayload struct {
	// Tenant ID for the request
	TenantID UUID
	// ID of the record to retrieve
	RecordID UUID
	// JWT auth token
	Auth string
	// Unique trace ID for request, initialized by the TraceRequest interceptor
	TraceID *UUID
	// Unique span ID for request, initialized by the TraceRequest interceptor
	SpanID *UUID
}

// GetResult is the result type of the interceptors service get method.
type GetResult struct {
	// ID of the record
	ID UUID
	// Value of the record
	Value string
	// Tenant the record belongs to
	Tenant string
	// Response status code
	Status int
	// Timestamp when processed, written by the RequestAudit interceptor
	ProcessedAt string
	// Processing duration in ms, written by the RequestAudit interceptor
	Duration int
	// Time at which the record was cached, written by the Cache interceptor
	CachedAt *string
	// Number of retry attempts made, written client-side by the Retry interceptor
	RetryCount *int
	// Total time spent retrying, written client-side by the Retry interceptor
	RetryDuration *int
}

// Valid UUID representation as per RFC 4122
type UUID string

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "NotFound", false, false, false)
}

// MakeUnavailable builds a goa.ServiceError from an error.
func MakeUnavailable(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Unavailable", false, true, false)
}

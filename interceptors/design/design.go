package design

import (
	. "goa.design/goa/v3/dsl"
)

// EncodeTenant is a client-side interceptor that writes the tenant ID to the JWT
var EncodeTenant = Interceptor("EncodeTenant", func() {
	Description("Client-side interceptor which writes the tenant ID to the signed JWT contained in the Authorization header")

	ReadPayload(func() {
		Attribute("tenantID", String, "Tenant ID to encode")
	})

	WritePayload(func() {
		Attribute("auth", String, "Generated JWT auth token")
	})
})

// JWTAuth handles authentication and authorization
var JWTAuth = Interceptor("JWTAuth", func() {
	Description("Server-side interceptor that validates JWT token and tenant ID")

	ReadPayload(func() {
		Attribute("auth", String, "JWT auth token")
		Attribute("tenantID", String, "Tenant ID to validate against")
	})
})

// TraceRequest handles request tracing
var TraceRequest = Interceptor("TraceRequest", func() {
	Description("Server-side interceptor that adds trace context to the request payload")

	WritePayload(func() {
		Attribute("traceID", String, "Unique trace ID for request")
		Attribute("spanID", String, "Unique span ID for request")
	})
})

// RequestAudit handles request/response logging and timing
var RequestAudit = Interceptor("RequestAudit", func() {
	Description("Server-side interceptor that provides comprehensive request/response audit logging")

	ReadResult(func() {
		Attribute("status", Int, "Response status code")
	})

	WriteResult(func() {
		Attribute("processedAt", String, "Timestamp when processed", func() {
			Format(FormatDateTime)
		})
		Attribute("duration", Int, "Processing duration in ms")
	})
})

// Cache implements a transparent cache for loaded records
var Cache = Interceptor("Cache", func() {
	Description("Server-side interceptor which implements a transparent cache for loaded records")

	ReadPayload(func() {
		Attribute("recordID", String, "Record ID to cache")
	})

	WriteResult(func() {
		Attribute("cachedAt", String, "Time at which the record was cached in RFC3339 format")
	})
})

// SetDeadline sets request context deadline
var SetDeadline = Interceptor("SetDeadline", func() {
	Description("Server-side interceptor which sets the context deadline for the request")
})

// Retry implements smart retry logic for failed requests
var Retry = Interceptor("Retry", func() {
	Description("Client-side interceptor which implements smart retry logic with exponential backoff")

	ReadResult(func() {
		Attribute("status", Int, "Response status code")
	})

	WriteResult(func() {
		Attribute("retryCount", Int, "Number of retry attempts made")
		Attribute("retryDuration", Int, "Total time spent retrying in ms")
	})
})

// Service definition
var _ = Service("interceptors", func() {
	Description(`The interceptors service demonstrates a comprehensive use of interceptors combining
authentication, tenant validation, caching, audit logging, and retry mechanisms. It showcases
both client-side and server-side interceptors working together to provide a robust service.`)

	// Server-side interceptors - ordered by execution sequence
	ServerInterceptor(TraceRequest) // Add trace context first
	ServerInterceptor(RequestAudit) // Start timing
	ServerInterceptor(JWTAuth)      // Then validate authentication
	ServerInterceptor(SetDeadline)  // Set deadline

	// Client-side interceptors
	ClientInterceptor(EncodeTenant)
	ClientInterceptor(Retry)

	// Define a method that uses all the interceptors
	Method("get", func() {
		Description("Get retrieves a record by ID with all interceptors in action")

		// Add cache interceptor only for get method
		ServerInterceptor(Cache)

		Payload(func() {
			Field(1, "tenantID", UUID, "Tenant ID for the request")
			Field(2, "recordID", UUID, "ID of the record to retrieve")
			Field(3, "auth", String, "JWT auth token")
			Field(4, "traceID", UUID, "Unique trace ID for request, initialized by the TraceRequest interceptor")
			Field(5, "spanID", UUID, "Unique span ID for request, initialized by the TraceRequest interceptor")
			Required("recordID", "tenantID", "auth")
		})

		Result(func() {
			Field(1, "id", UUID, "ID of the record")
			Field(2, "value", String, "Value of the record")
			Field(3, "tenant", String, "Tenant the record belongs to")
			Field(4, "status", Int, "Response status code")
			Field(5, "processedAt", String, "Timestamp when processed, written by the RequestAudit interceptor")
			Field(6, "duration", Int, "Processing duration in ms, written by the RequestAudit interceptor")
			Field(7, "cachedAt", String, "Time at which the record was cached, written by the Cache interceptor")
			Field(8, "retryCount", Int, "Number of retry attempts made, written client-side by the Retry interceptor")
			Field(9, "retryDuration", Int, "Total time spent retrying, written client-side by the Retry interceptor")
			Required("id", "value", "tenant", "status", "processedAt", "duration")
		})

		Error("NotFound", ErrorResult, "Record not found")
		Error("Unavailable", ErrorResult, "Temporary error", func() {
			Temporary()
		})

		HTTP(func() {
			GET("/records/{tenantID}/{recordID}")
			Header("auth:Authorization")
			Response("NotFound", StatusNotFound)
			Response("Unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response("NotFound", CodeNotFound)
			Response("Unavailable", CodeUnavailable)
		})
	})

	// Additional method showcasing write operations
	Method("create", func() {
		Description("Create a new record with all interceptors in action")

		Payload(func() {
			Field(1, "tenantID", UUID, "Tenant ID for the request")
			Field(2, "value", String, "Value to store in record")
			Field(3, "auth", String, "JWT auth token")
			Field(4, "traceID", UUID, "Unique trace ID for request, initialized by the TraceRequest interceptor")
			Field(5, "spanID", UUID, "Unique span ID for request, initialized by the TraceRequest interceptor")
			Required("value", "tenantID", "auth")
		})

		Result(func() {
			Field(1, "id", UUID, "ID of the created record")
			Field(2, "value", String, "Value of the record")
			Field(3, "tenant", String, "Tenant the record belongs to")
			Field(4, "status", Int, "Response status code")
			Field(5, "processedAt", String, "Timestamp when processed")
			Field(6, "duration", Int, "Processing duration in ms")
			Field(7, "retryCount", Int, "Number of retry attempts made")
			Field(8, "retryDuration", Int, "Total time spent retrying")
			Required("id", "value", "tenant", "status", "processedAt", "duration")
		})

		HTTP(func() {
			POST("/records/{tenantID}")
			Header("auth:Authorization")
			Response(StatusCreated)
		})
	})
})

var UUID = Type("UUID", String, func() {
	Description("Valid UUID representation as per RFC 4122")
	Format(FormatUUID)
})

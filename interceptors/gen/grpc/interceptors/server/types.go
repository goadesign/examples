// Code generated by goa v3.19.1, DO NOT EDIT.
//
// interceptors gRPC server types
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package server

import (
	interceptorspb "goa.design/examples/interceptors/gen/grpc/interceptors/pb"
	interceptors "goa.design/examples/interceptors/gen/interceptors"
	goa "goa.design/goa/v3/pkg"
)

// NewGetPayload builds the payload of the "get" endpoint of the "interceptors"
// service from the gRPC request type.
func NewGetPayload(message *interceptorspb.GetRequest) *interceptors.GetPayload {
	v := &interceptors.GetPayload{
		TenantID: interceptors.UUID(message.TenantId),
		RecordID: interceptors.UUID(message.RecordId),
		Auth:     message.Auth,
	}
	if message.TraceId != nil {
		traceID := interceptors.UUID(*message.TraceId)
		v.TraceID = &traceID
	}
	if message.SpanId != nil {
		spanID := interceptors.UUID(*message.SpanId)
		v.SpanID = &spanID
	}
	return v
}

// NewProtoGetResponse builds the gRPC response type from the result of the
// "get" endpoint of the "interceptors" service.
func NewProtoGetResponse(result *interceptors.GetResult) *interceptorspb.GetResponse {
	message := &interceptorspb.GetResponse{
		Id:          string(result.ID),
		Value:       result.Value,
		Tenant:      result.Tenant,
		Status:      int32(result.Status),
		ProcessedAt: result.ProcessedAt,
		Duration:    int32(result.Duration),
		CachedAt:    result.CachedAt,
	}
	if result.RetryCount != nil {
		retryCount := int32(*result.RetryCount)
		message.RetryCount = &retryCount
	}
	if result.RetryDuration != nil {
		retryDuration := int32(*result.RetryDuration)
		message.RetryDuration = &retryDuration
	}
	return message
}

// ValidateGetRequest runs the validations defined on GetRequest.
func ValidateGetRequest(message *interceptorspb.GetRequest) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.tenantID", string(message.TenantId), goa.FormatUUID))
	err = goa.MergeErrors(err, goa.ValidateFormat("message.recordID", string(message.RecordId), goa.FormatUUID))
	if message.TraceId != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("message.traceID", string(*message.TraceId), goa.FormatUUID))
	}
	if message.SpanId != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("message.spanID", string(*message.SpanId), goa.FormatUUID))
	}
	return
}

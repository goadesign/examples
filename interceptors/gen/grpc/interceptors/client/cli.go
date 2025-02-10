// Code generated by goa v3.19.1, DO NOT EDIT.
//
// interceptors gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package client

import (
	"encoding/json"
	"fmt"

	interceptorspb "goa.design/examples/interceptors/gen/grpc/interceptors/pb"
	interceptors "goa.design/examples/interceptors/gen/interceptors"
)

// BuildGetPayload builds the payload for the interceptors get endpoint from
// CLI flags.
func BuildGetPayload(interceptorsGetMessage string) (*interceptors.GetPayload, error) {
	var err error
	var message interceptorspb.GetRequest
	{
		if interceptorsGetMessage != "" {
			err = json.Unmarshal([]byte(interceptorsGetMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"auth\": \"Voluptas rerum non iste ducimus et recusandae.\",\n      \"recordID\": \"Aliquid necessitatibus asperiores iure.\",\n      \"spanID\": \"Aliquid necessitatibus asperiores iure.\",\n      \"tenantID\": \"Aliquid necessitatibus asperiores iure.\",\n      \"traceID\": \"Aliquid necessitatibus asperiores iure.\"\n   }'")
			}
		}
	}
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

	return v, nil
}

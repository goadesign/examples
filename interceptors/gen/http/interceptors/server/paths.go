// Code generated by goa v3.19.1, DO NOT EDIT.
//
// HTTP request path constructors for the interceptors service.
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package server

import (
	"fmt"
)

// GetInterceptorsPath returns the URL path to the interceptors service get HTTP endpoint.
func GetInterceptorsPath(tenantID string, recordID string) string {
	return fmt.Sprintf("/records/%v/%v", tenantID, recordID)
}

// CreateInterceptorsPath returns the URL path to the interceptors service create HTTP endpoint.
func CreateInterceptorsPath(tenantID string) string {
	return fmt.Sprintf("/records/%v", tenantID)
}

// StreamInterceptorsPath returns the URL path to the interceptors service stream HTTP endpoint.
func StreamInterceptorsPath(tenantID string) string {
	return fmt.Sprintf("/records/%v/stream", tenantID)
}

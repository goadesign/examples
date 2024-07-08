// Code generated by goa v3.17.2, DO NOT EDIT.
//
// session service
//
// Command:
// $ goa gen goa.design/examples/cookies/design

package session

import (
	"context"
)

// The session service illustrates managing user sessions with cookies.
type Service interface {
	// CreateSession implements create_session.
	CreateSession(context.Context, *CreateSessionPayload) (res *CreateSessionResult, err error)
	// UseSession implements use_session.
	UseSession(context.Context, *UseSessionPayload) (res *UseSessionResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "session"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "session"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"create_session", "use_session"}

// CreateSessionPayload is the payload type of the session service
// create_session method.
type CreateSessionPayload struct {
	// Name of session
	Name string
}

// CreateSessionResult is the result type of the session service create_session
// method.
type CreateSessionResult struct {
	// Session ID
	SessionID string
	// User message
	Message string
}

// UseSessionPayload is the payload type of the session service use_session
// method.
type UseSessionPayload struct {
	// Session ID
	SessionID string
}

// UseSessionResult is the result type of the session service use_session
// method.
type UseSessionResult struct {
	// User message
	Message string
}

// Code generated by goa v3.2.3, DO NOT EDIT.
//
// session HTTP server types
//
// Command:
// $ goa gen goa.design/examples/cookies/design -o
// $(GOPATH)/src/goa.design/examples/cookies

package server

import (
	session "goa.design/examples/cookies/gen/session"
	goa "goa.design/goa/v3/pkg"
)

// CreateSessionRequestBody is the type of the "session" service
// "create_session" endpoint HTTP request body.
type CreateSessionRequestBody struct {
	// Name of session
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CreateSessionResponseBody is the type of the "session" service
// "create_session" endpoint HTTP response body.
type CreateSessionResponseBody struct {
	// User message
	Message string `form:"message" json:"message" xml:"message"`
}

// UseSessionResponseBody is the type of the "session" service "use_session"
// endpoint HTTP response body.
type UseSessionResponseBody struct {
	// User message
	Message string `form:"message" json:"message" xml:"message"`
}

// NewCreateSessionResponseBody builds the HTTP response body from the result
// of the "create_session" endpoint of the "session" service.
func NewCreateSessionResponseBody(res *session.CreateSessionResult) *CreateSessionResponseBody {
	body := &CreateSessionResponseBody{
		Message: res.Message,
	}
	return body
}

// NewUseSessionResponseBody builds the HTTP response body from the result of
// the "use_session" endpoint of the "session" service.
func NewUseSessionResponseBody(res *session.UseSessionResult) *UseSessionResponseBody {
	body := &UseSessionResponseBody{
		Message: res.Message,
	}
	return body
}

// NewCreateSessionPayload builds a session service create_session endpoint
// payload.
func NewCreateSessionPayload(body *CreateSessionRequestBody) *session.CreateSessionPayload {
	v := &session.CreateSessionPayload{
		Name: *body.Name,
	}

	return v
}

// NewUseSessionPayload builds a session service use_session endpoint payload.
func NewUseSessionPayload(sessionID string) *session.UseSessionPayload {
	v := &session.UseSessionPayload{}
	v.SessionID = sessionID

	return v
}

// ValidateCreateSessionRequestBody runs the validations defined on
// create_session_request_body
func ValidateCreateSessionRequestBody(body *CreateSessionRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

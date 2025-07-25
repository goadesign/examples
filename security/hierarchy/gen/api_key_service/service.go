// Code generated by goa v3.21.5, DO NOT EDIT.
//
// api_key_service service
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package apikeyservice

import (
	"context"

	"goa.design/goa/v3/security"
)

// The svc service is secured with API key based authentication
type Service interface {
	// Default implements default.
	Default(context.Context, *DefaultPayload) (err error)
	// This method requires a valid JWT token.
	Secure(context.Context, *SecurePayload) (err error)
	// This method is not secured.
	Unsecure(context.Context) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// APIKeyAuth implements the authorization logic for the APIKey security scheme.
	APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error)
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "hierarchy"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "api_key_service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"default", "secure", "unsecure"}

// DefaultPayload is the payload type of the api_key_service service default
// method.
type DefaultPayload struct {
	// API key used for authentication
	Key string
}

// SecurePayload is the payload type of the api_key_service service secure
// method.
type SecurePayload struct {
	// JWT used for authentication
	Token string
}

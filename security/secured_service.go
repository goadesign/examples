package multiauth

import (
	"context"
	"fmt"
	"log"

	securedservice "goa.design/examples/security/gen/secured_service"
	"goa.design/goa/v3/security"
)

// secured_service service example implementation.
// The example methods log the requests and return zero values.
type securedServicesrvc struct {
	logger *log.Logger
}

// NewSecuredService returns the secured_service service implementation.
func NewSecuredService(logger *log.Logger) securedservice.Service {
	return &securedServicesrvc{logger}
}

// BasicAuth implements the authorization logic for service "secured_service"
// for the "basic" security scheme.
func (s *securedServicesrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// JWTAuth implements the authorization logic for service "secured_service" for
// the "jwt" security scheme.
func (s *securedServicesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// APIKeyAuth implements the authorization logic for service "secured_service"
// for the "api_key" security scheme.
func (s *securedServicesrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// OAuth2Auth implements the authorization logic for service "secured_service"
// for the "oauth2" security scheme.
func (s *securedServicesrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// Creates a valid JWT
func (s *securedServicesrvc) Signin(ctx context.Context, p *securedservice.SigninPayload) (res *securedservice.Creds, err error) {
	res = &securedservice.Creds{}
	s.logger.Print("securedService.signin")
	return
}

// This action is secured with the jwt scheme
func (s *securedServicesrvc) Secure(ctx context.Context, p *securedservice.SecurePayload) (res string, err error) {
	s.logger.Print("securedService.secure")
	return
}

// This action is secured with the jwt scheme and also requires an API key
// query string.
func (s *securedServicesrvc) DoublySecure(ctx context.Context, p *securedservice.DoublySecurePayload) (res string, err error) {
	s.logger.Print("securedService.doubly_secure")
	return
}

// This action is secured with the jwt scheme and also requires an API key
// header.
func (s *securedServicesrvc) AlsoDoublySecure(ctx context.Context, p *securedservice.AlsoDoublySecurePayload) (res string, err error) {
	s.logger.Print("securedService.also_doubly_secure")
	return
}

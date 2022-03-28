package hierarchicalsecurity

import (
	"context"
	"fmt"
	"log"

	apikeyservice "goa.design/examples/security/hierarchy/gen/api_key_service"
	"goa.design/goa/v3/security"
)

// api_key_service service example implementation.
// The example methods log the requests and return zero values.
type apiKeyServicesrvc struct {
	logger *log.Logger
}

// NewAPIKeyService returns the api_key_service service implementation.
func NewAPIKeyService(logger *log.Logger) apikeyservice.Service {
	return &apiKeyServicesrvc{logger}
}

// APIKeyAuth implements the authorization logic for service "api_key_service"
// for the "api_key" security scheme.
func (s *apiKeyServicesrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
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

// JWTAuth implements the authorization logic for service "api_key_service" for
// the "jwt" security scheme.
func (s *apiKeyServicesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// Default implements default.
func (s *apiKeyServicesrvc) Default(ctx context.Context, p *apikeyservice.DefaultPayload) (err error) {
	s.logger.Print("apiKeyService.default")
	return
}

// This method requires a valid JWT token.
func (s *apiKeyServicesrvc) Secure(ctx context.Context, p *apikeyservice.SecurePayload) (err error) {
	s.logger.Print("apiKeyService.secure")
	return
}

// This method is not secured.
func (s *apiKeyServicesrvc) Unsecure(ctx context.Context) (err error) {
	s.logger.Print("apiKeyService.unsecure")
	return
}

package hierarchicalsecurity

import (
	"context"
	"fmt"
	"log"

	defaultservice "goa.design/examples/security/hierarchy/gen/default_service"
	"goa.design/goa/v3/security"
)

// default_service service example implementation.
// The example methods log the requests and return zero values.
type defaultServicesrvc struct {
	logger *log.Logger
}

// NewDefaultService returns the default_service service implementation.
func NewDefaultService(logger *log.Logger) defaultservice.Service {
	return &defaultServicesrvc{logger}
}

// BasicAuth implements the authorization logic for service "default_service"
// for the "basic" security scheme.
func (s *defaultServicesrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
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

// The default service default_method is secured with basic authentication
func (s *defaultServicesrvc) Default(ctx context.Context, p *defaultservice.DefaultPayload) (err error) {
	s.logger.Print("defaultService.default")
	return
}

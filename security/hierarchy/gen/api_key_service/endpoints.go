// Code generated by goa v3.13.1, DO NOT EDIT.
//
// api_key_service endpoints
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design -o security/hierarchy

package apikeyservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "api_key_service" service endpoints.
type Endpoints struct {
	Default  goa.Endpoint
	Secure   goa.Endpoint
	Unsecure goa.Endpoint
}

// NewEndpoints wraps the methods of the "api_key_service" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Default:  NewDefaultEndpoint(s, a.APIKeyAuth),
		Secure:   NewSecureEndpoint(s, a.JWTAuth),
		Unsecure: NewUnsecureEndpoint(s),
	}
}

// Use applies the given middleware to all the "api_key_service" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Default = m(e.Default)
	e.Secure = m(e.Secure)
	e.Unsecure = m(e.Unsecure)
}

// NewDefaultEndpoint returns an endpoint function that calls the method
// "default" of service "api_key_service".
func NewDefaultEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DefaultPayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, p.Key, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Default(ctx, p)
	}
}

// NewSecureEndpoint returns an endpoint function that calls the method
// "secure" of service "api_key_service".
func NewSecureEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SecurePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Secure(ctx, p)
	}
}

// NewUnsecureEndpoint returns an endpoint function that calls the method
// "unsecure" of service "api_key_service".
func NewUnsecureEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return nil, s.Unsecure(ctx)
	}
}

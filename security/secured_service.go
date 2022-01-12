package multiauth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

var (
	// ErrUnauthorized is the error returned by Login when the request credentials
	// are invalid.
	ErrUnauthorized error = securedservice.Unauthorized("invalid username and password combination")

	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = securedservice.Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
	ErrInvalidTokenScopes error = securedservice.InvalidScopes("invalid scopes in token")

	// Key is the key used in JWT authentication
	Key = []byte("secret")
)

// BasicAuth implements the authorization logic for service "secured_service"
// for the "basic" security scheme.
func (s *securedServicesrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	if user != "goa" {
		return ctx, ErrUnauthorized
	}
	if pass != "rocks" {
		return ctx, ErrUnauthorized
	}
	ctx = contextWithAuthInfo(ctx, authInfo{
		user: user,
	})
	return ctx, nil
}

// JWTAuth implements the authorization logic for service "secured_service" for
// the "jwt" security scheme.
func (s *securedServicesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) { return Key, nil })
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, securedservice.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		claims: claims,
	})
	return ctx, nil
}

// APIKeyAuth implements the authorization logic for service "secured_service"
// for the "api_key" security scheme.
func (s *securedServicesrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	if key != "my_awesome_api_key" {
		return ctx, ErrUnauthorized
	}
	ctx = contextWithAuthInfo(ctx, authInfo{
		key: key,
	})
	return ctx, nil
}

// OAuth2Auth implements the authorization logic for service "secured_service"
// for the "oauth2" security scheme.
func (s *securedServicesrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) { return Key, nil })
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, securedservice.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		claims: claims,
	})
	return ctx, nil
}

// Creates a valid JWT
func (s *securedServicesrvc) Signin(ctx context.Context, p *securedservice.SigninPayload) (res *securedservice.Creds, err error) {
	// create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":    time.Now().Unix(),
		"scopes": []string{"api:read", "api:write"},
	})

	s.logger.Printf("user '%s' logged in", p.Username)

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	t, err := token.SignedString(Key)
	if err != nil {
		return nil, err
	}
	return &securedservice.Creds{
		JWT:        t,
		OauthToken: t,
		APIKey:     "my_awesome_api_key",
	}, nil
}

// This action is secured with the jwt scheme
func (s *securedServicesrvc) Secure(ctx context.Context, p *securedservice.SecurePayload) (res string, err error) {
	res = fmt.Sprintf("User authorized using JWT token %q", p.Token)
	authInfo := contextAuthInfo(ctx)
	s.logger.Printf(res)
	s.logger.Printf("%v", authInfo)
	if p.Fail != nil && *p.Fail {
		s.logger.Printf("Uh oh! `fail` passed in parameter. Auth failed!")
		return "", securedservice.Unauthorized("forced authentication failure")
	}
	return
}

// This action is secured with the jwt scheme and also requires an API key
// query string.
func (s *securedServicesrvc) DoublySecure(ctx context.Context, p *securedservice.DoublySecurePayload) (res string, err error) {
	res = fmt.Sprintf("User authorized using JWT token %q and API Key %q", p.Token, p.Key)
	authInfo := contextAuthInfo(ctx)
	s.logger.Printf(res)
	s.logger.Printf("%v", authInfo)
	return
}

// This action is secured with the jwt scheme and also requires an API key
// header.
func (s *securedServicesrvc) AlsoDoublySecure(ctx context.Context, p *securedservice.AlsoDoublySecurePayload) (res string, err error) {
	if p.Username != nil && p.Password != nil && p.OauthToken != nil {
		res = fmt.Sprintf("User authorized using username %q/password %q and OAuth2 token %q", *p.Username, *p.Password, *p.OauthToken)
		s.logger.Printf(res)
		return
	}
	res = fmt.Sprintf("User authorized using JWT token %q and API Key %q", *p.Token, *p.Key)
	authInfo := contextAuthInfo(ctx)
	s.logger.Print(res)
	s.logger.Printf("%v", authInfo)
	return
}

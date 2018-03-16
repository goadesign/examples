package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/satori/go.uuid"
)

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
//
// Note: the code below assumes the example is compiled against the master branch of goa.
// If compiling against goa v1 the call to jwt.New needs to be:
//
//    middleware := jwt.New(keys, ForceFail(), app.NewJWTSecurity())
func NewJWTMiddleware() (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(jwt.NewSimpleResolver(keys), ForceFail(), app.NewJWTSecurity()), nil
}

// JWTController implements the JWT resource.
type JWTController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewJWTController creates a JWT controller.
func NewJWTController(service *goa.Service) (*JWTController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &JWTController{
		Controller: service.NewController("JWTController"),
		privateKey: privKey,
	}, nil
}

// Signin creates JWTs for use by clients to access the secure endpoints.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// Generate JWT
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss":    "Issuer",                         // who creates the token and signs it
		"aud":    "Audience",                       // to whom the token is intended to be sent
		"exp":    in10m,                            // time when the token will expire (10 minutes from now)
		"jti":    uuid.Must(uuid.NewV4()).String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),                // when the token was issued/created (now)
		"nbf":    2,                                // time before which the token is not yet valid (2 minutes ago)
		"sub":    "subject",                        // the subject/principal is whom the token is about
		"scopes": "api:access",                     // token scope - not a standard claim
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	return ctx.NoContent()
}

// Secure runs the secure action.
func (c *JWTController) Secure(ctx *app.SecureJWTContext) error {
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}
	claims := token.Claims.(jwtgo.MapClaims)

	// Use the claims to authorize
	subject := claims["sub"]
	if subject != "subject" {
		// A real app would probably use an "Unauthorized" response here
		res := &app.Success{OK: false}
		return ctx.OK(res)
	}

	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *JWTController) Unsecure(ctx *app.UnsecureJWTContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() ([]jwt.Key, error) {
	keyFiles, err := filepath.Glob("./jwtkey/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]jwt.Key, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}

// ForceFail is a middleware illustrating the use of validation middleware with JWT auth.  It checks
// for the presence of a "fail" query string and fails validation if set to the value "true".
func ForceFail() goa.Middleware {
	errValidationFailed := goa.NewErrorClass("validation_failed", 401)
	forceFail := func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			if f, ok := req.URL.Query()["fail"]; ok {
				if f[0] == "true" {
					return errValidationFailed("forcing failure to illustrate Validation middleware")
				}
			}
			return h(ctx, rw, req)
		}
	}
	fm, _ := goa.NewMiddleware(forceFail)
	return fm
}

package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"golang.org/x/net/context"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
)

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
func NewJWTMiddleware() (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	middleware := jwt.New(keys, ForceFail(), app.NewJWTSecurity())
	return middleware, nil
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

// Signin creates JWTs for use by clients to access the secured endpoints.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// Generate JWT
	token := jwtgo.New(jwtgo.SigningMethodRS256)
	token.Claims = map[string]interface{}{"claim1": true}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return err // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	return ctx.NoContent()
}

// Secured runs the secured action.
func (c *JWTController) Secured(ctx *app.SecuredJWTContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *JWTController) Unsecured(ctx *app.UnsecuredJWTContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() ([]*rsa.PublicKey, error) {
	keyFiles, err := filepath.Glob("./jwtkeys/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]*rsa.PublicKey, len(keyFiles))
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
	fm, err := goa.NewMiddleware(forceFail)
	if err != nil {
		panic(err) // bug
	}
	return fm
}

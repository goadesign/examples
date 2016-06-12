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
	middleware := jwt.New(keys, fm, app.NewJWTSecurity())
	if err != nil {
		return nil, err
	}
	return middleware, nil
}

// JWTEndpointsController implements the JWTEndpoints resource.
type JWTEndpointsController struct {
	*goa.Controller
}

// NewJWTEndpointsController creates a JWTEndpoints controller.
func NewJWTEndpointsController(service *goa.Service) *JWTEndpointsController {
	return &JWTEndpointsController{Controller: service.NewController("JWTEndpointsController")}
}

// Secured runs the secured action.
func (c *JWTEndpointsController) Secured(ctx *app.SecuredJWTEndpointsContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *JWTEndpointsController) Unsecured(ctx *app.UnsecuredJWTEndpointsContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() ([]*rsa.PublicKey, error) {
	keyFiles, err := filepath.Glob("./pubkeys/*.pub")
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
		return nil, fmt.Errorf("couldn't load any public keys for security")
	}

	return keys, nil
}

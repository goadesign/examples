package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
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
	middleware := jwt.New(keys, nil, app.NewJWTSecurity())
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
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *JWTEndpointsController) Unsecured(ctx *app.UnsecuredJWTEndpointsContext) error {
	// TBD: implement
	res := &app.Success{}
	return ctx.OK(res)
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() ([]*rsa.PublicKey, error) {
	keyFiles, err := filepath.Glob("./pubkeys/*.pub")
	if err != nil {
		return err
	}
	keys := make([]*rsa.PublicKey, len(keyfiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return err
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load any public keys for security")
	}

	return keys, nil
}

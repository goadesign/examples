package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jws"

	goaclient "github.com/goadesign/goa/client"
)

type (
	// SASource is a token source that creates JWT token from the credentials stored in a Google
	// Developers service account JSON key file.
	SASource struct {
		RSA *rsa.PrivateKey
	}
)

// NewSASource creates a service account JWT token source from the credentials stored in the given
// Google Developers service account JSON key file.
func NewSASource(safile string) (goaclient.TokenSource, error) {
	sa, err := ioutil.ReadFile(safile)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(sa)
	if err != nil {
		return nil, fmt.Errorf("failed to load service account Key JSON file: %s", err)
	}
	b := conf.PrivateKey
	block, _ := pem.Decode(b)
	if block != nil {
		b = block.Bytes
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(b)
	if err != nil {
		parsedKey, err = x509.ParsePKCS1PrivateKey(b)
		if err != nil {
			return nil, fmt.Errorf("private key should be a PEM or plain PKSC1 or PKCS8; parse error: %v", err)
		}
	}
	rsa, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is invalid")
	}
	return &SASource{rsa}, nil
}

// Token returns a JWT token factory.
func (s *SASource) Token() (goaclient.Token, error) {
	return s, nil
}

// SetAuthHeader signs the request with a JWT token created from the token source.
func (s *SASource) SetAuthHeader(r *http.Request) {
	iat := time.Now()
	exp := iat.Add(time.Hour)
	jwt := &jws.ClaimSet{
		Iss:   "client.goa-endpoints.appspot.com",
		Sub:   "client@goa.design",
		Aud:   "goa-endpoints.appspot.com",
		Scope: "email",
		Iat:   iat.Unix(),
		Exp:   exp.Unix(),
	}
	header := &jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
	}
	auth, err := jws.Encode(header, jwt, s.RSA)
	if err != nil {
		fmt.Printf("Whoops, JWT encoding failed: %s\n", err)
		return
	}
	r.Header.Set("Authorization", "Bearer "+auth)
}

// Valid returns true.
func (s *SASource) Valid() bool {
	return true
}

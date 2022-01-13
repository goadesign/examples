// Code generated by goa v3.5.4, DO NOT EDIT.
//
// secured_service HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package client

import (
	"fmt"
	"strconv"

	securedservice "goa.design/examples/security/gen/secured_service"
)

// BuildSigninPayload builds the payload for the secured_service signin
// endpoint from CLI flags.
func BuildSigninPayload(securedServiceSigninUsername string, securedServiceSigninPassword string) (*securedservice.SigninPayload, error) {
	var username string
	{
		username = securedServiceSigninUsername
	}
	var password string
	{
		password = securedServiceSigninPassword
	}
	v := &securedservice.SigninPayload{}
	v.Username = username
	v.Password = password

	return v, nil
}

// BuildSecurePayload builds the payload for the secured_service secure
// endpoint from CLI flags.
func BuildSecurePayload(securedServiceSecureFail string, securedServiceSecureToken string) (*securedservice.SecurePayload, error) {
	var err error
	var fail *bool
	{
		if securedServiceSecureFail != "" {
			var val bool
			val, err = strconv.ParseBool(securedServiceSecureFail)
			fail = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for fail, must be BOOL")
			}
		}
	}
	var token string
	{
		token = securedServiceSecureToken
	}
	v := &securedservice.SecurePayload{}
	v.Fail = fail
	v.Token = token

	return v, nil
}

// BuildDoublySecurePayload builds the payload for the secured_service
// doubly_secure endpoint from CLI flags.
func BuildDoublySecurePayload(securedServiceDoublySecureKey string, securedServiceDoublySecureToken string) (*securedservice.DoublySecurePayload, error) {
	var key string
	{
		key = securedServiceDoublySecureKey
	}
	var token string
	{
		token = securedServiceDoublySecureToken
	}
	v := &securedservice.DoublySecurePayload{}
	v.Key = key
	v.Token = token

	return v, nil
}

// BuildAlsoDoublySecurePayload builds the payload for the secured_service
// also_doubly_secure endpoint from CLI flags.
func BuildAlsoDoublySecurePayload(securedServiceAlsoDoublySecureKey string, securedServiceAlsoDoublySecureOauthToken string, securedServiceAlsoDoublySecureToken string, securedServiceAlsoDoublySecureUsername string, securedServiceAlsoDoublySecurePassword string) (*securedservice.AlsoDoublySecurePayload, error) {
	var key *string
	{
		if securedServiceAlsoDoublySecureKey != "" {
			key = &securedServiceAlsoDoublySecureKey
		}
	}
	var oauthToken *string
	{
		if securedServiceAlsoDoublySecureOauthToken != "" {
			oauthToken = &securedServiceAlsoDoublySecureOauthToken
		}
	}
	var token *string
	{
		if securedServiceAlsoDoublySecureToken != "" {
			token = &securedServiceAlsoDoublySecureToken
		}
	}
	var username *string
	{
		if securedServiceAlsoDoublySecureUsername != "" {
			username = &securedServiceAlsoDoublySecureUsername
		}
	}
	var password *string
	{
		if securedServiceAlsoDoublySecurePassword != "" {
			password = &securedServiceAlsoDoublySecurePassword
		}
	}
	v := &securedservice.AlsoDoublySecurePayload{}
	v.Key = key
	v.OauthToken = oauthToken
	v.Token = token
	v.Username = username
	v.Password = password

	return v, nil
}

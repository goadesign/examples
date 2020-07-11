// Code generated by goa v2.2.0, DO NOT EDIT.
//
// secured_service gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package client

import (
	"encoding/json"
	"fmt"

	secured_servicepb "goa.design/examples/security/gen/grpc/secured_service/pb"
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
func BuildSecurePayload(securedServiceSecureMessage string, securedServiceSecureToken string) (*securedservice.SecurePayload, error) {
	var err error
	var message secured_servicepb.SecureRequest
	{
		if securedServiceSecureMessage != "" {
			err = json.Unmarshal([]byte(securedServiceSecureMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"fail\": true\n   }'")
			}
		}
	}
	var token string
	{
		token = securedServiceSecureToken
	}
	v := &securedservice.SecurePayload{
		Fail: &message.Fail,
	}
	v.Token = token

	return v, nil
}

// BuildDoublySecurePayload builds the payload for the secured_service
// doubly_secure endpoint from CLI flags.
func BuildDoublySecurePayload(securedServiceDoublySecureMessage string, securedServiceDoublySecureToken string) (*securedservice.DoublySecurePayload, error) {
	var err error
	var message secured_servicepb.DoublySecureRequest
	{
		if securedServiceDoublySecureMessage != "" {
			err = json.Unmarshal([]byte(securedServiceDoublySecureMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"key\": \"abcdef12345\"\n   }'")
			}
		}
	}
	var token string
	{
		token = securedServiceDoublySecureToken
	}
	v := &securedservice.DoublySecurePayload{
		Key: message.Key,
	}
	v.Token = token

	return v, nil
}

// BuildAlsoDoublySecurePayload builds the payload for the secured_service
// also_doubly_secure endpoint from CLI flags.
func BuildAlsoDoublySecurePayload(securedServiceAlsoDoublySecureMessage string, securedServiceAlsoDoublySecureOauthToken string, securedServiceAlsoDoublySecureToken string) (*securedservice.AlsoDoublySecurePayload, error) {
	var err error
	var message secured_servicepb.AlsoDoublySecureRequest
	{
		if securedServiceAlsoDoublySecureMessage != "" {
			err = json.Unmarshal([]byte(securedServiceAlsoDoublySecureMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"key\": \"abcdef12345\",\n      \"password\": \"password\",\n      \"username\": \"user\"\n   }'")
			}
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
	v := &securedservice.AlsoDoublySecurePayload{}
	if message.Username != "" {
		v.Username = &message.Username
	}
	if message.Password != "" {
		v.Password = &message.Password
	}
	if message.Key != "" {
		v.Key = &message.Key
	}
	v.OauthToken = oauthToken
	v.Token = token

	return v, nil
}

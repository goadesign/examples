// Code generated by goa v3.5.4, DO NOT EDIT.
//
// secured_service gRPC server types
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package server

import (
	secured_servicepb "goa.design/examples/security/gen/grpc/secured_service/pb"
	securedservice "goa.design/examples/security/gen/secured_service"
)

// NewSigninPayload builds the payload of the "signin" endpoint of the
// "secured_service" service from the gRPC request type.
func NewSigninPayload(username string, password string) *securedservice.SigninPayload {
	v := &securedservice.SigninPayload{}
	v.Username = username
	v.Password = password
	return v
}

// NewSigninResponse builds the gRPC response type from the result of the
// "signin" endpoint of the "secured_service" service.
func NewSigninResponse(result *securedservice.Creds) *secured_servicepb.SigninResponse {
	message := &secured_servicepb.SigninResponse{
		Jwt:        result.JWT,
		ApiKey:     result.APIKey,
		OauthToken: result.OauthToken,
	}
	return message
}

// NewSecurePayload builds the payload of the "secure" endpoint of the
// "secured_service" service from the gRPC request type.
func NewSecurePayload(message *secured_servicepb.SecureRequest, token string) *securedservice.SecurePayload {
	v := &securedservice.SecurePayload{
		Fail: &message.Fail,
	}
	v.Token = token
	return v
}

// NewSecureResponse builds the gRPC response type from the result of the
// "secure" endpoint of the "secured_service" service.
func NewSecureResponse(result string) *secured_servicepb.SecureResponse {
	message := &secured_servicepb.SecureResponse{}
	message.Field = result
	return message
}

// NewDoublySecurePayload builds the payload of the "doubly_secure" endpoint of
// the "secured_service" service from the gRPC request type.
func NewDoublySecurePayload(message *secured_servicepb.DoublySecureRequest, token string) *securedservice.DoublySecurePayload {
	v := &securedservice.DoublySecurePayload{
		Key: message.Key,
	}
	v.Token = token
	return v
}

// NewDoublySecureResponse builds the gRPC response type from the result of the
// "doubly_secure" endpoint of the "secured_service" service.
func NewDoublySecureResponse(result string) *secured_servicepb.DoublySecureResponse {
	message := &secured_servicepb.DoublySecureResponse{}
	message.Field = result
	return message
}

// NewAlsoDoublySecurePayload builds the payload of the "also_doubly_secure"
// endpoint of the "secured_service" service from the gRPC request type.
func NewAlsoDoublySecurePayload(message *secured_servicepb.AlsoDoublySecureRequest, oauthToken *string, token *string) *securedservice.AlsoDoublySecurePayload {
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
	return v
}

// NewAlsoDoublySecureResponse builds the gRPC response type from the result of
// the "also_doubly_secure" endpoint of the "secured_service" service.
func NewAlsoDoublySecureResponse(result string) *secured_servicepb.AlsoDoublySecureResponse {
	message := &secured_servicepb.AlsoDoublySecureResponse{}
	message.Field = result
	return message
}

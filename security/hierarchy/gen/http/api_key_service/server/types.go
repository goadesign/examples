// Code generated by goa v3.18.2, DO NOT EDIT.
//
// api_key_service HTTP server types
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package server

import (
	apikeyservice "goa.design/examples/security/hierarchy/gen/api_key_service"
)

// NewDefaultPayload builds a api_key_service service default endpoint payload.
func NewDefaultPayload(key string) *apikeyservice.DefaultPayload {
	v := &apikeyservice.DefaultPayload{}
	v.Key = key

	return v
}

// NewSecurePayload builds a api_key_service service secure endpoint payload.
func NewSecurePayload(token string) *apikeyservice.SecurePayload {
	v := &apikeyservice.SecurePayload{}
	v.Token = token

	return v
}

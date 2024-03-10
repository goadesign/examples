// Code generated by goa v3.15.1, DO NOT EDIT.
//
// HTTP request path constructors for the api_key_service service.
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package server

// DefaultAPIKeyServicePath returns the URL path to the api_key_service service default HTTP endpoint.
func DefaultAPIKeyServicePath() string {
	return "/svc/default"
}

// SecureAPIKeyServicePath returns the URL path to the api_key_service service secure HTTP endpoint.
func SecureAPIKeyServicePath() string {
	return "/svc/secure"
}

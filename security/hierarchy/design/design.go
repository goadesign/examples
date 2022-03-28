package design

import (
	. "goa.design/goa/v3/dsl"
)

// BasicAuth defines a security scheme using basic authentication.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication")
})

// APIKeyAuth defines a security scheme that uses API keys.
var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token.`)
})

var _ = API("hierarchy", func() {
	Title("Security Example API")
	Description("This API demonstrates the effect of using Security at the API, Service or Method levels")

	Security(BasicAuth)
})

var _ = Service("default_service", func() {
	Method("default", func() {
		Description("The default service default_method is secured with basic authentication")
		Payload(func() {
			Username("username")
			Password("password")
			Required("username", "password")
		})
		HTTP(func() { GET("/default") })
	})
})

var _ = Service("api_key_service", func() {
	Description("The svc service is secured with API key based authentication")
	HTTP(func() { Path("/svc") })

	Security(APIKeyAuth)

	Method("default", func() {
		Payload(func() {
			APIKey("api_key", "key", String, func() {
				Description("API key used for authentication")
			})
			Required("key")
		})

		HTTP(func() { GET("/default") })
	})

	Method("secure", func() {
		Security(JWTAuth)
		Description("This method requires a valid JWT token.")

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
			})
			Required("token")
		})

		HTTP(func() { GET("/secure") })
	})

	Method("unsecure", func() {
		Description("This method is not secured.")
		NoSecurity()
	})
})

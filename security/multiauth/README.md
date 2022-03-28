# MultiAuth Example

The MultiAuth example defines a service that exposes one endpoint for each
security scheme natively supported by Goa (i.e. basic auth, API key, JWT auth
and OAuth2). The example also includes a couple of endpoints that show how to
apply multiple security schemes to a single endpoint.

## Design

The key design sections for the `multi_auth` service define the various security
requirements. The `JWTAuth`, `BasicAuth`, `APIKeyAuth` and `OAuth2Auth` variables
define the security schemes used by the service endpoints. The `doubly_secure`
endpoint requires both a JWT and an API key to authorize requests:

```go
Security(JWTAuth, APIKeyAuth, func() { // Use JWT and an API key to secure this endpoint.
	Scope("api:read")  // Enforce presence of both "api:read"
	Scope("api:write") // and "api:write" scopes in JWT claims.
})
```

The `Payload` DSL defines two attributes `key` and `token` that hold the API key
and JWT token respectively:

```go
Payload(func() {
	APIKey("api_key", "key", String, func() {
		Description("API key")
	})
	Token("token", String, func() {
		Description("JWT used for authentication")
	})
})
```

The `also_double_secure` endpoint requires either a JWT and API key pair or a
basic auth header and a OAuth2 token pair:

```go
Security(JWTAuth, APIKeyAuth, func() { // Use JWT and an API key to secure this endpoint.
    Scope("api:read")  // Enforce presence of both "api:read"
    Scope("api:write") // and "api:write" scopes in JWT claims.
})

Security(OAuth2Auth, BasicAuth, func() { // Or basic auth and OAuth2 token.
    Scope("api:read")  // Enforce presence of both "api:read"
    Scope("api:write") // and "api:write" scopes in OAuth2 claims.
})
```

The `HTTP` DSL defines where the values for each security element is loaded from:

```go
HTTP(func() {
    POST("/secure")
    Header("token:X-Authorization") // JWT token passed in "X-Authorization" header
    Param("key:k")                  // API key "key" sent in query parameter "k"
    Param("oauth_token:oauth")      // OAuth token sent in query parameter "oauth"
    Response(StatusOK)
    Response("invalid-scopes", StatusForbidden)
})
```

The `Payload` DSL defines the mapping to the payload struct that is given as
parameter to the endpoint method:

```go
Payload(func() {
    // UsernameField identifies the basic auth user.
    UsernameField(1, "username", String, "Username used to perform signin", func() {
        Example("user")
    })
    // PasswordField identifies the basic auth password.
    PasswordField(2, "password", String, "Password used to perform signin", func() {
        Example("password")
    })
    // APIKeyField identifies the API key.
    APIKeyField(3, "api_key", "key", String, func() {
        Description("API key")
        Example("abcdef12345")
    })
    // TokenField identifies the JWT.
    TokenField(4, "token", String, func() {
        Description("JWT used for authentication")
    })
    // AccessTokenField identifies the OAuth2 access token.
    AccessTokenField(5, "oauth_token", String)
})
```

# Security Hierarchy

## Security Schemes

Security schemes are global constucts in a Goa design that can be used to secure
entire APIs, services or individual methods. The schemes are defined using one
of the security DSLs:

* [APIKeySecurity](https://pkg.go.dev/goa.design/goa/v3/dsl#APIKeySecurity) defines a
  security scheme that uses an API key.
* [BasicAuthSecurity](https://pkg.go.dev/goa.design/goa/v3/dsl#BasicAuthSecurity)
  defines a security scheme that uses basic auth.
* [JWTSecurity](https://pkg.go.dev/goa.design/goa/v3/dsl#JWTSecurity) defines a security
  scheme that uses JWT tokens.
* [OAuth2Security](https://pkg.go.dev/goa.design/goa/v3/dsl#OAuth2Security) defines a
  security scheme that uses OAuth2.

Once a security scheme is defined using one of the functions above it can be
applied to all the endpoints defined in a design using the
[Security](https://pkg.go.dev/goa.design/goa/v3/dsl#Security) function in an
[API](https://pkg.go.dev/goa.design/goa/v3/dsl#API) definition: 

```go
var BasicAuth = BasicAuthSecurity("realm", func() {
    Description("Basic auth")
})

var _ = API("secure", func() {
    Security(BasicAuth)
    // ...
})
```

The example above defines a security scheme named `BasicAuth` that uses HTTP
authentication to secure all the endpoints defined in the design package by
default.

A service can also use the [Security](https://pkg.go.dev/goa.design/goa/v3/dsl#Security)
function to override (if the API already defined a default security scheme) or
define the security scheme applied to the service endpoints. 

```go
var BasicAuth = BasicAuthSecurity("realm", func() {
    Description("Basic auth")
})

var APIKeyAuth = APIKeySecurity("key", func() {
    Description("API key based authentication")
})

var _ = API("secure", func() {
    Security(BasicAuth)
    // ...
})

var _ = Service("svc", func() {
    Security(APIKeyAuth)
    // ...
})
```

In the example above the `svc` service endpoints are secured with the `APIKeyAuth`
security scheme while any other endpoint in the `secure` API is secured with the
`BasicAuth` security scheme.

Similarly a method may specify a different security scheme than the one defined
for the service or API if any using the same
[Security](https://pkg.go.dev/goa.design/goa/v3/dsl#Security) function:

```go
var BasicAuth = BasicAuthSecurity("realm", func() {
    Description("Basic auth")
})

var APIKeyAuth = APIKeySecurity("key", func() {
    Description("API key based authentication")
})

var JWTAuth = JWTSecurity("jwt", func() {
    Description("JWT based authentication")
})

var _ = API("secure", func() {
    Security(BasicAuth)
    // ...
})

var _ = Service("svc", func() {
    Security(APIKeyAuth)
    Method("method", func() {
        Security(JWTAuth)
        // ...
    })
})
```

In the example above the `method` endpoint is secured with the `JWTAuth`
security scheme while other `svc` endpoints are secured with the `APIKeyAuth`
security scheme.

## NoSecurity

The [NoSecurity](https://pkg.go.dev/goa.design/goa/v3/dsl#NoSecurity) function
can be used to explicitly remove any security scheme applied to a specific
endpoint:

```go
var _ = Service("svc", func() {
    Security(BasicAuth)
    Method("default", func() {
        // Secured with BasicAuth
    })
    Method("method", func() {
        Security(JWTAuth)
        // Secured with JWTAuth
    })
    Method("other", func() {
        NoSecurity()
        // Not secured
    })
})
```

## Generated Code

The generated code for the default service 
[endpoints](https://github.com/goadesign/examples/tree/master/security/hierarchy/gen/api_key_service/endpoints.go)
shows the basic auth security scheme defined at the API level being applied.

The [generated endpoints code](https://github.com/goadesign/examples/tree/master/security/hierarchy/gen/api_key_service/endpoints.go)
code the `api_key_service` shows the API key scheme applied the default
endpoint (using the security scheme defined at the service leve), the JWT
scheme used by the `secure` method that overrides the default set at the
service level and the `unsecure` method with no security scheme applied.

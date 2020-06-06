# Cookies Example

This example illustrates how to read and write HTTP cookies from a Goa service.

The `session` service exposes two methods: `create_session` and `use_session`.
`create_session` writes a cookie to the HTTP response and `use_session` reads
it back.

The available DSL functions used to manipulate HTTP cookies are:

* [Cookie](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#Cookie)
* [CookieMaxAge](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#CookieMaxAge)
* [CookiePath](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#CookiePath)
* [CookieDomain](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#CookieDomain)
* [CookieSecure](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#CookieSecure)
* [CookieHTTPOnly](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#CookieHTTPOnly)

The following shows the DSL used to write the cookie:

```go
HTTP(func() {
	POST("/")
	Response(StatusOK, func() {
		Cookie("session_id:SID") // Return session ID in "SID" cookie
		CookieMaxAge(3600)       // Sessions last one hour
	})
})
```

The value of the HTTP cookie `SID` is read from the `session_id` field of the
method result object.

The extract below shows the DSL used to read the cookie:

```go
HTTP(func() {
	GET("/")
	Cookie("session_id:SID") // Load session ID from "SID" cookie
	Response(StatusOK)
})
```

## Building and Running the example

Build the server and the client:

```
go build ./cmd/session && go build ./cmd/session-cli
```

Run the server:

```
./session --http-port 8080
```

Run the client to create a session:

```
./session-cli -v -url http://localhost:8080 session create-session --body '{"name":"foo"}' 
```

Run the client to use a session:

```
./session-cli -url http://localhost:8080 session use-session -session-id abcd
```

# Security Examples

This repo contains a set of examples that show how to design and implement secure APIs using goa.

There are examples covering [basic authentication](https://tools.ietf.org/html/rfc2617#section-2)
, the use of API keys, [JWT authentication](https://jwt.io) and
[OAuth2](https://tools.ietf.org/html/rfc6749).

Each example contains a `secured` and an `unsecured` endpoint. The `secured` endpoint requires the
client to authenticate using the scheme defined by the example.

Generate, compile and start the API with:

```
go generate
go build -o secured
./secured
```

The rest of this document uses the generated client to illustrate the various endpoints. Compile it
with:

```
cd tool/secured-cli
go build
```

## Basic Auth

The file [design/basic_auth.go](design/basic_auth.go) defines a basic authorization security scheme.
Clients of the API authenticate using basic auth as described in [section 2 of RFC 2617](https://tools.ietf.org/html/rfc2617#section-2).

The controller authorizes all requests with a non empty password. Real applications may want to
implement a better algorithm :)

Sending requests to the secured endpoint using an unauthenticated request:

```
./secured-cli 




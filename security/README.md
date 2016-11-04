# Security Examples

This repo contains a set of examples that show how to design and implement secure APIs using goa.

There are examples covering [basic authentication](https://tools.ietf.org/html/rfc2617#section-2)
, the use of API keys, [JWT authentication](https://jwt.io) and
[OAuth2](https://tools.ietf.org/html/rfc6749).

Each example contains a `secure` and an `unsecure` endpoint. The `secure` endpoint requires the
client to authenticate using the scheme defined by the example.

Generate, compile and start the API with:

```
go generate
go build -o secure
./secure
```

The rest of this document uses the generated client to illustrate the various endpoints. Compile it
with:

```
cd tool/secure-cli
go build
```

This illustrates how the security schemes defined in the design affect both the generated
controllers and client. With minimal code you get a secure API and a client package and tool that
can properly sign requests for the specific scheme.

Use the `--dump` flag with the client to see the details for each request and response.

## Basic Auth

The file [design/basic_auth.go](design/basic_auth.go) defines a basic authorization security scheme.
Clients of the API authenticate using basic auth as described in [section 2 of RFC 2617](https://tools.ietf.org/html/rfc2617#section-2).

The controller authorizes all requests with a non empty password. Real applications may want to
implement a better algorithm :)

Sending requests to the secure endpoint using an unauthenticated request:

```
❯ ./secure-cli secure basic
2016/06/14 01:14:31 [INFO] started id=VywOGRJD GET=http://localhost:8080/basic
2016/06/14 01:14:31 [INFO] completed id=VywOGRJD status=401 time=1.308411ms
error: 401: {"code":"unauthorized","status":401,"detail":"missing auth"}
```

Sending requests to the secure endpoint using an authenticated request:

```
❯ ./secure-cli --user foo --pass bar secure basic
2016/06/14 01:14:44 [INFO] started id=8sdzObwy GET=http://localhost:8080/basic
2016/06/14 01:14:44 [INFO] completed id=8sdzObwy status=200 time=1.824404ms
{"ok":true}
```

Sending requests to the unsecure endpoint using an unauthenticated request:

```
❯ ./secure-cli unsecure basic
2016/06/14 01:24:37 [INFO] started id=uUL8u8Vo GET=http://localhost:8080/basic/unsecure
2016/06/14 01:24:37 [INFO] completed id=uUL8u8Vo status=200 time=902.653µs
{"ok":true}
```

## API Key

The file [design/api_key.go](design/api_key.go) defines an API key security scheme using the
`X-Shared-Secret` header to pass in a static API key value.

The controller authorizes all requests with a non empty `X-Shared-Secret` header. As with the
basic auth examples real applications may want to do something more interesting...

Sending requests to the secure endpoint using an unauthenticated request:

```
❯ ./secure-cli secure api_key
2016/06/14 01:29:43 [INFO] started id=Mk3HKcfs GET=http://localhost:8080/api_key
2016/06/14 01:29:43 [INFO] completed id=Mk3HKcfs status=401 time=1.159544ms
error: 401: {"code":"unauthorized","status":401,"detail":"missing auth"}
```

Sending requests to the secure endpoint using an authenticated request:

```
❯ ./secure-cli --key foo secure api_key
2016/06/14 01:37:49 [INFO] started id=FCsDiXxb GET=http://localhost:8080/api_key
2016/06/14 01:37:49 [INFO] completed id=FCsDiXxb status=200 time=945.243µs
{"ok":true}
```

Sending requests to the unsecure endpoint using an unauthenticated request:

```
❯ ./secure-cli unsecure api_key 
2016/06/14 01:35:41 [INFO] started id=CbCgqGmi GET=http://localhost:8080/api_key/unsecure
2016/06/14 01:35:41 [INFO] completed id=CbCgqGmi status=200 time=859.733µs
{"ok":true}
```

## JWT

The file [design/jwt.go](design/jwt.go) defines a JWT security scheme using the key in the `jwtkey`
directory for validating the incoming requests.

The controller validates the incoming requests `Authorization` header using the private key to
decode its content. The `signin` action uses basic auth to authenticate incoming requests and
creates valid JWT tokens that can be used to send requests to the secure endpoint.

Sending requests to the secure endpoint using an unauthenticated request:

```
❯ ./secure-cli secure jwt
2016/06/14 01:43:02 [INFO] started id=iTBZsjo6 GET=http://localhost:8080/jwt?fail=false
2016/06/14 01:43:02 [INFO] completed id=iTBZsjo6 status=401 time=1.389178ms
error: 401: {"code":"jwt_security_error","status":401,"detail":"invalid or malformed \"Bearer\" header, expected 'Authorization: Bearer JWT-token...'"}
```

Retrieving a valid JWT token:

```
❯ ./secure-cli --user foo --pass bar signin jwt --dump
2016/06/14 02:06:18 [INFO] started id=uRpfVTBx POST=http://localhost:8080/jwt/signin
2016/06/14 02:06:18 [INFO] request headers Authorization=Basic Zm9vOmJhcg== User-Agent=Secure-cli/0
2016/06/14 02:06:18 [INFO] completed id=uRpfVTBx status=204 time=5.419953ms
2016/06/14 02:06:18 [INFO] response headers Authorization=Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJBdWRpZW5jZSIsImV4cCI6MTQ2NTg5NTc3OCwiaWF0IjoiMjAxNi0wNi0xNFQwMjowNjoxOC45NzM0MzgzMDItMDc6MDAiLCJpc3MiOiJJc3N1ZXIiLCJqdGkiOiJmNGI2ZTVhZC02OTdiLTRlYjQtYTczNi04NzFmYzM3MmQzNDUiLCJuYmYiOjIsInNjb3BlcyI6ImFwaTphY2Nlc3MiLCJzdWIiOiJzdWJqZWN0In0.SbzvKFSzzV-1Dt24C6Cpon7X7z5l8_jvkuY3FHtuQWtO5JSlrGBbe1iYxxA2Wb71qQ7FqSrW6s8l2ir4377S-V9ZyuUPEfJPPUmxmDeKy64a1VB4Gu9VxuwfcBa8DQLi7yKwnuYbzSKW4Y4ETz8myF5ywi7CKP4rEwHvl8VVoAXVsiVoukXcN8XCqNb9Slnf2yl8OdWTr8FiXm7dPDM5m5RcHIbXOZWpHJK_q1lcFJqlLlBKTDFmKUP2-Uh1x80xfVYVV3BV6fsXGOvODM9qgSUagL4CTZ8i_IU_puRcWp9y4MhikwqlkjHGC7uz3tsM2INY757rLGDpEDq6ey2xlg Date=Tue, 14 Jun 2016 09:06:18 GMT
2016/06/14 02:06:18 [INFO] response body=
```

Sending requests to the secure endpoint using an authenticated request:

```
❯ ./secure-cli --key eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJBdWRpZW5jZSIsImV4cCI6MTQ2NTg5NTc3OCwiaWF0IjoiMjAxNi0wNi0xNFQwMjowNjoxOC45NzM0MzgzMDItMDc6MDAiLCJpc3MiOiJJc3N1ZXIiLCJqdGkiOiJmNGI2ZTVhZC02OTdiLTRlYjQtYTczNi04NzFmYzM3MmQzNDUiLCJuYmYiOjIsInNjb3BlcyI6ImFwaTphY2Nlc3MiLCJzdWIiOiJzdWJqZWN0In0.SbzvKFSzzV-1Dt24C6Cpon7X7z5l8_jvkuY3FHtuQWtO5JSlrGBbe1iYxxA2Wb71qQ7FqSrW6s8l2ir4377S-V9ZyuUPEfJPPUmxmDeKy64a1VB4Gu9VxuwfcBa8DQLi7yKwnuYbzSKW4Y4ETz8myF5ywi7CKP4rEwHvl8VVoAXVsiVoukXcN8XCqNb9Slnf2yl8OdWTr8FiXm7dPDM5m5RcHIbXOZWpHJK_q1lcFJqlLlBKTDFmKUP2-Uh1x80xfVYVV3BV6fsXGOvODM9qgSUagL4CTZ8i_IU_puRcWp9y4MhikwqlkjHGC7uz3tsM2INY757rLGDpEDq6ey2xlg secure jwt
2016/06/14 02:06:44 [INFO] started id=f1gRUZe0 GET=http://localhost:8080/jwt?fail=false
2016/06/14 02:06:44 [INFO] completed id=f1gRUZe0 status=200 time=1.132171ms
{"ok":true}
```

Sending requests to the unsecure endpoint using an unauthenticated request:

```
❯ ./secure-cli unsecure jwt
2016/06/14 02:07:43 [INFO] started id=FWyDR/Ta GET=http://localhost:8080/jwt/unsecure
2016/06/14 02:07:43 [INFO] completed id=FWyDR/Ta status=200 time=878.099µs
{"ok":true}
```

## OAuth2

OAuth2 is not solely an authentication mechanism, it's also a way to have 3rd party services send
requests to the service on behalf of end users. In OAuth2 RFC speak OAuth2 makes it possible for
Clients (3rd party services) to send requests to the Resource Server (your service) after being
granted access by the Resource Owner (end user). The Authorization Server is in charge of
implementing the corresponding request handlers.

For a more detailed description of OAuth2 and the "Authorization Code" flow being implemented in
this example refer to the [github.com/goadesign/oauth2](https://github.com/goadesign/oauth2) package
README and documentation.

The goa oauth2 package takes care of implementing the Authorization Server. All you have to do is to
provide it with the business logic that validates authorization codes and refresh tokens and creates
access tokens. This is done by implementing the oauth2.Provider interface as done by the
`DummyProvider` included in this example.

Once the provider is implemented the OAuth2 security middleware validates the access tokens included
in requests targetting secure endpoints. There are thus a few steps involved in exercising the
OAuth2 authorization:

1. The client needs to retrieve an authorization code by making a request to the "authorization
   code" endpoint
2. The client should then exchange the authorization code for a pair of refresh an access tokens by
   making a request to the "token" endpoint.
3. Finally the client can make requests to secure endpoints by using the access token in the
   "Authorization" header.

Sending requests to retrieve authorization codes:

```
❯ ./secure-cli authorize oauth2_provider --response_type=code --client_id=foo --redirect_uri=http://localhost:8080/oauth2/handle_redirect
2016/06/14 18:00:47 [INFO] started id=WcC6KWC6 GET=http://localhost:8080/oauth2/authorize?client_id=foo&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Foauth2%2Fhandle_redirect&response_type=code&scope=&state=
2016/06/14 18:00:47 [INFO] completed id=WcC6KWC6 status=204 time=1.213178ms
```

Looking at the service logs we can see it received the authorization redirect with the code:

```
2016/06/14 18:00:47 [INFO] started req_id=1bd2n3cfyC-2 GET=/oauth2/handle_redirect?code=authcode from=::1 ctrl=OAuth2Controller action=HandleRedirect
2016/06/14 18:00:47 [INFO] params req_id=1bd2n3cfyC-2 code=authcode
2016/06/14 18:00:47 [INFO] got redirect request req_id=1bd2n3cfyC-2 code=authcode
2016/06/14 18:00:47 [INFO] completed req_id=1bd2n3cfyC-2 status=204 bytes=0 time=54.136µs
```

Using the code to retrieve refresh and access tokens:

```
❯ ./secure-cli --user client --pass secret getToken oauth2_provider --payload='{"grant_type":"authorization_code","code":"authcode","redirect_uri":"http://localhost:8080/handle_redirect"}' --content="application/x-www-form-urlencoded"
2016/06/14 22:36:23 [INFO] started id=iwE21eIG POST=http://localhost:8080/oauth2/token
2016/06/14 22:36:23 [INFO] completed id=iwE21eIG status=200 time=1.126004ms
{"access_token":"accesstoken","expires_in":3600,"refresh_token":"refreshtoken","token_type":"Bearer"}
```

Using the access token to send requests to the secure endpoint:

```
❯ ./secure-cli --token accesstoken secure oauth2
2016/06/14 23:05:52 [INFO] started id=rW3rVFc2 GET=http://localhost:8080/oauth2/read
2016/06/14 23:05:52 [INFO] completed id=rW3rVFc2 status=200 time=1.821886ms
{"ok":true}
```

Refreshing the access token:

```
❯ ./secure-cli --user client --pass secret getToken oauth2_provider --payload='{"grant_type":"refresh_token","refresh_token":"refreshtoken"}' --content="application/x-www-form-urlencoded"
2016/06/14 23:07:42 [INFO] started id=km2N7Udg POST=http://localhost:8080/oauth2/token
2016/06/14 23:07:42 [INFO] completed id=km2N7Udg status=200 time=1.001427ms
{"access_token":"accesstoken","expires_in":3600,"refresh_token":"refreshtoken","token_type":"Bearer"}
```

That's it! Now go forth and secure your APIs.

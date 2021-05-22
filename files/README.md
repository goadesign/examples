# File Systems

This example illustrates how to change the file system of a file server.

## Design

The example defines a `openapi` service and four file servers using `Files` DSL.

## Implementation

The HTTP server initializer receives file systems for each file server.

```go
		oapi3 := http.FS(openapiapi.OpenAPI3)
		openapiServer = openapisvr.New(nil, mux, dec, enc, eh, nil,
			// http.FileSystem system for openapi.json.
			// http.Dir(".") will be used if nil.
			nil,

			// http.FileSystem for openapi.yaml.
			// It's converted from embed.FS using http.FS().
			http.FS(openapiapi.OpenAPI),

			// http.FileSystem for openapi3.json and openapi3.yaml.
			// A file system can be used for multiple file servers.
			oapi3,
			oapi3,
		)
```

## Running the Example

First compile and start the server:

```bash
cd $GOPATH/src/goa.design/examples/files/cmd/openapi
go build; ./openapi --http-port 8080
```

This shoud produce output similar to:

```
[openapiapi] 17:32:39 HTTP "gen/http/openapi.json" mounted on GET /openapi.json
[openapiapi] 17:32:39 HTTP "gen/http/openapi.yaml" mounted on GET /openapi.yaml
[openapiapi] 17:32:39 HTTP "gen/http/openapi3.json" mounted on GET /openapi3.json
[openapiapi] 17:32:39 HTTP "gen/http/openapi3.yaml" mounted on GET /openapi3.yaml
[openapiapi] 17:32:39 HTTP server listening on "localhost:8080"
```

We use [httpie](https://httpie.io/) to make the requests:

```bash
$ http GET :8080/openapi.json
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 1120
Content-Type: application/json
Date: Sat, 22 May 2021 08:43:58 GMT
Last-Modified: Sat, 22 May 2021 07:39:58 GMT

{ ... }


$ http GET :8080/openapi.yaml
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 1333
Content-Type: text/plain; charset=utf-8
Date: Sat, 22 May 2021 08:45:25 GMT

...


$ http GET :8080/openapi3.json
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 927
Content-Type: application/json
Date: Sat, 22 May 2021 08:46:09 GMT

{ ... }


$ http GET :8080/openapi3.yaml
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 1031
Content-Type: text/plain; charset=utf-8
Date: Sat, 22 May 2021 08:47:23 GMT

...


```
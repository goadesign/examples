# Upload and Download

This example illustrates how to implement upload and download of binary files
using Goa. The example uses straight HTTP and shows how both server and client
code can stream content without having to load the entire payload in memory.

## Design

The `updown` service exposes two methods: `upload` which allows clients to
stream content to the server and `download` which does the opposite. The key DSL
functions that enable the streaming are
[SkipRequestBodyEncodeDecode](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#SkipRequestBodyEncodeDecode)
and
[SkipResponseBodyEncodeDecode](https://pkg.go.dev/goa.design/goa/v3/dsl?tab=doc#SkipResponseBodyEncodeDecode).
The DSL functions tell Goa to bypass the generation of an encoder and decoder for the HTTP request
and response bodies respectively.

Here is an highlight of the key parts of the `upload` method design:

```go
Method("upload", func() {
    // A payload may be defined for methods that use SkipRequestBodyEncodeDecode.
    // The attributes of the payload must all be mapped to either HTTP headers or parameters.
    Payload(func() {
        Attribute("content_type", String, "...")
        Attribute("dir", String, "...")
    })

    // The HTTP DSL makes use of SkipRequestBodyEncodeDecode instructing
    // Goa to bypass the code generation for the request encoder and decoder
    // and instead make the underlying request body io.Reader available to
    // the service. The generated code also lets the client provide a
    // io.Reader that gets streamed to the server.
    HTTP(func() {
        POST("/upload/{*dir}")
        Header("content_type:Content-Type")
        SkipRequestBodyEncodeDecode()
    })
})
```

and of the `download` method design:

```go
Method("download", func() {
    Payload(String) // Name of downloaded file
    
    // A result may be defined for methods that use SkipResponseBodyEncodeDecode.
    // The attributes of the result must all be mapped to HTTP headers.
	Result(func() {
		Attribute("length", Int64, "...")
		Required("length")
	})

    // The HTTP DSL makes use of SkipResponseBodyEncodeDecode instructing
    // Goa to bypass the code generation for the response encode and decoder
    // and instead make the underlying response body io.Reader available to
    // the client. The generated code also lets the service provide a
    // io.Reader that gest streamed to the client.
	HTTP(func() {
        GET("/download/{*filename}")
		SkipResponseBodyEncodeDecode()
		Response(func() {
			Header("length:Content-Length")
		})
	})
})
```

## Building and Running the Generated Server and Client

The generated example server and client can be built and run as follows:

```
$ go build ./cmd/upload_download && go build ./cmd/upload_download-cli 

# Run the server

$ ./upload_download

# Run the client

# The generated client tool defines a --stream flag that makes it possible to upload a file.
$ ./upload_download-cli -url http://localhost:8080  updown upload --stream public/goa.png --dir upload
```

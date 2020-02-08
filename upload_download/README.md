# Upload and Download

This example illustrates how to implement upload and download of binary files
using Goa. The example uses straight HTTP and shows how both server and client
code can stream content without having to load the entire payload in memory.

## Design

The `upload` service exposes two methods: `upload` which allows clients to
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
	Payload(func() {
		Attribute("length", UInt, "...")
		Attribute("name", String, "...")
		Required("length", "name")
	})
	Result(String)
	HTTP(func() {
		POST("/{*name}")
		Header("length:Content-Length")
		SkipRequestBodyEncodeDecode()
	})
})
```

and of the `download` method design:

```go
Method("download", func() {
	Payload(String) // Name of downloaded file
	Result(func() {
		Attribute("length", UInt, "...")
		Required("length")
	})

	HTTP(func() {
		GET("/{*name}")
		Response(func() {
			Header("length:Content-Length")
			SkipResponseBodyEncodeDecode()
		})
	})
})
```

## Building and Running the Generated Server and Client

The generated example server and client can be built and run as follows

```
$ go build ./cmd/upload && go build ./cmd/upload-cli

# Run the server

$ ./upload

# Run the client

# The generated client tool --file argument makes it possible to upload a file.
$ ./upload-cli --url="http://localhost:8000" updown upload --file goa.png

$ ./upload-cli --url="http://localhost:8000" updown download --id goa.png

```

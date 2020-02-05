package design

import . "goa.design/goa/v3/dsl"

var _ = API("upload_download", func() {
	Description("Simple file upload and download example.")
})

var _ = Service("updown", func() {
	Description(`Service updown exposes both a upload and download methods.

This example demonstrates how to implement upload and download of big files in
Goa without having to load the entire content in memory first.

The upload method uses SkipRequestBodyEncodeDecode to delegate reading the HTTP
request body to the service logic. This alleviates the need for loading the
full body content in memory first to decode it into a data structure. Note that
using SkipRequestBodyDecode is incompatible with gRPC and can only be used on
methods that only define a HTTP transport mapping. This example implementation
leverages package "mime/multipart" to read the request body.

Similarly the download method uses SkipResponseBodyEncodeDecode to stream the 
file to the client without requiring to load the complete content in memory
first. As with SkipRequestBodyDecode using SkipResponseBodyEncodeDecode is
incompatible with gRPC.`)

	Method("upload", func() {

		// The use of Payload here illustrates how HTTP headers and parameters can still be
		// properly decoded and validated when using SkipRequestBodyDecode. It is not
		// generally required to implement an upload method.
		Payload(func() {
			Attribute("length", UInt, "Length is the upload content length in bytes.", func() {
				Example(4 * 1024 * 1024)
			})
			Attribute("name", String, "Name is the name of the file being uploaded", func() {
				Example("goa.png")
			})
			Required("length", "name")
		})

		Result(String)

		HTTP(func() {
			POST("/{*name}")
			Header("length:Content-Length")

			// Bypass request body decoder code generation to alleviate need for loading
			// the entire request body in memory. The service gets direct access to the
			// HTTP request body reader.
			SkipRequestBodyEncodeDecode()
		})
	})

	Method("download", func() {
		Payload(String) // Name of downloaded file

		// The use of Result here illustrates how HTTP headers can still be properly encoded
		// and validated when using SkipResponseBodyEncode. It is not generally required to
		// implement a download method.
		Result(func() {
			Attribute("length", UInt, "Length is the downloaded content length in bytes.", func() {
				Example(4 * 1024 * 1024)
			})
			Required("length")
		})

		HTTP(func() {
			GET("/{*name}") // Encode payload in URL path

			// Bypass response body encoder code generation to alleviate need for
			// loading the entire response body in memory.
			SkipResponseBodyEncodeDecode()

			Response(func() {
				Header("length:Content-Length")
			})
		})
	})
})

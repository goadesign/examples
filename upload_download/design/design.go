package design

import . "goa.design/goa/v3/dsl"

var _ = API("upload_download", func() {
	Description("Simple file upload and download example.")
})

var _ = Service("updown", func() {
	Description(`Service updown demonstrates how to implement upload and download of files in
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

		// The payload defines the request headers and parameters. It cannot
		// define body attributes as the endpoint makes use of
		// SkipRequestBodyEncodeDecode.
		Payload(func() {
			Attribute("content_type", String, "Content-Type header, must define value for multipart boundary.", func() {
				Default("multipart/form-data; boundary=goa")
				Pattern("multipart/[^;]+; boundary=.+")
				Example("multipart/form-data; boundary=goa")
			})
			Attribute("dir", String, "Dir is the relative path to the file directory where the uploaded content is saved.", func() {
				Default("upload")
				Example("upload")
			})
		})

		Error("invalid_media_type", ErrorResult, "Error returned when the Content-Type header does not define a multipart request.")
		Error("invalid_multipart_request", ErrorResult, "Error returned when the request body is not a valid multipart content.")
		Error("internal_error", ErrorResult, "Fault while processing upload.")

		HTTP(func() {
			POST("/upload/{*dir}")
			Header("content_type:Content-Type")

			// Bypass request body decoder code generation to alleviate need for
			// loading the entire request body in memory. The service gets
			// direct access to the HTTP request body reader.
			SkipRequestBodyEncodeDecode()

			// Define error HTTP statuses.
			Response("invalid_media_type", StatusBadRequest)
			Response("invalid_multipart_request", StatusBadRequest)
			Response("internal_error", StatusInternalServerError)
		})
	})

	Method("download", func() {
		Payload(String, func() {
			Description("Path to downloaded file.")
		})

		// The use of Result here illustrates how HTTP headers can still be
		// properly encoded and validated when using SkipResponseBodyEncode. It
		// is not generally required to implement a download method.
		Result(func() {
			Attribute("length", Int64, "Length is the downloaded content length in bytes.", func() {
				Example(4 * 1024 * 1024)
			})
			Required("length")
		})

		Error("invalid_file_path", ErrorResult, "Could not locate file for download")
		Error("internal_error", ErrorResult, "Fault while processing download.")

		HTTP(func() {
			GET("/download/{*filename}") // Encode payload in URL path

			// Bypass response body encoder code generation to alleviate need for
			// loading the entire response body in memory.
			SkipResponseBodyEncodeDecode()

			Response(func() {
				Header("length:Content-Length")
			})

			// Define error HTTP statuses
			Response("invalid_file_path", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
		})
	})
})

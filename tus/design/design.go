package design

import (
	. "goa.design/goa/v3/dsl"
)

// ErrorResponse defines a response with an empty body.
func ErrorResponse(name string, code int, fn ...func()) {
	Response(name, code, func() {
		Body(Empty)
		if len(fn) > 0 {
			fn[0]()
		}
	})
}

var _ = API("tus upload", func() {
	Title("File Upload Service")
	Description("HTTP service for uploading files using the tus protocol (https://tus.io)")

	Server("Upload", func() {
		Description("Upload hosts the upload service.")

		Services("tus")

		Host("development", func() {
			Description("Development host.")
			URI("http://localhost:8000/upload")
		})
	})
})

var _ = Service("tus", func() {
	Description("The tus service exposes the methods required to implement the tus protocol")

	Error("InvalidTusResumable", ErrInvalidTUSResumable, func() {
		Description("If the version specified by the Client is not supported by the Server, it MUST respond with the 412 Precondition Failed status.")
	})

	HTTP(func() {
		// Base path for all endpoints.
		Path("/upload")
	})

	Method("head", func() {
		Description("Clients use the HEAD request to determine the offset at which the upload should be continued.")
		Payload(func() {
			Reference(TUSCoreHeaders)
			ID()
			Attribute("tusResumable")
			Required("id", "tusResumable")
		})

		Result(func() {
			Reference(TUSCoreResponseHeaders)
			Reference(TUSExtensionHeaders)
			Attribute("tusResumable")
			Attribute("uploadOffset")
			Attribute("uploadLength")
			Attribute("uploadDeferLength")
			Attribute("uploadMetadata")
			Required("tusResumable", "uploadOffset")
		})

		Error("NotFound", func() {
			Description("If the resource is not found, the Server SHOULD return either the 404 Not Found, 410 Gone or 403 Forbidden status without the Upload-Offset header.")
		})
		Error("Gone", func() {
			Description("If the resource is not found, the Server SHOULD return either the 404 Not Found, 410 Gone or 403 Forbidden status without the Upload-Offset header.")
		})

		HTTP(func() {
			HEAD("/{id}")
			Header("tusResumable:Tus-Resumable")
			Response(StatusOK, func() {
				Header("tusResumable:Tus-Resumable")
				Header("uploadOffset:Upload-Offset")
				Header("uploadLength:Upload-Length")
				Header("uploadDeferLength:Upload-Defer-Length")
				Header("uploadMetadata:Upload-Metadata")
			})
			ErrorResponse("InvalidTusResumable", StatusPreconditionFailed, func() { Header("tusVersion:Tus-Version") })
			ErrorResponse("NotFound", StatusNotFound)
			ErrorResponse("Gone", StatusGone)
		})
	})

	Method("patch", func() {
		Description(`Clients use the PATCH method to start or resume an upload.`)
		Payload(func() {
			Reference(TUSCoreHeaders)
			Reference(TUSExtensionHeaders)
			ID()
			Attribute("tusResumable")
			Attribute("uploadOffset")
			Attribute("uploadChecksum")
			Required("id", "tusResumable", "uploadOffset")
		})

		Result(func() {
			Reference(TUSCoreResponseHeaders)
			Reference(TUSExtensionHeaders)
			Attribute("tusResumable")
			Attribute("uploadOffset")
			Attribute("uploadExpires")
			Required("tusResumable", "uploadOffset")
		})

		Error("InvalidContentType", func() {
			Description("All PATCH requests MUST use Content-Type: application/offset+octet-stream, otherwise the server SHOULD return a 415 Unsupported Media Type status.")
		})
		Error("InvalidOffset", func() {
			Description("If the offsets do not match, the Server MUST respond with the 409 Conflict status without modifying the upload resource.")
		})
		Error("NotFound", func() {
			Description("If a Client does attempt to resume an upload which has since been removed by the Server, the Server SHOULD respond with the404 Not Found or 410 Gone status.")
		})
		Error("Gone", func() {
			Description("If a Client does attempt to resume an upload which has since been removed by the Server, the Server SHOULD respond with the404 Not Found or 410 Gone status.")
		})
		Error("InvalidChecksumAlgorithm", func() {
			Description("The checksum algorithm is not supported by the server.")
		})
		Error("ChecksumMismatch", func() {
			Description("The checksums mismatch.")
		})
		Error("Internal", func() {
			Description("Internal error")
		})

		HTTP(func() {
			PATCH("/{id}")
			Header("tusResumable:Tus-Resumable")
			Header("uploadOffset:Upload-Offset")
			Header("uploadChecksum:Upload-Checksum")
			SkipRequestBodyEncodeDecode()
			Response(StatusNoContent, func() {
				Header("tusResumable:Tus-Resumable")
				Header("uploadOffset:Upload-Offset")
				Header("uploadExpires:Upload-Expires")
			})
			ErrorResponse("InvalidTusResumable", StatusPreconditionFailed, func() { Header("tusVersion:Tus-Version") })
			ErrorResponse("InvalidContentType", StatusUnsupportedMediaType)
			ErrorResponse("InvalidOffset", StatusConflict)
			ErrorResponse("NotFound", StatusNotFound)
			ErrorResponse("Gone", StatusGone)
			ErrorResponse("InvalidChecksumAlgorithm", StatusBadRequest)
			ErrorResponse("ChecksumMismatch", 460 /*StatusChecksumMismatch*/)
			ErrorResponse("Internal", StatusInternalServerError)
		})
	})

	Method("options", func() {
		Description("Clients use the OPTIONS method to gather information about the Server’s current configuration.")
		Payload(Empty)

		Result(func() {
			Reference(TUSCoreResponseHeaders)
			Reference(TUSExtensionHeaders)
			Attribute("tusResumable")
			Attribute("tusVersion")
			Attribute("tusExtension")
			Attribute("tusMaxSize")
			Attribute("tusChecksumAlgorithm")
			Required("tusResumable", "tusVersion", "tusExtension", "tusChecksumAlgorithm")
		})

		HTTP(func() {
			OPTIONS("/")
			Response(StatusNoContent, func() {
				Header("tusResumable:Tus-Resumable")
				Header("tusVersion:Tus-Version")
				Header("tusExtension:Tus-Extension")
				Header("tusMaxSize:Tus-Max-Size")
				Header("tusChecksumAlgorithm:Tus-Checksum-Algorithm")
			})
			ErrorResponse("InvalidTusResumable", StatusPreconditionFailed, func() { Header("tusVersion:Tus-Version") })
		})
	})

	// creation and creation-wiht-upload extensions support.
	Method("post", func() {
		Description("Clients use the POST method against a known upload creation URL to request a new upload resource.")
		Payload(func() {
			Reference(TUSCoreHeaders)
			Reference(TUSExtensionHeaders)
			Attribute("tusResumable")
			Attribute("uploadLength")
			Attribute("uploadDeferLength")
			Attribute("uploadChecksum")
			Attribute("uploadMetadata")
			Attribute("tusMaxSize", Int64, "Length of the upload")
		})
		Result(func() {
			Reference(TUSCoreResponseHeaders)
			Reference(TUSExtensionHeaders)
			Attribute("tusResumable")
			Attribute("uploadOffset")
			Attribute("uploadExpires")
			Attribute("location", String, "URL of created resource", func() {
				Format(FormatURI)
				Example("/files/123")
			})
			Required("tusResumable", "uploadOffset", "location")
		})

		Error("MissingHeader", func() {
			Description("The request MUST include one of the following headers: a) Upload-Length -or- b) Upload-Defer-Length: 1")
		})
		Error("InvalidDeferLength", func() {
			Description("If the Upload-Defer-Length header contains any other value than 1 the server should return a 400 Bad Request status.")
		})
		Error("MaximumSizeExceeded", func() {
			Description("If the length of the upload exceeds the maximum, which MAY be specified using the Tus-Max-Size header, the Server MUST respond with the 413 Request Entity Too Large status.")
		})
		Error("InvalidChecksumAlgorithm", func() {
			Description("The checksum algorithm is not supported by the server.")
		})
		Error("ChecksumMismatch", func() {
			Description("The checksums mismatch.")
		})

		HTTP(func() {
			POST("/")
			Header("tusResumable:Tus-Resumable")
			Header("uploadLength:Upload-Length")
			Header("uploadDeferLength:Upload-Defer-Length")
			Header("uploadChecksum:Upload-Checksum")
			Header("uploadMetadata:Upload-Metadata")
			Header("tusMaxSize:Tus-Max-Size")
			SkipRequestBodyEncodeDecode()
			Response(StatusCreated, func() {
				Header("location:Location")
				Header("tusResumable:Tus-Resumable")
				Header("uploadOffset:Upload-Offset")
				Header("uploadExpires:Upload-Expires")
			})
			ErrorResponse("InvalidTusResumable", StatusPreconditionFailed, func() { Header("tusVersion:Tus-Version") })
			ErrorResponse("MissingHeader", StatusBadRequest)
			ErrorResponse("InvalidDeferLength", StatusBadRequest)
			ErrorResponse("MaximumSizeExceeded", StatusRequestEntityTooLarge)
			ErrorResponse("InvalidChecksumAlgorithm", StatusBadRequest)
			ErrorResponse("ChecksumMismatch", 460 /*StatusChecksumMismatch*/)
		})
	})

	// termination extension support
	Method("delete", func() {
		Description("Clients use the DELETE method to terminate completed and unfinished uploads allowing the Server to free up used resources.")
		Payload(func() {
			Reference(TUSCoreHeaders)
			ID()
			Attribute("tusResumable")
			Required("id", "tusResumable")
		})

		Result(func() {
			Reference(TUSCoreResponseHeaders)
			Attribute("tusResumable")
			Required("tusResumable")
		})

		Error("NotFound", func() {
			Description("For all future requests to this URL, the Server SHOULD respond with the 404 Not Found or 410 Gone status.")
		})
		Error("Gone", func() {
			Description("For all future requests to this URL, the Server SHOULD respond with the 404 Not Found or 410 Gone status.")
		})

		HTTP(func() {
			DELETE("/{id}")
			Header("tusResumable:Tus-Resumable")
			Response(StatusNoContent, func() {
				Header("tusResumable:Tus-Resumable")
			})
			ErrorResponse("InvalidTusResumable", StatusPreconditionFailed, func() { Header("tusVersion:Tus-Version") })
			ErrorResponse("NotFound", StatusNotFound)
			ErrorResponse("Gone", StatusGone)
		})
	})
})

var ErrInvalidTUSResumable = Type("ErrInvalidTUSResumable", func() {
	Description("ErrInvalidTUSResumable describes the error returned when a non-supported Tus-Resumable header is provided by the client.")
	Attribute("tusVersion", String, "Comma separated list of supported versions.", func() {
		Enum("1.0.0")
	})
	Required("tusVersion")
})

var TUSCoreHeaders = Type("TUSCoreHeaders", func() {
	Description("TUSCoreHeaders defines the tus core protocol headers used in both requests and responses.")
	Attribute("uploadOffset", Int64, "uploadOffset represents a byte offset within a resource.")
	Attribute("uploadLength", Int64, "uploadLength represents the size of the entire upload in bytes.")
	Attribute("tusResumable", String, "tusResumable represents a tus protocol version.", func() {
		Pattern(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\+[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?$`)
		Example("1.0.0")
	})
	Required("tusResumable")
})

var TUSCoreRequestHeaders = Type("TUSCoreRequestHeaders", func() {
	Description("TUSCoreRequestHeaders defines the tus core protocol request headers.")
	Extend(TUSCoreHeaders)
})

var TUSCoreResponseHeaders = Type("TUSCoreResponseHeaders", func() {
	Description("TUSCoreResponseHeaders defines the tus core protocol response headers.")
	Extend(TUSCoreHeaders)
	Attribute("tusVersion", String, "tusVersion is a comma separated list of protocol versions supported by the server. This implementation only supports 1.0.0.", func() {
		Enum("1.0.0")
	})
	Attribute("tusExtension", String, "tusExtension is a comma separated list of extensions supported by the server. This implementation supports the creation, creation-with-upload, expiration, checksum and termination extensions", func() {
		Enum("creation,creation-with-upload,creation-defer-length,expiration,checksum,termination")
	})
	Attribute("tusMaxSize", Int64, "tusMaxSize represents the maximum allowed size of an entire upload in bytes.")
})

var TUSExtensionHeaders = Type("TUSExtensionHeaders", func() {
	Description("TUSExtensionHeaders defines the suppoterd tus protocol extension headers used in both requests and responses.")
	Attribute("uploadDeferLength", Int, "The Upload-Defer-Length request and response header indicates that the size of the upload is not known currently and will be transferred later.", func() {
		Enum(1)
	})
	Attribute("uploadMetadata", String, "The Client MAY supply the Upload-Metadata header to add additional metadata to the upload creation request.", func() {
		Example("key1 val1,key2 val2")
	})
	Attribute("uploadExpires", String, "The Upload-Expires response header indicates the time after which the unfinished upload expires.", func() {
		Example("Wed, 25 Jun 2014 16:00:00 GMT")
	})
	Attribute("uploadChecksum", String, "A Client MAY include the Upload-Checksum header in a PATCH request. Once the entire request has been received, the Server MUST verify the uploaded chunk against the provided checksum using the specified algorithm.", func() {
		Example("sha1 Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
	})
	Attribute("tusChecksumAlgorithm", String, "A Client MAY include the Upload-Checksum header in a PATCH request. Once the entire request has been received, the Server MUST verify the uploaded chunk against the provided checksum using the specified algorithm.", func() {
		Enum("md5,sha1,crc32")
	})
})

// ID is the attribute used to represent upload identifiers.
func ID() {
	Attribute("id", String, "id is the unique upload identifier.", func() {
		Description("IDs are generated using Xid: https://github.com/rs/xid")
		Pattern(`[0-9a-v]{20}`)
	})
}

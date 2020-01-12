package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("tus upload", func() {
	Title("File Upload Service")
	Description("HTTP service for uploading files using the tus protocol (https://tus.io)")

	Server("Upload", func() {
		Description("Upload hosts the upload service.")

		Services("upload")

		Host("development", func() {
			Description("Development host.")
			URI("http://localhost:8000/upload")
		})
	})
})

var _ = Service("upload", func() {
	Description("The upload service exposes the methods required to implement the tus protocol")

	HTTP(func() {
		// Base path for all HTTP methods.
		Path("/files")
	})

	Method("head", func() {
		Description("Clients use the HEAD request to determine the offset at which the upload should be continued.")
		Payload(func() {
			Reference(TUSCoreHeaders)
			ID()
			Attribute("tusResumable")
			Attribute("uploadOffset")
			Required("id", "tusResumable")
		})

		Result(func() {
			Reference(TUSCoreHeaders)
			Reference(TUSExtensionHeaders)
			Attribute("tusResumable")
			Attribute("uploadOffset")
			Attribute("uploadLength")
			Attribute("uploadDeferLength")
			Attribute("uploadMetadata")
			Required("tusResumable", "uploadOffset")
		})

		HTTP(func() {
			HEAD("/{id}")
			Header("tusResumable")
			Header("uploadOffset")
			Response(StatusOK, func() {
				Header("tusResumable")
				Header("uploadOffset")
				Header("uploadLength")
				Header("uploadDeferLength")
				Header("uploadMetadata")
			})
		})
	})

	Method("patch", func() {
		Description("Clients use the PATCH method to start or resume an upload.")
		Payload(func() {
			Reference(TUSCoreHeaders)
			Reference(TUSExtensionHeaders)
			ID()
			Attribute("tusResumable")
			Attribute("uploadOffset")
			Attribute("uploadChecksum")
			Attribute("content", Bytes, "Initial upload content")
			Required("tusResumable", "uploadOffset", "content")
		})

		Result(func() {
			Reference(TUSCoreHeaders)
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
			Description("If the server receives a PATCH request against a non-existent resource it SHOULD return a 404 Not Found status.")
		})
		Error("InvalidChecksumAlgorithm", func() {
			Description("The checksum algorithm is not supported by the server.")
		})
		Error("ChecksumMismatch", func() {
			Description("The checksums mismatch.")
		})

		HTTP(func() {
			PATCH("/{id}")
			Header("tusResumable")
			Header("uploadOffset")
			Header("uploadChecksum")
			Body("content")
			Response(StatusNoContent, func() {
				Header("tusResumable")
				Header("uploadOffset")
				Header("uploadExpires")
			})
			Response("InvalidContentType", StatusUnsupportedMediaType)
			Response("InvalidOffset", StatusConflict)
			Response("NotFound", StatusNotFound)
			Response("InvalidChecksumAlgorithm", StatusBadRequest)
			Response("ChecksumMismatch", 460 /*StatusChecksumMismatch*/)
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
				Header("tusResumable")
				Header("tusVersion")
				Header("tusExtension")
				Header("tusMaxSize")
				Header("tusChecksumAlgorithm")
			})
		})
	})

	// creation and creation-wiht-upload extensions support.
	Method("post", func() {
		Description("Clients use the POST method against a known upload creation URL to request a new upload resource.")
		Payload(func() {
			Reference(TUSCoreHeaders)
			Reference(TUSExtensionHeaders)
			Attribute("tusResumable")
			Attribute("uploadDeferLength")
			Attribute("uploadChecksum")
			Attribute("tusMaxSize", UInt, "Length of the upload")
			Attribute("content", Bytes, "Initial upload content")
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
			Header("tusResumable")
			Header("uploadDeferLength")
			Header("uploadChecksum")
			Header("tusMaxSize")
			Body("content")
			Response(StatusCreated, func() {
				Header("location")
				Header("tusResumable")
				Header("uploadOffset")
				Header("uploadExpires")
			})
			Response("InvalidDeferLength", StatusBadRequest)
			Response("MaximumSizeExceeded", StatusRequestEntityTooLarge)
			Response("InvalidChecksumAlgorithm", StatusBadRequest)
			Response("ChecksumMismatch", 460 /*StatusChecksumMismatch*/)
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
			Reference(TUSCoreHeaders)
			Attribute("tusResumable")
			Required("tusResumable")
		})

		Error("NotFound", func() {
			Description("For all future requests to this URL, the Server SHOULD respond with the 404 Not Found.")
		})

		HTTP(func() {
			DELETE("/{id}")
			Header("tusResumable")
			Response(StatusNoContent, func() {
				Header("tusResumable")
			})
			Response("NotFound", StatusNotFound)
		})
	})
})

var TUSCoreHeaders = Type("TUSCoreHeaders", func() {
	Description("TUSCoreHeaders defines the tus core protocol headers used in both requests and responses.")
	Attribute("uploadOffset", UInt, "uploadOffset represents a byte offset within a resource.")
	Attribute("uploadLength", UInt, "uploadLength represents the size of the entire upload in bytes.")
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
	Attribute("tusMaxSize", UInt, "tusMaxSize represents the maximum allowed size of an entire upload in bytes.")
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
		Pattern(`[0-9a-f]{32}`)
	})
}

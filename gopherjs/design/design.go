package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("upload", func() {
	Description("This API exposes an image resource that allows uploading and downloading images")
	BasePath("/api")
})

var _ = Resource("image", func() {
	BasePath("/images")

	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload an image")
		Response(OK, ImageMedia)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Description("Show an image metadata")
		Params(func() {
			Param("id", Integer, "Image ID")
		})
		Response(OK, ImageMedia)
		Response(NotFound)
	})

	Files("/download/*filename", "images/") // Serve files from the "images" directory
})

var ImageMedia = MediaType("application/vnd.goa.examples.upload.image", func() {
	Description("Image metadata")
	TypeName("ImageMedia")
	Attributes(func() {
		Attribute("id", Integer, "Image ID")
		Attribute("filename", String, "Image filename")
		Attribute("uploaded_at", DateTime, "Upload timestamp")
		Required("id", "filename", "uploaded_at")
	})
	View("default", func() {
		Attribute("id")
		Attribute("filename")
		Attribute("uploaded_at")
	})
})

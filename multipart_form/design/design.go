package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("profile", func() {
	Title("A profile API")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("profiles", func() {
	Action("submit", func() {
		Routing(POST("profiles"))
		Payload(ProfilePayload)
		MultipartForm()
		Description("Post accepts a multipart form encoded request")
		Response(OK, ResultMedia)
	})
})

var ProfilePayload = Type("ProfilePayload", func() {
	Attribute("name", String, "Name")
	Attribute("birthday", DateTime, "Birthday")
	Attribute("icon", File, "Icon")
	Required("name", "birthday", "icon")
})

var ResultMedia = MediaType("application/vnd.goa.example.form", func() {
	TypeName("ResultMedia")
	Attributes(func() {
		Attribute("name", String, "Name")
		Attribute("birthday", DateTime, "Birthday")
		Required("name", "birthday")
	})
	View("default", func() {
		Attribute("name")
		Attribute("birthday")
	})
})

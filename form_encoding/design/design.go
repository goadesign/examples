package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("survey", func() {
	Title("A yes/no survey API demonstrating the use of x-www-form-urlencoded encoding in goa")
	Host("localhost:8080")
	Scheme("http")

	// This is the relevant part:
	Consumes("application/x-www-form-urlencoded", func() {
		Package("github.com/goadesign/goa/encoding/form")
	})
	Produces("application/x-www-form-urlencoded", func() {
		Package("github.com/goadesign/goa/encoding/form")
	})
})

var _ = Resource("survey_form", func() {
	Action("submit", func() {
		Routing(POST("survey/"))
		Payload(YesNoPayload)
		Description("Post accepts a form encoded request and returns a form encoded response")
		Response(OK, ResultMedia)
	})

})

var YesNoPayload = Type("YesNoPayload", func() {
	Attribute("name", String, "Voter name")
	Attribute("vote", String, "Yes or no", func() {
		Enum("yes", "no")
	})
	Required("name", "vote")
})

var ResultMedia = MediaType("application/vnd.goa.example.form", func() {
	TypeName("ResultMedia")
	Attributes(func() {
		Attribute("name", String, "Voter name")
		Attribute("message", String, "Thank you message")
		Required("name", "message")
	})
	View("default", func() {
		Attribute("name")
		Attribute("message")
	})
})

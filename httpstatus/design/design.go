package design

import (
	. "goa.design/goa/v3/dsl"
)

var GreetingResult = ResultType("application/vnd.hello", func() {
	Attribute("greeting", String, "The greeting message")
	Attribute("outcome", func() {
		Meta("struct:tag:json", "-")     // hide from response
		Meta("swagger:example", "false") // hide from swagger
		Meta("swagger:generate", "false")
	})
	Required("outcome", "greeting")
})

var _ = Service("hello", func() {
	Description("The hello service returns greetings with various statuses.")

	Method("hello", func() {
		Payload(func() {
			Attribute("greeting", String, "The greeting message")
			Required("greeting")
		})
		Result(GreetingResult)
		HTTP(func() {
			GET("/hello/{greeting}")
			Params(func() {
				Param("greeting")
			})
			Response(StatusCreated, func() {
				Tag("outcome", "created")
			})
			Response(StatusAccepted, func() {
				Tag("outcome", "accepted")
			})
			Response(StatusOK) // default response status
		})
	})
})

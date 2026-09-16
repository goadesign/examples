package design

import (
	. "goa.design/goa/v3/dsl"
)

var GreetingResult = ResultType("application/vnd.hello", func() {
	Attribute("greeting", String, "The greeting message")
	Attribute("outcome", String, "The HTTP response selected for the greeting", func() {
		Default("defaultStatus")
		Enum("created", "accepted", "defaultStatus")
	})
	Required("greeting")
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
			Response(StatusOK, func() {
				Body(func() {
					Attribute("greeting")
				})
			})
		})
	})
})

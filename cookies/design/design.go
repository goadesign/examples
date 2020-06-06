package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("session", func() {
	Title("Session service")
	Description("HTTP service that showcase the use of HTTP cookies in Goa services.")
})

var _ = Service("session", func() {
	Description("The session service illustrates managing user sessions with cookies.")

	Method("create_session", func() {
		Payload(func() {
			Attribute("name", String, "Name of session")
			Required("name")
		})

		Result(func() {
			Attribute("session_id", String, "Session ID")
			Attribute("message", String, "User message")
			Required("session_id", "message")
		})

		HTTP(func() {
			POST("/")
			Response(StatusOK, func() {
				Cookie("session_id:SID") // Return session ID in "SID" cookie
				CookieMaxAge(3600)       // Sessions last one hour
			})
		})
	})

	Method("use_session", func() {
		Payload(func() {
			Attribute("session_id", String, "Session ID")
			Required("session_id")
		})

		Result(func() {
			Attribute("message", String, "User message")
			Required("message")
		})

		HTTP(func() {
			GET("/")
			Cookie("session_id:SID") // Load session ID from "SID" cookie
			Response(StatusOK)
		})
	})

})

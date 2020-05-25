package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("concat", func() {
	Description(`The concat service performs operations on strings.

The service uses the CBOR binary serialization standard to encode responses.
It supports reading requests encoded with CBOR, JSON, XML or GOB.
`)

	Method("concat", func() {
		Payload(func() {
			Attribute("a", String, "Left operand")
			Attribute("b", String, "Right operand")
			Required("a", "b")
		})

		Result(String)

		HTTP(func() {
			GET("/concat/{a}/{b}")
		})
	})
})

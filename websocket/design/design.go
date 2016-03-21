package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("echo", func() {
	Title("Websocket echo server")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("echo", func() {
	Action("connect", func() {
		Routing(GET("echo"))
		Scheme("ws")
		Description("echo websocket server")
		Params(func() {
			Param("initial", String, "Initial message to echo")
		})
		Response(SwitchingProtocols)
	})

})

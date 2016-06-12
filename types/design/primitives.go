package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("prism", func() {
	Description("Resource prism contains artifical actions that showcase the use of primitive types in the DSL.")
	Action("show", func() {
		Description("Action show accepts one querystring parameter for each primitive type")
		Routing(GET(""))
		Params(func() {
			Param("bool_param", Boolean, "Boolean parameter")
			Param("int_param", Integer, "Integer parameter")
			Param("num_param", Number, "Number parameter")
			Param("string_param", String, "String parameter")
			Param("date_time_param", DateTime, "DateTime parameter")
			Param("uuid_param", UUID, "UUID parameter")
			Param("any_param", Any, "Any parameter")
		})
		Response(OK, PrismMedia)
	})
	Action("create", func() {
		Description("Action create accepts a payload with one member for each primitive type")
		Routing(POST(""))
		Payload(func() {
			Member("bool_member", Boolean, "Boolean member")
			Member("int_member", Integer, "Integer member")
			Member("num_member", Number, "Number member")
			Member("string_member", String, "String member")
			Member("date_time_member", DateTime, "DateTime member")
			Member("uuid_member", UUID, "UUID member")
			Member("any_member", Any, "Any member")
		})
	})
})

var PrismMedia = MediaType("application/vnd.goadesign.examples.prism", func() {
	Description("PrismMedia is a media type with one attribute per primitive type.")
	Attributes(func() {
		Attribute("bool_att", Boolean, "Boolean attribute")
		Attribute("int_att", Integer, "Integer attribute")
		Attribute("num_att", Number, "Number attribute")
		Attribute("string_att", String, "String attribute")
		Attribute("date_time_att", DateTime, "DateTime attribute")
		Attribute("uuid_att", UUID, "UUID attribute")
		Attribute("any_att", Any, "Any attribute")
	})
	View("default", func() {
		Attribute("bool_att")
		Attribute("int_att")
		Attribute("num_att")
		Attribute("string_att")
		Attribute("date_time_att")
		Attribute("uuid_att")
		Attribute("any_att")
	})
})

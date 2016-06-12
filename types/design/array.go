package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("array_prism", func() {
	BasePath("/array")
	Description("Resource array_prism contains artifical actions that showcase the use of the array type in the DSL.")
	Action("show", func() {
		Description(`Action show accepts one querystring array parameter for each primitive type.
Array parameters are constructed using all the values given to the same querystring key, e.g.:

	GET /array?int_param=1&int_param=2
`)
		Routing(GET(""))
		Params(func() {
			Param("bool_array", ArrayOf(Boolean), "Boolean array parameter")
			Param("int_array", ArrayOf(Integer), "Integer array parameter")
			Param("num_array", ArrayOf(Number), "Number array parameter")
			Param("string_array", ArrayOf(String), "String array parameter")
			Param("date_time_array", ArrayOf(DateTime), "DateTime array parameter")
			Param("uuid_array", ArrayOf(UUID), "UUID array parameter")
			Param("any_array", ArrayOf(Any), "Any array parameter")
		})
		Response(OK, ArrayPrismMedia)
	})
	Action("create", func() {
		Description("Action create accepts a payload with one array member for each primitive type")
		Routing(POST(""))
		Payload(func() {
			Member("bool_array", ArrayOf(Boolean), "Boolean array member")
			Member("int_array", ArrayOf(Integer), "Integer array member")
			Member("num_array", ArrayOf(Number), "Number array member")
			Member("string_array", ArrayOf(String), "String array member")
			Member("date_time_array", ArrayOf(DateTime), "DateTime array member")
			Member("uuid_array", ArrayOf(UUID), "UUID array parameter")
			Member("any_array", ArrayOf(Any), "Any array member")
		})
	})
})

var ArrayPrismMedia = MediaType("application/vnd.goadesign.examples.arrayprism", func() {
	Description("ArrayPrismMedia is a media type with one attribute per primitive type.")
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

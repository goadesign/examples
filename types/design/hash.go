package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Note: goa supports hash types with arbitrary types as key. However JSON and therefore Swagger
// do not allow for anything but strings as keys. This example generates Swagger so only uses
// hashes with string keys. Just know that if you don't use JSON to serialize requests or
// responses you can use hashes with keys of any type.
var _ = Resource("hash_prism", func() {
	BasePath("/hash")
	Description("Resource hash_prism contains artifical actions that showcase the use of the hash type in the DSL.")
	Action("show", func() {
		Description("Action show returns a media type with one hash member per primitive type, each using the primitive type as value and String as key")
		Routing(GET(""))
		Response(OK, HashPrismMedia)
	})
	Action("create", func() {
		Description("Action create accepts a payload with one hash member per primitive type, each using the type as value and String as key")
		Routing(POST(""))
		Payload(func() {
			Member("bool_val_hash", HashOf(String, Boolean), "Hash with Boolean value member")
			Member("int_val_hash", HashOf(String, Integer), "Hash with Integer value member")
			Member("num_val_hash", HashOf(String, Number), "Hash with Number value member")
			Member("String_val_hash", HashOf(String, String), "Hash with String value member")
			Member("date_time_val_hash", HashOf(String, DateTime), "Hash with DateTime value member")
			Member("uuid_val_hash", HashOf(String, UUID), "Hash with UUID value member")
			Member("any_val_hash", HashOf(String, Any), "Hash with Any value member")
		})
	})
})

var HashPrismMedia = MediaType("application/vnd.goadesign.examples.hashprism", func() {
	Description("HashPrismMedia is a media type with one hash member per primitive type, each using the type as value and String as key")
	Attributes(func() {
		Attribute("bool_val_hash", HashOf(String, Boolean), "Hash with Boolean value member")
		Attribute("int_val_hash", HashOf(String, Integer), "Hash with Integer value member")
		Attribute("num_val_hash", HashOf(String, Number), "Hash with Number value member")
		Attribute("String_val_hash", HashOf(String, String), "Hash with String value member")
		Attribute("date_time_val_hash", HashOf(String, DateTime), "Hash with DateTime value member")
		Attribute("uuid_val_hash", HashOf(String, UUID), "Hash with UUID value member")
		Attribute("any_val_hash", HashOf(String, Any), "Hash with Any value member")
	})
	View("default", func() {
		Attribute("bool_val_hash")
		Attribute("int_val_hash")
		Attribute("num_val_hash")
		Attribute("String_val_hash")
		Attribute("date_time_val_hash")
		Attribute("uuid_val_hash")
		Attribute("any_val_hash")
	})
})

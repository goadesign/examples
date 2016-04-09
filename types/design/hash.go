package design

var _ = Resource("hash_prism", func() {
	BasePath("/hash")
	Description("Resource hash_prism contains artifical actions that showcase the use of the hash type in the DSL.")
	Action("val_show", func() {
		Description("Action val_show returns a media type with one hash member per primitive type, each using the primitive type as value and String as key")
		Routing(GET("/val"))
		Response(OK, HashValPrismMedia)
	})
	Action("key_show", func() {
		Description("Action key_show returns a media type with one hash member per primitive type, each using the primitive type as key and String as value")
		Routing(GET("/key"))
		Response(OK, HashKeyPrismMedia)
	})
	Action("val_create", func() {
		Description("Action val_create accepts a payload with one hash member per primitive type, each using the type as value and String as key")
		Routing(POST("/val"))
		Payload(func() {
			Member("bool_val_hash", HashOf(String, Boolean), "Hash with Boolean value member")
			Member("int_val_hash", HashOf(String, Integer), "Hash with Integer value member")
			Member("num_val_hash", HashOf(String, Number), "Hash with Number value member")
			Member("String_val_hash", HashOf(String, String), "Hash with String value member")
			Member("date_val_time_hash", HashOf(String, DateTime), "Hash with DateTime value member")
			Member("any_val_hash", HashOf(String, Any), "Hash with Any value member")
		})
	})
	Action("key_create", func() {
		Description("Action key_create accepts a payload with one hash member per primitive type, each using the type as key and String as value")
		Routing(POST("/key"))
		Payload(func() {
			Member("bool_key_hash", HashOf(Boolean, String), "Hash with Boolean key member")
			Member("int_key_hash", HashOf(Integer, String), "Hash with Integer key member")
			Member("num_key_hash", HashOf(Number, String), "Hash with Number key member")
			Member("String_key_hash", HashOf(String, String), "Hash with String key member")
			Member("date_key_time_hash", HashOf(DateTime, String), "Hash with DateTime key member")
			Member("any_key_hash", HashOf(Any, String), "Hash with Any key member")
		})
	})
})

var HashValPrismMedia = MediaType("vnd.goadesign.examples.valprism", func() {
	Description("HashValPrismMedia is a media type with one hash member per primitive type, each using the type as value and String as key")
	Attributes(func() {
		Attribute("bool_val_hash", HashOf(String, Boolean), "Hash with Boolean value member")
		Attribute("int_val_hash", HashOf(String, Integer), "Hash with Integer value member")
		Attribute("num_val_hash", HashOf(String, Number), "Hash with Number value member")
		Attribute("String_val_hash", HashOf(String, String), "Hash with String value member")
		Attribute("date_val_time_hash", HashOf(String, DateTime), "Hash with DateTime value member")
		Attribute("any_val_hash", HashOf(String, Any), "Hash with Any value member")
	})
	View("default", func() {
		Attribute("bool_val_hash")
		Attribute("int_val_hash")
		Attribute("num_val_hash")
		Attribute("String_val_hash")
		Attribute("date_time_val_hash")
		Attribute("any_val_hash")
	})
})

var HashKeyPrismMedia = MediaType("vnd.goadesign.examples.keyprism", func() {
	Description("HashKeyPrismMedia is a media type with one hash member per primitive type, each using the type as key and String as value")
	Attributes(func() {
		Attribute("bool_key_hash", HashOf(Boolean, String), "Hash with Boolean key member")
		Attribute("int_key_hash", HashOf(Integer, String), "Hash with Integer key member")
		Attribute("num_key_hash", HashOf(Number, String), "Hash with Number key member")
		Attribute("String_key_hash", HashOf(String, String), "Hash with String key member")
		Attribute("date_key_time_hash", HashOf(DateTime, String), "Hash with DateTime key member")
		Attribute("any_key_hash", HashOf(Any, String), "Hash with Any key member")
	})
	View("default", func() {
		Attribute("bool_key_hash")
		Attribute("int_key_hash")
		Attribute("num_key_hash")
		Attribute("String_key_hash")
		Attribute("date_time_key_hash")
		Attribute("any_key_hash")
	})
})

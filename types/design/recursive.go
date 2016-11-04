package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("recursive_prism", func() {
	BasePath("/recursive")
	Description("Resource recursive_prism contains artifical actions that showcase the use of recursive data types")
	Action("array_array_show", func() {
		Description("Action array_array_show returns a media type with a member of type array containing another array")
		Routing(GET("/array_array"))
		Response(OK, ArrayArrayPrismMedia)
	})
	Action("array_hash_show", func() {
		Description("Action array_hash_show returns a media type with a member of type array containing a hash")
		Routing(GET("/array_hash"))
		Response(OK, ArrayHashPrismMedia)
	})
	Action("hash_array_show", func() {
		Description("Action hash_array_show returns a media type with a member of type hash containing an array")
		Routing(GET("/hash_array"))
		Response(OK, HashArrayPrismMedia)
	})
	Action("hash_hash_show", func() {
		Description("Action hash_hash_show returns a media type with a member of type hash containing another hash")
		Routing(GET("/hash_hash"))
		Response(OK, HashHashPrismMedia)
	})
})

var ArrayArrayPrismMedia = MediaType("application/vnd.goadesign.examples.arrayarrayprism", func() {
	Description("ArrayArrayPrismMedia is a media type with a array member that contains another array")
	Attributes(func() {
		Attribute("array_array", ArrayOf(ArrayOf(String)), "Array of array member")
	})
	View("default", func() {
		Attribute("array_array")
	})
})

var ArrayHashPrismMedia = MediaType("application/vnd.goadesign.examples.arrayhashprism", func() {
	Description("ArrayHashPrismMedia is a media type with a array member that contains a hash")
	Attributes(func() {
		Attribute("array_hash", ArrayOf(HashOf(String, String)), "Array of hashes")
	})
	View("default", func() {
		Attribute("array_hash")
	})
})

var HashArrayPrismMedia = MediaType("application/vnd.goadesign.examples.hasharrayprism", func() {
	Description("HashArrayPrismMedia is a media type with a hash member that contains an array")
	Attributes(func() {
		Attribute("hash_array", HashOf(String, ArrayOf(String)), "Hash of array member")
	})
	View("default", func() {
		Attribute("hash_array")
	})
})

var HashHashPrismMedia = MediaType("application/vnd.goadesign.examples.hashhashprism", func() {
	Description("HashHashPrismMedia is a media type with a hash member that contains another hash")
	Attributes(func() {
		Attribute("hash_hash", HashOf(String, HashOf(String, String)), "Hash of hash member")
	})
	View("default", func() {
		Attribute("hash_hash")
	})
})

var RecursivePrismMedia = MediaType("vnd.goadesign.examples.recursiveprism", func() {
	Description("RecursivePrismMedia is a media type that contains children of the same type")
	Attributes(func() {
		Attribute("children", CollectionOf("vnd.goadesign.examples.recursiveprism"), "Children")
	})
	View("default", func() {
		Attribute("children")
	})
})

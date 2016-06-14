package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// APIKey defines a security scheme using an API key (shared secret).  The scheme uses the
// "X-Shared-Secret" header to lookup the key.
var APIKey = APIKeySecurity("api_key", func() {
	Header("X-Shared-Secret")
})

var _ = Resource("api_key", func() {
	Description("This resource uses an API key to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(APIKey)

	Action("secure", func() {
		Description("This action is secure with the api_key scheme")
		Routing(GET("/api_key"))
		Response(OK)
		Response(Unauthorized)
	})

	Action("unsecure", func() {
		Description("This action does not require auth")
		Routing(GET("/api_key/unsecure"))
		NoSecurity()
		Response(OK)
	})
})

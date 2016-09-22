package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("An example of a goa service deployed behind Google Cloud Endpoints - https://cloud.google.com/endpoints")
	Host("goa-endpoints.appspot.com")
	Scheme("http")
})

// APIKey defines a security scheme using an API key (shared secret).
var APIKey = APIKeySecurity("api_key", func() {
	Query("key")
})

// JWT defines a security scheme using Google Endpoints JWT.
var JWT = OAuth2Security("jwt", func() {
	// Dummy value to make OpenAPI spec valid, Endpoints take care of implementation.
	ImplicitFlow("/auth")
})

var _ = Resource("operands", func() {
	Action("add", func() {
		Security(APIKey)
		Routing(GET("add/:left/:right"))
		Description("add returns the sum of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "text/plain")
	})
})

var _ = Resource("auth", func() {
	DefaultMedia(Auth)
	BasePath("/auth")

	Action("basic", func() {
		Security("api_key")
		Routing(GET("info/basic"))
		Response(OK)
	})

	Action("jwt", func() {
		Security("jwt", func() {
			// Swagger extensions as per https://cloud.google.com/endpoints/docs/authenticating-users
			Metadata("swagger:extension:x-issuer", "jwt-client.endpoints.sample.google.com")
			Metadata("swagger:extension:x-jwks_uri", "https://www.googleapis.com/service_accounts/v1/jwk/account-1@goa-swagger.iam.gserviceaccount.com")
		})
		Routing(GET("info/jwt"))
		Response(OK)
	})
})

// Auth is the auth info media type.
var Auth = MediaType("application/vnd.goa-cellar.auth+json", func() {
	Description("User info extracted from security token")
	TypeName("Auth")

	Attributes(func() {
		Attribute("issuer", String, "Token issuer")
		Attribute("id", String, "User ID")
		Attribute("email", String, "User email", func() {
			Format("email")
		})
		Required("id")
	})

	View("default", func() {
		Attribute("issuer")
		Attribute("id")
		Attribute("email")
	})
})

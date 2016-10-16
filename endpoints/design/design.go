package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("An example of a goa service deployed behind Google Cloud Endpoints - https://cloud.google.com/endpoints")
	Host("goa-endpoints.appspot.com")
	Scheme("https")

	Origin("http://swagger.goa.design", func() {
		Methods("GET", "POST")
		Headers("*")
	})
})

// APIKey defines a security scheme using an API key (shared secret).
var APIKey = APIKeySecurity("api_key", func() {
	Query("key")
})

// JWT defines a security scheme using Google Endpoints JWT.
var JWT = OAuth2Security("google_jwt", func() {
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
		Routing(GET("/basic"))
		Response(OK)
	})

	Action("jwt", func() {
		Security(JWT, func() {
			// Swagger extensions as per https://cloud.google.com/endpoints/docs/authenticating-users
			Metadata("swagger:extension:x-issuer", "client.goa-endpoints.appspot.com")
			Metadata("swagger:extension:x-jwks_uri", "https://www.googleapis.com/service_accounts/v1/jwk/endpoints-demo@goa-endpoints.iam.gserviceaccount.com")
		})
		Routing(GET("/jwt", func() {
			Metadata("swagger:extension:x-security", `json:[{"google_jwt":{"audiences":["goa-endpoints.appspot.com"]}}]`)
		}))
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

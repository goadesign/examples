package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("openapi", func() {
	// Does not work.
	Files("/openapi.json", "openapiSpecial.yaml")
	Files("/openapi3.json", "openapiSpecial.yaml")

	// Works
	Files("/openapi.yaml", "gen/openapiSpecial.yaml")
	Files("/openapi3.yaml", "gen/openapiSpecial.yaml")
})

package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("openapi", func() {
	Files("/openapi.json", "gen/http/openapi.json")
	Files("/openapi.yaml", "gen/http/openapi.yaml")
	Files("/openapi3.json", "gen/http/openapi3.json")
	Files("/openapi3.yaml", "gen/http/openapi3.yaml")
})

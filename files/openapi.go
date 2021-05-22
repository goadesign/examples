package openapiapi

import (
	"embed"
	"log"

	openapi "goa.design/examples/files/gen/openapi"
)

// openapi service example implementation.
// The example methods log the requests and return zero values.
type openapisrvc struct {
	logger *log.Logger
}

// NewOpenapi returns the openapi service implementation.
func NewOpenapi(logger *log.Logger) openapi.Service {
	return &openapisrvc{logger}
}

var (
	//go:embed gen/http/openapi.yaml
	OpenAPI embed.FS

	//go:embed gen/http/openapi3.json
	//go:embed gen/http/openapi3.yaml
	OpenAPI3 embed.FS
)

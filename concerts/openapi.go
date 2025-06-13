package concerts

import (
	"embed"
	"io/fs"
	"net/http"
)

// Embed the OpenAPI specifications at compile time
//
//go:embed gen/http/openapi3.json gen/http/openapi3.yaml
var openAPIFiles embed.FS

// OpenAPIFileSystem returns an http.FileSystem that serves the embedded OpenAPI files
func OpenAPIFileSystem() http.FileSystem {
	// Create a sub-filesystem from gen/http directory
	fsys, _ := fs.Sub(openAPIFiles, "gen/http")
	return http.FS(fsys)
}
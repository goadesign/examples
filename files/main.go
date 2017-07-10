package main

import (
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/goadesign/examples/files/app"
	"github.com/goadesign/examples/files/public/swagger"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("files")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "schema" controller
	c := NewSchemaController(service)
	app.MountSchemaController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	// You can override FileSystem of the controller.
	// For example using github.com/elazarl/go-bindata-assetfs is like below.
	c2.FileSystem = func(dir string) http.FileSystem {
		return &assetfs.AssetFS{
			Asset:     swagger.Asset,
			AssetDir:  swagger.AssetDir,
			AssetInfo: swagger.AssetInfo,
			Prefix:    dir,
		}
	}
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

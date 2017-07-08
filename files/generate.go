//go:generate goagen -d github.com/goadesign/examples/files/design app
//go:generate goagen -d github.com/goadesign/examples/files/design main
//go:generate goagen -d github.com/goadesign/examples/files/design client
//go:generate goagen -d github.com/goadesign/examples/files/design swagger -o public
//go:generate goagen -d github.com/goadesign/examples/files/design schema -o public
//go:generate go-bindata -ignore bindata.go -pkg swagger -o public/swagger/bindata.go ./public/swagger/...
package main

//go:generate goagen -d github.com/goadesign/examples/gopherjs/design app
//go:generate goagen -d github.com/goadesign/examples/gopherjs/design main
//go:generate goagen -d github.com/goadesign/examples/gopherjs/design client
//go:generate goagen -d github.com/goadesign/examples/gopherjs/design swagger
//go:generate gopherjs build github.com/goadesign/examples/gopherjs/public -o public/website.js
//go:generate mkdir -p images
package main

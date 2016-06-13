# Gopherjs Example

This example shows how to build an API that allows uploading and downloading assets. It also provides
a small website to show how to use the client with gopherjs

The file [generate.go](generate.go) contains `go generate` directives that uses `goagen` to generate
the necessary files. Run it with:

```bash
go generate
```

Then you need to generate the javascript code.
```bash
gopherjs build github.com/goadesign/examples/gopherjs/public -o public/website.js
```

Then compile and run the service:

```bash
go build
./gopherjs
```

Images can be uploaded by sending a multipart mime POST request to `/api/images`. The part
containing the file content must be named "file". This API returns the image metadata in the
response including an image ID that can be used to show the image medata and download it.


the website can be accessed by http://localhost:8080/index.html

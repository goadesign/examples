# Files

This example shows how to build an API that serves static assets.

The file [generate.go](generate.go) contains `go generate` directives that uses `goagen` to generate
the necessary files. Run it with:

```bash
go generate
```

Then compile and run the service:

```bash
go build
./files
```

## Resources

This example contains the following two resources.

* schema
* swagger

### Schema

The `schema` resource contains a single endpoint which serves the static files contained in the `public/schema` directory. This means the directory must exist and its relative path must be `public/schema`.

### Swagger

The `swagger` resource also contains a single endpoint however it leverages [https://github.com/elazarl/go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs) to serve assets embedded within the compiled binary.

* [go-bindata](https://github.com/jteeuwen/go-bindata) is

    > A small utility which generates Go code from any file. Useful for embedding binary data in a Go program.

* [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs)

    > Serves embedded files from `jteeuwen/go-bindata` with `net/http`.

Using these packages alleviates the need for deploying the `public/swagger` directory.

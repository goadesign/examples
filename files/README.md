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

This example contains following two resources.

* schema
* swagger

### Schema

Schema is served from native file system so you have to deploy the binary with `public/schema` directory and keep their path relationship.

### Swagger

Swagger is served from file system of [https://github.com/elazarl/go-bindata-assetfs](github.com/elazarl/go-bindata-assetfs).

* [go-bindata](https://github.com/jteeuwen/go-bindata) is

    > A small utility which generates Go code from any file. Useful for embedding binary data in a Go program.

* [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs)

    > Serves embedded files from `jteeuwen/go-bindata` with `net/http`.

When you use these packages, the binary contains static assets. It means you don't need to deploy `public/swagger` directory.

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

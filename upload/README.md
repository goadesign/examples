# Upload Example

This example shows how to build an API that allows uploading and downloading assets.

The file [generate.go](generate.go) contains `go generate` directives that uses `goagen` to generate
the necessary files. Run it with:

```bash
go generate
```

Then compile and run the service:

```bash
go build
./upload
```

Images can be uploaded by sending a multipart mime POST request to `/api/images`. The part
containing the file content must be named "file". This API returns the image metadata in the
response including an image ID that can be used to show the image medata and download it.

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

Images can be uploaded by sending a multipart mime POST request to `/api/images`. 
The API returns the images metadata in the response including an image IDs that 
can be used to show the image medata and download it.

Using CURL as an example files can be uploaded with:

```
curl -XPOST http://localhost:8080/api/images -F"image1=@foo.png" -F"image2=@bar.png"
```

where `foo.png` and `bar.png` are files that live in the current directory.

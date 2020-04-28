# File upload using the tus protocol.

This example implements a [tus](https://tus.io) compliant file upload mechanism.
tus is a protocol built on top of HTTP that allows for uploads to be resumed.
There are a number of client libraries that support tus written in different
languages such as [go-tus](https://github.com/eventials/go-tus) or
[tus-js-client](https://github.com/tus/tus-js-client). This example illustrates
how to write a Goa service that can be used in conjunction with such clients to
implement resumable uploads.

## Methods

The [tus core protocol](https://tus.io/protocols/resumable-upload.html#requests)
requires the server to implement a number of HTTP methods:

* `HEAD` requests should return status information on on-going uploads.
* `PATCH` requests are used to do the actual upload.
* `OPTIONS` requests allow for discovery (e.g. tus protocol version, supported
  extensions)

The protocol also suggests implementing additional HTTP methods:

* `POST` to create the upload target (`creation` and `creation-with-upload`
  extensions)
* `DELETE` to cancel uploads (`termination` extension).

## Extensions

This example implements the tus core protocol as well as the following
extensions:

* `creation`
* `creation-with-upload`
* `creation-defer-length`
* `expiration`
* `checksum`
* `termination`

This example supports version `1.0.0` of the tus protocol.

## Usage

This example implementation can be adapted for "real production" scenarios by
doing two modifications:
1. Modifying [the function](https://github.com/goadesign/examples/blob/master/tus/cmd/upload/main.go#L71-L77)
   that creates the writer used to write incoming bytes. The default implementation writes to local disk.
2. Providing an implementation of the [metadata store](https://github.com/goadesign/examples/blob/master/tus/persist/store.go)
   interface that persists the data instead of simply using an in-memory map.
   This data is what makes is possible to resume uploads and thus must be shared
   by all application servers that participate to the production cluster.

## Limitations

This implementation does not support the [Concatenation](https://tus.io/protocols/resumable-upload.html#concatenation)
extension. There is no fundamental reason why it couldn't (other than time!)
and would make a great addition!

## Example

Build and start the upload server:
```bash
cd $GOPATH/src/goa.design/examples/tus/cmd/upload
go build; and ./upload
[tus] 17:19:14 HTTP "Head" mounted on HEAD /upload/{id}
[tus] 17:19:14 HTTP "Patch" mounted on PATCH /upload/{id}
[tus] 17:19:14 HTTP "Options" mounted on OPTIONS /upload
[tus] 17:19:14 HTTP "Post" mounted on POST /upload
[tus] 17:19:14 HTTP "Delete" mounted on DELETE /upload/{id}
[tus] 17:19:14 HTTP server listening on "localhost:8080"
```

Build the client and show usage:
```bash
cd $GOPATH/src/goa.design/examples/tus/cmd/upload-cli
go build
./upload-cli --help
./upload-cli is a command line client for the tus upload API.

Usage:
    ./upload-cli [-host HOST][-url URL][-timeout SECONDS][-verbose|-v] SERVICE ENDPOINT [flags]

    -host HOST:  server host (development). valid values: development
    -url URL:    specify service URL overriding host URL (http://localhost:8080)
    -timeout:    maximum number of seconds to wait for response (30)
    -verbose|-v: print request and response details (false)

Commands:
    tus (head|patch|options|post|delete)

Additional help:
    ./upload-cli SERVICE [ENDPOINT] --help

Example:
    ./upload-cli tus head --id "6m6dfuts16k6ac7a5gna" --tus-resumable "1.0.0"
```

Perform an upload of the file `images/goa.png`:
```bash
./upload-cli --url http://localhost:8080 tus post --tus-resumable "1.0.0" --stream ../../images/goa.png --upload-defer-length 1
```
*note:* here we are uploading the file in a single chunk. If the image was big we would cut it in chunks and upload each one individually using a series of POST requests and specifying the offset each time.

Retrieve the status of the upload (**replace `bqjnf2cl42v556c7du4g` with the value you got above**):
```bash
./upload-cli --url http://localhost:8080 tus head --tus-resumable "1.0.0" --id bqjnf2cl42v556c7du4g
```

Complete the upload (again replacing the id with the proper value):
```bash
./upload-cli --url http://localhost:8080 tus delete --tus-resumable "1.0.0" --id bqjnf2cl42v556c7du4g
```

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
1. Modifying [the function](https://github.com/goadesign/examples/blob/tus/tus/cmd/upload/main.go#L71-L77)
   that creates the writer used to write incoming bytes. The default implementation writes to local disk.
2. Providing an implementation of the [metadata store](https://github.com/goadesign/examples/blob/tus/tus/persist/store.go)
   interface that persists the data instead of simply using an in-memory map.
   This data is what makes is possible to resume uploads and thus must be shared
   by all application servers that participate to the production cluster.

## Limitations

This implementation does not support the [Concatenation](https://tus.io/protocols/resumable-upload.html#concatenation)
extension. There is no fundamental reason why it couldn't (other than time!)
and would make a great addition!

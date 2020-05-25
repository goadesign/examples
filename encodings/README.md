# Goa Custom Encodings Examples

This directory contains two examples showing two different ways to implement
custom encodings in Goa. The Goa default encoders and decoders support the
JSON, XML and GOB serializations provided by the Go standard library. Goa
makes it possible to provide any arbitrary encoder and decoder to the
generated code via the generated HTTP server factory functions. Goa also
makes it possible to implement arbitrary content type negotation algorithms
(that is inferring the proper encoding to be used in the response given the
request). The default response encoder factory
[function](https://github.com/goadesign/goa/blob/v3/http/encoding.go#L99)
implements a basic content type negotatiation that simply matches the request
`Accept` header with a set of predefined supported mime types.

The `text` example illustrates how to specify the response content type in
the design thereby informing the response encoder factory function. In this
example the design specifies `text/plain` which
[forces](https://github.com/goadesign/goa/blob/v3/http/encoding.go#L109) the
default Goa content-type negotiation algorithm to use a text encoder.

The `cbor` example shows how to add a completely new type of encodings to Goa
services. The details on how to do that are described on the
[website](https://goa.design/implement/encoding/). To summarize: the
generated HTTP server and client factory functions accept request and
response encoder and decoder factory fucntions and invoke them on each
request appropriately.

The two functions used server side are:

```go
// RequestDecoder returns a HTTP request body decoder suitable for the given
// request.
func RequestDecoder(r *http.Request) Decoder 

// ResponseEncoder returns a HTTP response encoder leveraging the mime type
// set in the context under the AcceptTypeKey or the ContentTypeKey if any.
func ResponseEncoder(ctx context.Context, w http.ResponseWriter) Encoder
```

The two functions used client side are:

```go
// RequestEncoder returns a HTTP request encoder.
func RequestEncoder(r *http.Request) Encoder

// ResponseDecoder returns a HTTP response decoder.
func ResponseDecoder(resp *http.Response) Decoder
```

The `cbor` example top level `encoding.go` file implements these 4 functions
leveraging the [https://github.com/fxamacker/cbor] package. The
[server](https://github.com/goadesign/examples/tree/master/encodings/cbor/cmd/cbor/http.go#L32)
and
[client](https://github.com/goadesign/examples/tree/master/encodings/cbor/cmd/cbor-cli/http.go#L28)
main packages leverage these functions to create the HTTP server and client.

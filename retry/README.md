# Idempotent Client Retries

This example shows how a Goa service declares that repeating a method call with
the exact same input is safe. Generated HTTP and gRPC clients use that contract
to retry one temporary failure automatically.

The `get_message` implementation intentionally fails the first call for each
`id`. The generated client waits briefly, sends the same payload again,
and returns the successful response. The CLI and service contain no retry loop.

## Design

The method declares two independent facts:

```go
Method("get_message", func() {
    Idempotent()

    Error("unavailable", func() {
        Temporary()
    })

    // Payload, result, HTTP, and gRPC definitions...
})
```

`Idempotent()` says replaying the call is safe. `Temporary()` says a particular
service error may succeed when tried again. Goa retries only when both the
method and the failure permit it. Transport failures such as an HTTP connection
reset or gRPC `Unavailable` also qualify for methods marked `Idempotent()`.

Goa projects the method contract into each generated transport:

- `gen/http/retry/client/client.go` wraps the typed endpoint with
  `goa.RetryEndpoint`.
- `gen/grpc/retry/client/client.go` uses the same endpoint retry behavior.
- `gen/grpc/retry/pb/goagen_retry_retry.proto` sets protobuf
  `idempotency_level = IDEMPOTENT`.
- `gen/http/openapi*.{json,yaml}` adds `x-goa-idempotent: true`.

Automatic endpoint retries apply to unary calls with generated, replayable
request bodies. Goa does not replay established streams, WebSocket or
Server-Sent Events connections, multipart requests, or custom raw bodies.

## Generate and build

```sh
make gen
go build ./cmd/retry
go build ./cmd/retry-cli
```

Run the server in one terminal:

```sh
go run ./cmd/retry
```

Call it over HTTP:

```sh
go run ./cmd/retry-cli \
  --url=http://localhost:18080 \
  retry get-message --id http-demo
```

Call the same method over gRPC with a different demonstration request:

```sh
go run ./cmd/retry-cli \
  --url=grpc://localhost:18081 \
  retry get-message --message '{"id":"grpc-demo"}'
```

Each command succeeds even though the server reports an intentional temporary
failure on its first attempt:

```json
{
    "Message": "request http-demo succeeded after an automatic retry"
}
```

The service remembers request IDs in memory. Reusing an ID returns the same
successful result immediately; restarting the server resets the demonstration.

# Multipart Form Encoding

This example shows how to use encoding multipart/form-data (a.k.a. multipart form encoding) in a
goa service. The example uses multipart form encoding for the request payload.

The instructions below assume that goagen has been installed:

```
go get github.com/goadesign/goa/goagen
```

## Setup and Code Generation

The file [profiles.go](profiles.go) contains a `go:generate` comment which invokes `goagen` to
generate the source of the example from the `design` package. `go get` the example and run 
`go generate` to produce the entire source:

```
go get github.com/goadesign/examples/multipart_form
cd $GOPATH/src/github.com/goadesign/examples/multipart_form
go generate
```

> *Note*: `go get` displays an error because it cannot find the package `github.com/goadesign/examples/multipart_form/app` - this is expected as this package is generated and not committed. The error can be safely ignored.

## Running the Example

The example can then be built and run, nothing special:
```
go build -o profiles
./profiles
```

## Compile and Run the Client

Compile and run the client in a different terminal:

```
cd ./tool/profile-cli
go build
./profile-cli submit profiles --payload '{"name":"me","birthday":"2006-01-02T15:04:05Z"}' --dump
2018/03/15 17:15:40 [INFO] started id=Zs2fVjFT POST=http://localhost:8080/profiles
2018/03/15 17:15:40 [INFO] request headers Content-Type=multipart/form-data; boundary=ba3e72204014e59f22da6646866202df401d8ef7590a2ad53e928b5cd0da User-Agent=profile-cli/0
2018/03/15 17:15:40 [INFO] request body=--ba3e72204014e59f22da6646866202df401d8ef7590a2ad53e928b5cd0da
Content-Disposition: form-data; name="birthday"

2006-01-02T15:04:05Z
--ba3e72204014e59f22da6646866202df401d8ef7590a2ad53e928b5cd0da
Content-Disposition: form-data; name="name"

me
--ba3e72204014e59f22da6646866202df401d8ef7590a2ad53e928b5cd0da--
2018/03/15 17:15:40 [INFO] completed id=Zs2fVjFT status=200 time=2.216568ms
2018/03/15 17:15:40 [INFO] response headers Content-Length=48 Content-Type=application/vnd.goa.example.form Date=Thu, 15 Mar 2018 08:15:40 GMT
2018/03/15 17:15:40 [INFO] response body={"birthday":"2006-01-02T15:04:05Z","name":"me"}
```

Note how the body of the request is multiform form encoded.


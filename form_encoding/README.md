# Form Encoding

This example shows how to use encoding application/x-www-form-urlencoded (a.k.a. form encoding) in a
goa service. The example uses form encoding for both the request payload and response media type.

The instructions below assume that goagen has been installed:

```
go get github.com/goadesign/goa/goagen
```

## Setup and Code Generation

The file [survey_form.go](survey_form.go) contains a `go:generate` comment which invokes `goagen` to
generate the source of the example from the `design` package. `go get` the example and run 
`go generate` to produce the entire source:

```
go get github.com/goadesign/examples/form_encoding
cd $GOPATH/src/github.com/goadesign/examples/form_encoding
go generate
```

> *Note*: `go get` displays an error because it cannot find the package `github.com/goadesign/examples/form_encoding/app` - this is expected as this package is generated and not committed. The error can be safely ignored.

## Running the Example

The example can then be built and run, nothing special:
```
go build -o survey
./survey
```

## Compile and Run the Client

Compile and run the client in a different terminal:

```
cd ./tool/survey-cli
go build
./survey-cli submit survey_form --payload '{"name":"me","vote":"no"}' --dump
2016/06/11 23:03:24 [INFO] started id=MRMd1Ayw POST=http://localhost:8080/survey
2016/06/11 23:03:24 [INFO] request headers User-Agent=survey-cli/0
2016/06/11 23:03:24 [INFO] request body=name=me&vote=no
2016/06/11 23:03:24 [INFO] completed id=MRMd1Ayw status=200 time=1.446369ms
2016/06/11 23:03:24 [INFO] response headers Content-Type=application/vnd.goa.example.form Date=Sun, 12 Jun 2016 06:03:24 GMT Content-Length=26
2016/06/11 23:03:24 [INFO] response body=message=nay+%3A%28&name=me
```

Note how both the bodies of the request and response are form encoded.


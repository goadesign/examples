# Getting Started Guide

This is the example used in the [Getting Started Guide](https://goa.design/learn/guide/). Note that
apart from [design/design.go](design/design.go) the only code that is *not* generated is the body
of the [show](bottle.go) `bottle` controller action method.

The instructions below assume that goagen has been installed:
```
go get github.com/goadesign/goa/goagen
```

## Setup and Code Generation

The file [bottle.go](bottle.go) contains a `go:generate` comment which invokes `goagen` to
generate the source of the example from the `design` package. `go get` the example and run
`go generate` to produce the entire source:
```
go get github.com/goadesign/examples/guide
cd $GOPATH/src/github.com/goadesign/examples/guide
go generate
```

> *Note*: `go get` displays an error because it cannot find the pagkage
> `github.com/goadesign/examples/guide/app` - this is expected as this package is generated and not
> committed. The error can be safely ignored.

## Running the Example

The example can then be built and run, nothing special:
```
go build
./guide
2016/04/10 15:58:39 [INFO] mount ctrl=Bottle action=Show route=GET /bottles/:bottleID
2016/04/10 15:58:39 [INFO] mount file name=swagger/swagger.json route=GET /swagger.json
2016/04/10 15:58:39 [INFO] listen transport=http addr=:8080
```

## Compile and Run the Client

Compile and run the client in a different terminal:
```
cd client/cellar-cli
go build
./cellar-cli show bottle /bottles/1
2016/04/10 16:00:53 [INFO] started id=TNt3gVdt GET=http://localhost:8080/bottles/1
2016/04/10 16:00:53 [INFO] completed id=TNt3gVdt status=200 time=1.046666ms
{"href":"/bottles/1","id":1,"name":"Bottle #1"}
```

# adder

This is the example used in the README. Note that apart from [design/design.go](design/design.go)
the only line of code that is *not* generated is the line that does the addition and writes the
response in [operands.go](operands.go).

The instructions below assume that goagen has been installed:
```
go get github.com/goadesign/goa/goagen
```

## Setup and Code Generation

The file [operands.go](operands.go) contains a `go:generate` comment which invokes `goagen` to
generate the source of the example from the `design` package. `go get` the example and run
`go generate` to produce the entire source:
```
go get github.com/goadesign/examples/adder
cd $GOPATH/src/github.com/goadesign/examples/adder
go generate
```

> *Note*: `go get` displays an error because it cannot find the pagkage `github.com/goadesign/examples/adder/app` - this is expected as this package is generated and not committed. The error can be safely ignored.

## Running the Example

The example can then be built and run, nothing special:
```
go build
./adder
```

## Compile and Run the Client

Compile and run the client in a different terminal:
```
cd client/adder-cli
go build
./adder-cli add operands /add/1/2
2016/03/21 00:33:10 [INFO] started id=nclom9xa GET=http://localhost:8080/add/1/2
2016/03/21 00:33:10 [INFO] completed id=nclom9xa status=200 time=20.102916ms
3
```

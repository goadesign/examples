# Websocket

This example sets up a websocket echo server. Clients create websocket connections to `/echo`. Any
websocket message sent to that connection results in a response containing the same data.

## Design

Setting up a websocket endpoint in goa is as simple as setting the scheme of an action to `ws` or
`wss`. In this example the `echo` action is defined as:
```go
Action("add", func() {
	Routing(GET("echo"))
	Scheme("ws")
	Description("echo websocket server")
	Params(func() {
		Param("initial", String, "Initial message to echo")
	})
	Response(SwitchingProtocols)
})
```
Note how it's also possible to send parameters to the endpoint.
The best practice consists of defining a single response returning status 101. This example uses
the goa built-in `SwitchingProtocols` response. This response will not cause `goagen` to generate
a response method on the context but will still allow other generators such as Swagger to produce
the right output.

## Building

The [main.go](main.go) file contains a `go:generate` directive that uses `goagen` to generate the
code from the design. Clone, generate and build the code:
```
git clone https://github.com/goadesign/examples/websocket
cd $GOPATH/src/github.com/goadesign/examples/websocket
go generate
go build
```

## Running

Once the code is compiled and run the service:
```
./websocket
```

### Websocket Client

The instructions below use the `wsc` websocket command line client. Any other websocket client
would work as well. To use the `wsc` client install it with:
```
go install https://github.com/raphael/wsc
```

### Sending Websocket Messages

Open a new terminal and run the websocket command line client:
```
$ wsc ws://localhost:8080/echo?initial=first
2016/03/20 22:44:31 connecting to ws://localhost:8080/echo?initial=first...
2016/03/20 22:44:31 ready, exit with CTRL+C.
<< first
foo
>> foo
<< foo
^C
exiting
```
The first line above establishes the websocket connection. The first message to be returned is
`first` as specified in the `initial` querystring value. Typing `foo` causes the value to be sent
and received back from the goa websocket server. Finally CTRL+C exits the client.

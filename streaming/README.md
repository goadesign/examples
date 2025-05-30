# Chatter Service

The "chatter" service illustrates the **streaming endpoint** features in
Goa.

## Design

An endpoint becomes a streaming endpoint if any of the following DSLs are used
in the `Method` expression.

* `StreamingPayload` - client streams payload to the server
* `StreamingResult` - server streams result to the client

When both the above DSLs are defined in a `Method` expression, the endpoint
becomes a bidirectional stream. The syntax for the `StreamingPayload` and
`StreamingResult` DSLs are similar to the `Payload` and `Result` DSLs.

### `login` Endpoint

This is a non-streaming endpoint which authenticates the user using the
basic auth scheme and returns a valid JWT token to be sent by the other
endpoints.

```
$ ./chatter_cli chatter login --user "goa" --password "rocks"
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzQxOTg0NjgsIm5iZiI6MTQ0NDQ3ODQwMCwic2NvcGV
zIjpbInN0cmVhbTpyZWFkIiwic3RyZWFtOndyaXRlIl19.frlyMLZeSCSFthtaYe3tZYkg0nMqwREOj-55J6IUxyg"
$ export JWT_TOKEN=`echo $_`
```

### `listener` Endpoint

This endpoint illustrates a **streaming payload**. The client sends a payload
defined in `Payload` DSL containing the JWT token. Once auth succeeds, the
client streams the payload defined in `StreamingPayload` DSL to the server.
The server patiently listens until the client stops sending the payload.

```
$ ./chatter_cli chatter listener --token $JWT_TOKEN
Press Ctrl+D to stop chatting.
thanks for listening
you are very patient
```

### `summary` Endpoint

This endpoint is similar to the `listener` endpoint. The only difference is
the server responds back with a summary of all the messages sent by the client.

```
$ ./chatter_cli chatter summary --token $JWT_TOKEN
testing 1 2 3
check check check
[
    {
        "Message": "testing 1 2 3",
        "Length": 13,
        "SentAt": "2018-08-14T12:32:26-07:00"
    },
    {
        "Message": "check check check",
        "Length": 17,
        "SentAt": "2018-08-14T12:32:30-07:00"
    }
]
```

### `subscribe` Endpoint

This endpoint illustrates a **streaming result**. The client sends a payload
defined in the `Payload` DSL and that contains a JWT token. Once auth succeeds,
the server establishes a stream. The sends messages to the client using the
stream. The subscription ends when the client goes away or the server shuts
down.

```
$ ./chatter_cli chatter subscribe --token $JWT_TOKEN
# Now send messages to server in a separate `chatter_cli` process by calling
# the `summary`, `listener`, or `echoer` endpoints.
{
    "Message": "hello",
    "Action": "added",
    "AddedAt": "2018-08-14T12:32:30-07:00"
}
```

### `history` Endpoint

This endpoint illustrates a **streaming result** with views. The client sends
a payload defined in the `Payload` DSL and that contains a JWT token and an
optional "view" parameter. Once auth succeeds, the server sends results to the
client using the stream. Each result struct is first projected to the view
matching the "view" parameter.

```
$ ./chatter_cli chatter history --token $JWT_TOKEN --view tiny
{
    "Message": "thanks for listening",
    "Length": null,
    "SentAt": null
}
{
    "Message": "you are very patient",
    "Length": null,
    "SentAt": null
}
{
    "Message": "testing 1 2 3",
    "Length": null,
    "SentAt": null
}
{
    "Message": "check check check",
    "Length": null,
    "SentAt": null
}
```

### `echoer` Endpoint

This endpoint illustrates a **bidirectional streaming**. The client sends a
payload defined in `Payload` DSL containing the JWT token. Once auth
succeeds, the client and server communicates via the stream until one of them
quits. The server simply echoes the client messages.

```
$ ./chatter_cli chatter echoer --token $JWT_TOKEN
Press Ctrl+D to stop chatting.
say my name
"say my name"
stop repeating me 
"stop repeating me"
```

## Customizing HTTP Websocket Connections

Goa uses [gorilla websocket](https://godoc.org/github.com/gorilla/websocket)
underneath to implement streaming via websocket in the HTTP transport layer.  By
default, Goa uses the default
[`Upgrader`](https://godoc.org/github.com/gorilla/websocket#Upgrader) server
side to upgrade HTTP connection to a websocket connection and the
[`DefaultDialer`](https://godoc.org/github.com/gorilla/websocket#pkg-variables)
client side to dial a websocket connection.

Developers can use custom Upgrader and Dialer as shown below

```
// In server main.go

var (
  chatterServer *chattersvcsvr.Server
)
{
  eh := ErrorHandler(logger)

  // my custom websocket Upgrader
  myUpgrader := &websocket.Upgrader{
    ReadBufferSize: 512,
    WriteBufferSize: 512,
  }

  // pass a ConnConfigureFunc type as an argument to apply the connection
  // configurer to all the streaming endpoints
  chatterConfigurer := chattersvcsvr.NewConnConfigurer(myConfigurerFn)
  // or override the connection configurer for a specific streaming endpoint
  chatterConfigurer.SubscribeFn = mySubscriberConfigurerFn

  chatterServer = chattersvcsvr.New(chatterEndpoints, mux, dec, enc, eh, myUpgrader, chatterConfigurer)
}

// In client main.go

var (
  myDialer          *websocket.Dialer
  chatterConfigurer *chattersvcclient.ConnConfigurer
)
{
  chatterConfigurer = chattersvcclient.NewConnConfigurer(myConfigurerFn)
  myDialer = websocket.Dialer{
    ReadBufferSize: 512,
    WriteBufferSize: 512,
  }
}

endpoint, payload, err := cli.ParseEndpoint(
  scheme,
  host,
  doer,
  goahttp.RequestEncoder,
  goahttp.ResponseDecoder,
  debug,
  myDialer,
  chatterConfigurer,
)
```

Check out the `pingPonger` defined in `cmd/chatter/http.go` for how to use
`ConnConfigureFunc` to customize the websocket connection to send periodic
pings to the client and cancel the request if the client does not respond
with a pong.

## Gotchas

* Goa uses websockets to implement streaming in the HTTP transport layer.
The [websocket protocol](https://tools.ietf.org/html/rfc6455) has two parts,
an opening handshake and the data transfer. The opening handshake always
uses a `GET` request to the server to upgrade the HTTP connection to a
websocket connection. Therefore, even though your Goa HTTP endpoint in a
streaming method defines a verb other than `GET`, the first HTTP request is
always a `GET`. This could potentially lead to requests to your streaming
endpoint get routed to some other endpoint using the `GET` verb. To avoid this,
if your design defines a streaming method, the corresponding HTTP endpoint
should always use a `GET` method. Any other verb will generate a validation
error.

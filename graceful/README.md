# Graceful Example

This example shows how to setup a goa service to use the [graceful](https://github.com/tylerb/graceful)
server.

The graceful server enables graceful shutdown where upon receiving an interrupt signal (SIGINT,
SIGTERM) or when explicitly requested the server:

1. Disables keepalive connections.
2. Closes the listening socket, allowing another process to listen on that port immediately.
3. Starts a timer of timeout duration to give active requests a chance to finish.
4. When timeout expires, closes all active connections.
5. Closes the stopChan, waking up any blocking goroutines.
6. Returns from the function, allowing the server to terminate.

## Running the Example

The add action in the example sleeps 10s prior to returning the result and the graceful
timeout is set to 15s. Interrupting the server during a request triggers graceful shutdown: the
server quits after the request response is sent.

To run the example, first compile it:
```
git clone https://github.com/goadesign/examples/graceful
cd graceful
go build
```
Run it:
```
./graceful
```
Open another console and make a request:
```
curl localhost:8080/add/1/2
```
Interrupt the server by typing CTRL+C in the terminal running `graceful`. See how the request
completes then the server quits. Also note how sending new requests to the server after it is
interrupted fails right away as the server stops accepting new connections.

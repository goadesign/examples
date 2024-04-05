Sample endpoints that return multiple status codes in HTTP transport.
===

This sample changes the response status according to the contents of the request body. See the [Tag documentation](https://pkg.go.dev/goa.design/goa/v3/dsl#Tag) for details.

### OUTPUT:

Run
```shell
$ go run ./cmd/hello -http-port=8080
```

Access
```sh
$ curl -iii -XGET localhost:8080/hello/hello
HTTP/1.1 201 Created
Content-Type: application/json
Date: Fri, 05 Apr 2024 13:12:13 GMT
Content-Length: 21

{"greeting":"hello"}
```

```sh
$ curl -iii -XGET localhost:8080/hello/bye
HTTP/1.1 202 Accepted
Content-Type: application/json
Date: Fri, 05 Apr 2024 13:12:31 GMT
Content-Length: 19

{"greeting":"bye"}
```

```sh
$ curl -iii -XGET localhost:8080/hello/こんにちは
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 05 Apr 2024 13:12:41 GMT
Content-Length: 31

{"greeting":"こんにちは"}
```
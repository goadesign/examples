# Error Handling

This example illustrates how to return and handle errors in Goa. Refer to
[Error Handling](https://goa.design/implement/error_handling/) for additional
information.

## Design

The example defines a `calc` service and two methods `add` and `divide`. The
service defines a `timeout` error which applies to both methods. The `divide`
method also defines a method specific `div_by_zero` error.

The `timeout` error does not specify a type and thus uses the default
([ErrorResult](https://pkg.go.dev/goa.design/goa/v3/expr#ErrorResult)):

```go
Error("timeout") // Use default error type
```

The `div_by_zero` error specifies a custom error type `DivByZero`:

```go
Error("div_by_zero", DivByZero, "division by 0") // Use custom error type
```

## Implementation

Given the above Goa generates a function `MakeTimeout` and a Go struct
`DivByZero` both in the `calc` service Go package. The `MakeTimeout` function
makes it possible to return `timeout` errors given a Go error and is
leveraged by the service implementation to handle exceeded deadlines:

```go
	// Wait for results or timeout
	select {
	case <-ctx.Done():
		// Timeout triggered, return timeout error
		return nil, calc.MakeTimeout(ctx.Err())
	case res := <-c:
		// Division completed successfully
		return res, nil
    }
```

The `DivByZero` struct can be used to return division by zero errors:

```go
	if p.Divisor == 0 {
        return nil, &calc.DivByZero{Message: "divide by zero"}
    }
```

### Overridding Default Validation Errors

This example also illustrates how to override default validation error
responses. In this case the response returned for missing required field
errors. This is accomplished in the `main` package by providing a non-nil
error formatter function to the HTTP server `New` function:

```go
calcServer = calcsvr.New(calcEndpoints, mux, dec, enc, eh, FormatError)
```

The `FormatError` function is called by the generated code prior to writing
the HTTP response. The function takes the error as argument and returns the
struct that should be serialized in the response body. The struct retuned by
the error formatter must implement the
[Statuser](https://pkg.go.dev/goa.design/goa/v3/http#Statuser) interface
which defines a single method `StatusCode()` that returns the response
status code.

 The example uses a string to send the error back. The `FormatError` function
 checks whether the error is an instance of `ServiceError` and if so whether
 it corresponds to a missing field validation error. If that's the case it
 returns the custom error type otherwise it returns the default Goa error
 response by calling
 [NewErrorResponse](https://pkg.go.dev/goa.design/goa/v3/http#NewErrorResponse).
	
```go
// missingFieldError is the type used to serialize missing required field
// errors. It overrides the default provided by Goa.
type missingFieldError string

// StatusCode returns 400 (BadRequest).
func (missingFieldError) StatusCode() int { return http.StatusBadRequest }

// FormatError is the error formatter used to format error responses returned by
// the calc server.
func FormatError(err error) goahttp.Statuser {
	if serr, ok := err.(*goa.ServiceError); ok {
		switch serr.Name {
		case "missing_field":
			return missingFieldError(serr.Message)
		default:
			// Use Goa default
			return goahttp.NewErrorResponse(err)
		}
	}
	// Use Goa default for all other error types
	return goahttp.NewErrorResponse(err)
}
```

## Running the Example

First compile and start the server:

```bash
cd $GOPATH/src/goa.design/examples/error/cmd/calc
go build; ./calc --http-port 8080
```

This shoud produce output similar to:

```
[calcapi] 11:50:56 HTTP "Divide" mounted on POST /
[calcapi] 11:50:56 serving gRPC method calc.Calc/Divide
[calcapi] 11:50:56 HTTP server listening on "localhost:8000"
[calcapi] 11:50:56 gRPC server listening on "localhost:8080"
```

In a different terminal, compile and run the client:

```bash
cd $GOPATH/src/goa.design/examples/error/cmd/calc-cli
go build; ./calc-cli -url http://localhost:8000 calc divide --body '{"dividend":1,"divisor":1}'
{
    "Quotient": 1,
    "Reminder": 0
}

./calc-cli -v -url http://localhost:8000 calc divide --body '{"dividend":1,"divisor":0}'
> POST http://localhost:8000/
{"dividend":1,"divisor":0}

< 400 Bad Request
< Content-Length: 29
< Content-Type: application/json
< Date: Sun, 22 Nov 2020 21:27:36 GMT
< Goa-Error: div_by_zero
{"message":"divide by zero"}

DivByZero is the error returned when using value 0 as divisor.
```

Note that in that last example the message `DivByZero is the error...` is
produced by the generated client, the HTTP response is shown in the debug
output. The HTTP response body contains `{"message":"divide by zero"}` as
expected (a serialized `DivByZero` object). Now let's trigger a timeout
error:

```bash
./calc-cli -v -url http://localhost:8000 calc divide --body '{"dividend":42,"divisor":1}'
> POST http://localhost:8000/
{"dividend":42,"divisor":1}

< 504 Gateway Timeout
< Content-Length: 120
< Content-Type: application/json
< Date: Sun, 22 Nov 2020 21:33:37 GMT
< Goa-Error: timeout
{"name":"timeout","id":"_KL-qvO3","message":"context deadline exceeded","temporary":false,"timeout":true,"fault":false}

context deadline exceeded
```

Here we can see the server returning the default Goa error response type (and
as previously the message `context deadline exceeded` is produced by the
client). Note that the `timeout` field of the error response object is set to
`true`, this is because the design calls `Timeout()` in the `timeout` error
definition.

Finally we can check that the custom error type is used for missing required
field errors. We use [httpie](https://httpie.io/) to make the request since
the generated client knows the field is required and would fail the request:

```bash
http POST localhost:8000 dividend:=42
HTTP/1.1 400 Bad Request
Content-Length: 35
Content-Type: application/json
Date: Mon, 23 Nov 2020 03:49:50 GMT

"\"divisor\" is missing from body"
```

The response body consists of a string instead of the default Goa error
response object.

## Additional Information

Consult the
[error handling section](https://goa.design/implement/error_handling/) of
[goa.design](https://goa.design) to learn more about error handling in Goa.
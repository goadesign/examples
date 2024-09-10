// Code generated by goa v3.15.2, DO NOT EDIT.
//
// hello service
//
// Command:
// $ goa gen goa.design/examples/httpstatus/design

package hello

import (
	"context"

	helloviews "goa.design/examples/httpstatus/gen/hello/views"
)

// The hello service returns greetings with various statuses.
type Service interface {
	// Hello implements hello.
	HelloEndpoint(context.Context, *HelloPayload) (res *Hello, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "hello"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "hello"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"hello"}

// Hello is the result type of the hello service hello method.
type Hello struct {
	// The greeting message
	Greeting string
	Outcome  string `json:"-"`
}

// HelloPayload is the payload type of the hello service hello method.
type HelloPayload struct {
	// The greeting message
	Greeting string
}

// NewHello initializes result type Hello from viewed result type Hello.
func NewHello(vres *helloviews.Hello) *Hello {
	return newHello(vres.Projected)
}

// NewViewedHello initializes viewed result type Hello from result type Hello
// using the given view.
func NewViewedHello(res *Hello, view string) *helloviews.Hello {
	p := newHelloView(res)
	return &helloviews.Hello{Projected: p, View: "default"}
}

// newHello converts projected type Hello to service type Hello.
func newHello(vres *helloviews.HelloView) *Hello {
	res := &Hello{}
	if vres.Greeting != nil {
		res.Greeting = *vres.Greeting
	}
	if vres.Outcome != nil {
		res.Outcome = *vres.Outcome
	}
	return res
}

// newHelloView projects result type Hello to projected type HelloView using
// the "default" view.
func newHelloView(res *Hello) *helloviews.HelloView {
	vres := &helloviews.HelloView{
		Greeting: &res.Greeting,
		Outcome:  &res.Outcome,
	}
	return vres
}
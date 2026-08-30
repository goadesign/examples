// This file verifies that the custom HTTP error formatter changes missing-field
// errors without changing the response declared for division by zero.
package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	calcapi "goa.design/examples/error"
	gencalc "goa.design/examples/error/gen/calc"
	gencalcsvr "goa.design/examples/error/gen/http/calc/server"
	goahttp "goa.design/goa/v3/http"
)

func TestErrorResponses(t *testing.T) {
	service := calcapi.NewCalc(log.New(io.Discard, "", 0))
	endpoints := gencalc.NewEndpoints(service)
	mux := goahttp.NewMuxer()
	server := gencalcsvr.New(
		endpoints,
		mux,
		goahttp.RequestDecoder,
		goahttp.ResponseEncoder,
		nil,
		formatError,
	)
	gencalcsvr.Mount(mux, server)

	tests := []struct {
		name string
		body string
		want string
	}{
		{
			name: "missing divisor",
			body: `{"dividend":42}`,
			want: `"\"divisor\" is missing from body"` + "\n",
		},
		{
			name: "division by zero",
			body: `{"dividend":1,"divisor":0}`,
			want: `{"message":"divide by zero"}` + "\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.body))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			mux.ServeHTTP(response, request)

			if response.Code != http.StatusBadRequest {
				t.Errorf("status code = %d, want %d", response.Code, http.StatusBadRequest)
			}
			if body := response.Body.String(); body != test.want {
				t.Errorf("response body = %q, want %q", body, test.want)
			}
		})
	}
}

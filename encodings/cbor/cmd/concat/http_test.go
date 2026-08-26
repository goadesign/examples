// These tests verify that the example server uses its CBOR transport codecs.
package main

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/fxamacker/cbor/v2"
	"goa.design/clue/log"
	concatapi "goa.design/examples/encodings/cbor"
	genconcat "goa.design/examples/encodings/cbor/gen/concat"
)

// TestHTTPServerUsesCBOR checks the response produced by the configured HTTP
// server rather than testing the codec functions on their own.
func TestHTTPServerUsesCBOR(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("reserve HTTP address: %v", err)
	}
	address := listener.Addr().String()
	if err := listener.Close(); err != nil {
		t.Fatalf("release HTTP address: %v", err)
	}

	serverURL := &url.URL{Scheme: "http", Host: address}
	endpoints := genconcat.NewEndpoints(concatapi.NewConcat())
	logContext := log.Context(context.Background(), log.WithFormat(log.FormatJSON))
	ctx, cancel := context.WithCancel(logContext)
	var waitGroup sync.WaitGroup
	handleHTTPServer(ctx, serverURL, endpoints, &waitGroup, make(chan error, 1), false)
	t.Cleanup(func() {
		cancel()
		waitGroup.Wait()
	})

	response := waitForHTTPResponse(t, "http://"+address+"/concat/a/b")
	t.Cleanup(func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("close response body: %v", err)
		}
	})
	if got := response.Header.Get("Content-Type"); got != "application/cbor" {
		t.Fatalf("Content-Type = %q, want application/cbor", got)
	}

	var result string
	if err := cbor.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fatalf("decode CBOR response: %v", err)
	}
	if result != "ab" {
		t.Fatalf("result = %q, want ab", result)
	}
}

// waitForHTTPResponse waits until the server accepts requests or the test
// timeout expires.
func waitForHTTPResponse(t *testing.T, requestURL string) *http.Response {
	t.Helper()

	client := &http.Client{Timeout: 100 * time.Millisecond}
	deadline := time.Now().Add(2 * time.Second)
	for {
		response, err := client.Get(requestURL)
		if err == nil {
			return response
		}
		if time.Now().After(deadline) {
			t.Fatalf("GET %s: %v", requestURL, err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

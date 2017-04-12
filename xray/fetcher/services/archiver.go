package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/goadesign/examples/xray/archiver/client"
	goaclient "github.com/goadesign/goa/client"
	"github.com/goadesign/goa/middleware"
)

type (
	// Archiver is the client interface to the archiver service specific to
	// the fetcher service.
	Archiver interface {
		// Archive stores a HTTP response in the archiver service and
		// returns the corresponding resource href.
		Archive(ctx context.Context, status int, body string) (string, error)
		// Healthy returns true if the downstream service is responding
		// to health check requests with status code OK.
		// Healthy returns an error if the request cannot be made.
		Healthy(context.Context) (bool, error)
		// WaitUntilHealthy returns true if the downstream service
		// responds successfully to health check requests, false if the
		// given timeout is exceeded.
		// Health check requests are attempted every 100ms until a
		// successful response is received or the timeout triggers.
		// WaitUntilHealthy returns an error if the request cannot be
		// made.
		WaitUntilHealthy(context.Context, time.Duration) (bool, error)
	}

	// archiver is the client implementation.
	archiver struct {
		// host is the archiver service host:port
		host string
		// doer is the http client used to make requests.
		doer goaclient.Doer
	}
)

// NewArchiver returns a client to the archiver service.
func NewArchiver(host string, client *http.Client) Archiver {
	return &archiver{host, goaclient.HTTPClientDoer(client)}
}

// Archive stores a HTTP response in the archiver service and returns the
// corresponding resource href.
func (a *archiver) Archive(ctx context.Context, status int, body string) (string, error) {
	// Wrap client with xray to trace request
	c := client.New(middleware.TraceDoer(a.doer))
	c.Host = a.host

	// Create payload
	payload := &client.ArchivePayload{
		Status: status,
		Body:   body,
	}

	// Make request
	resp, err := c.ArchiveArchiver(ctx, "/archive", payload, "")

	// Check for errors
	if err != nil {
		return "", fmt.Errorf("failed to make request to archiver: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Archive request failed: %s", errorFromResponse(resp))
	}

	// Read response
	defer resp.Body.Close()
	var archive client.ArchiveMedia
	err = json.NewDecoder(resp.Body).Decode(&archive)
	if err != nil {
		return "", fmt.Errorf("Failed to decode response: %s", err)
	}

	// We're done
	return archive.Href, nil
}

// Healthy returns true if the downstream service is responding to health check
// requests with status code OK.
// Healthy returns an error if the request cannot be made.
func (a *archiver) Healthy(ctx context.Context) (bool, error) {
	// We chose not to trace health-check requests
	c := client.New(a.doer)
	c.Host = a.host
	resp, err := c.ShowHealth(ctx, "/health")
	if err != nil {
		return false, err
	}
	return resp.StatusCode == http.StatusOK, nil
}

// WaitUntilHealthy returns true if the downstream service responds successfully
// to health check requests, false if the given timeout is exceeded.
// Health check requests are attempted every 100ms until a successful response is
// received or the timeout triggers.
// WaitUntilHealthy returns an error if the request cannot be made.
func (a *archiver) WaitUntilHealthy(ctx context.Context, t time.Duration) (bool, error) {
	// Quick check first
	if ok, err := a.Healthy(ctx); ok || err != nil {
		return ok, err
	}

	// Now wait actively
	var (
		ticker  = time.NewTicker(100 * time.Millisecond)
		timeout = time.NewTimer(t)
	)
	defer ticker.Stop()
	defer timeout.Stop()
	for {
		select {
		case <-ticker.C:
			if ok, err := a.Healthy(ctx); ok || err != nil {
				return ok, err
			}
		case <-timeout.C:
			return false, nil
		case <-ctx.Done():
			return false, nil
		}
	}
}

// errorFromResponse formats the HTTP response into an error message.
func errorFromResponse(resp *http.Response) string {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var suffix string
	if len(body) > 0 {
		suffix = ": " + string(body)
	}
	return fmt.Sprintf("%d %s%s", resp.StatusCode, resp.Status, suffix)
}

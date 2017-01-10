package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goadesign/examples/xray/fetcher/app/test"
	"github.com/goadesign/examples/xray/fetcher/services"
	"github.com/goadesign/goa"
)

func TestFetchOK(t *testing.T) {
	const (
		archiveHref = "/archive/123"
	)
	var (
		service = goa.New("test service")
		ctx     = service.Context
	)
	cases := map[string]struct {
		Status int
		Body   string
	}{
		"ok":                {http.StatusOK, "body"},
		"ok-empty":          {http.StatusOK, ""},
		"bad_request":       {http.StatusBadRequest, "body"},
		"bad_request-empty": {http.StatusBadRequest, ""},
		"internal":          {http.StatusInternalServerError, "body"},
		"internal-empty":    {http.StatusInternalServerError, ""},
	}

	for k, c := range cases {
		var (
			s  *httptest.Server
			ac = services.NewTestArchiver()
		)

		// Setup expectations
		s = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(c.Status)
			fmt.Fprint(w, c.Body)
		}))
		ac.Expect("Archive", func(ctx context.Context, status int, body string) (string, error) {
			if status != c.Status {
				t.Errorf("%s: incorrect archived status %d, expected %d", k, status, c.Status)
			}
			if body != c.Body {
				t.Errorf("%s: incorrect archived body %#v, expected %#v", k, body, c.Body)
			}
			return archiveHref, nil
		})

		// Run tested code
		ctrl := NewFetcherController(service, ac)
		_, m := test.FetchFetcherOK(t, ctx, service, ctrl, s.URL)

		// Assert on results
		if err := ac.MetExpectations(); err != nil {
			t.Errorf("%s: %s", k, err)
		}
		if m.ArchiveHref != archiveHref {
			t.Errorf("%s: invalid media archive href %#v, expected %#v", k, m.ArchiveHref, archiveHref)
		}
		if m.Status != c.Status {
			t.Errorf("%s: invalid media archive status %d, expected %d", k, m.Status, c.Status)
		}

		// Cleanup
		if s != nil {
			s.Close()
		}
	}
}

func TestFetchInternalError(t *testing.T) {
	var (
		service      = goa.New("test service")
		ctx          = service.Context
		archiveError = fmt.Errorf("expected")
	)
	cases := map[string]struct {
		Status int
		Body   string
	}{
		"ok":                {http.StatusOK, "body"},
		"ok-empty":          {http.StatusOK, ""},
		"bad_request":       {http.StatusBadRequest, "body"},
		"bad_request-empty": {http.StatusBadRequest, ""},
		"internal":          {http.StatusInternalServerError, "body"},
		"internal-empty":    {http.StatusInternalServerError, ""},
	}

	for k, c := range cases {
		var (
			s  *httptest.Server
			ac = services.NewTestArchiver()
		)

		// Setup expectations
		s = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(c.Status)
			fmt.Fprint(w, c.Body)
		}))
		ac.Expect("Archive", func(ctx context.Context, status int, body string) (string, error) {
			if status != c.Status {
				t.Errorf("%s: incorrect archived status %d, expected %d", k, status, c.Status)
			}
			if body != c.Body {
				t.Errorf("%s: incorrect archived body %#v, expected %#v", k, body, c.Body)
			}
			return "", archiveError
		})

		// Run tested code
		ctrl := NewFetcherController(service, ac)
		test.FetchFetcherInternalServerError(t, ctx, service, ctrl, s.URL)

		// Assert on results
		if err := ac.MetExpectations(); err != nil {
			t.Errorf("%s: %s", k, err)
		}

		// Cleanup
		if s != nil {
			s.Close()
		}
	}
}

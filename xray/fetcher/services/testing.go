package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	// TestArchiver is a mock CRS client which is NOT goroutine safe.
	TestArchiver struct {
		// Mutex used to synchronize access to Stopped.
		*sync.Mutex
		expected   *expectations
		unexpected []string
	}

	// expectations is the data structure used to record expected function
	// calls and the corresponding behavior.
	expectations struct {
		*sync.Mutex
		Entries map[string][]interface{}
	}
)

// NewTestArchiver returns a mock archiver service client that can be used by
// tests to set and check expectations.
func NewTestArchiver() *TestArchiver {
	return &TestArchiver{
		Mutex: &sync.Mutex{},
		expected: &expectations{
			Mutex:   &sync.Mutex{},
			Entries: make(map[string][]interface{}),
		},
	}
}

// Expect records the request handler in the list of expected request calls.
func (c *TestArchiver) Expect(fn string, e interface{}) {
	c.expected.Lock()
	defer c.expected.Unlock()
	c.expected.Entries[fn] = append(c.expected.Entries[fn], e)
}

// Expectation removes the expectation for the function with the given name from
// the expected calls if there is one and returns it. It returns nil if there is
// no (more) expectations for the function.
func (c *TestArchiver) Expectation(fn string) interface{} {
	c.expected.Lock()
	defer c.expected.Unlock()
	es, ok := c.expected.Entries[fn]
	if !ok {
		c.unexpected = append(c.unexpected, fn)
		return nil
	}
	e := es[0]
	if len(es) == 1 {
		delete(c.expected.Entries, fn)
	} else {
		c.expected.Entries[fn] = c.expected.Entries[fn][1:]
	}
	return e
}

// MetExpectations returns nil if there no expectation left to be called and if
// there is no call that was made that did not match an expectation. It returns
// an error describing what is left to be called or what was called with no
// expectation otherwise.
func (c *TestArchiver) MetExpectations() error {
	c.expected.Lock()
	defer c.expected.Unlock()
	var msg string
	if len(c.unexpected) > 0 {
		msg = fmt.Sprintf("%s was called but wasn't expected.", strings.Join(c.unexpected, ", "))
	}
	if len(c.expected.Entries) > 0 {
		if len(msg) > 0 {
			msg += "\n"
		}
		i := 0
		keys := make([]string, len(c.expected.Entries))
		for e := range c.expected.Entries {
			keys[i] = e
			i++
		}
		msg += fmt.Sprintf("%s was expected to be called but wasn't.", strings.Join(keys, ", "))
	}
	if msg == "" {
		return nil
	}
	return errors.New(msg)
}

// Archive runs the "Archive" expectations
func (c *TestArchiver) Archive(ctx context.Context, status int, body string) (string, error) {
	if e := c.Expectation("Archive"); e != nil {
		return e.(func(context.Context, int, string) (string, error))(ctx, status, body)
	}
	return "", nil
}

// Healthy runs the "Healthy" expectations
func (c *TestArchiver) Healthy(ctx context.Context) (bool, error) {
	if e := c.Expectation("Healthy"); e != nil {
		return e.(func(context.Context) (bool, error))(ctx)
	}
	return false, nil
}

// WaitUntilHealthy runs the "WaitUntilHealthy" expectations
func (c *TestArchiver) WaitUntilHealthy(ctx context.Context, t time.Duration) (bool, error) {
	if e := c.Expectation("WaitUntilHealthy"); e != nil {
		return e.(func(context.Context, time.Duration) (bool, error))(ctx, t)
	}
	return false, nil
}

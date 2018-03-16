package main

import (
	"io/ioutil"
	"net/http"

	"github.com/goadesign/examples/xray/fetcher/app"
	"github.com/goadesign/examples/xray/fetcher/services"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/client"
	"github.com/goadesign/goa/middleware/xray"
)

// FetcherController implements the fetcher resource.
type FetcherController struct {
	*goa.Controller
	Archiver services.Archiver
}

// NewFetcherController creates a fetcher controller.
func NewFetcherController(service *goa.Service, a services.Archiver) *FetcherController {
	return &FetcherController{
		Controller: service.NewController("FetcherController"),
		Archiver:   a,
	}
}

// Fetch runs the fetch action.
func (c *FetcherController) Fetch(ctx *app.FetchFetcherContext) error {
	// Create traced client
	cl := xray.WrapDoer(client.HTTPClientDoer(http.DefaultClient))

	// Make request to external endpoint
	req, err := http.NewRequest("GET", ctx.URL, nil)
	if err != nil {
		return goa.ErrBadRequest("failed to build request", "err", err)
	}
	resp, err := cl.Do(ctx, req)
	if err != nil {
		return goa.ErrBadRequest("failed to make request", "err", err)
	}

	// Read response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return goa.ErrBadRequest("failed to decode response", "err", err)
	}

	// Annotate trace with response status code
	s := xray.ContextSegment(ctx)
	s.AddAnnotation("remote-url", req.URL.String())
	s.AddInt64Annotation("remote-status", int64(resp.StatusCode))

	// Archive response using archiver service
	href, err := c.Archiver.Archive(ctx, resp.StatusCode, string(body))
	if err != nil {
		return goa.ErrInternal("failed to archive response", "err", err)
	}

	// Return response
	m := &app.FetchMedia{
		ArchiveHref: href,
		Status:      resp.StatusCode,
	}
	return ctx.OK(m)
}

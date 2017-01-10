package main

import (
	"sync"

	"github.com/goadesign/examples/xray/archiver/app"
	"github.com/goadesign/goa"
)

type (
	// ArchiverController implements the archiver resource.
	ArchiverController struct {
		*goa.Controller
		db *Archive
	}

	// Document represents a single archive document.
	Document struct {
		// Unique ID
		ID int
		// Status is the archived response HTTP status
		Status int
		// Body is the archive response HTTP body
		Body string
	}

	// Archive is the archived documents in-memory "DB"
	Archive struct {
		*sync.RWMutex
		docs []*Document
	}
)

// NewArchiverController creates a archiver controller.
func NewArchiverController(service *goa.Service) *ArchiverController {
	return &ArchiverController{
		Controller: service.NewController("ArchiverController"),
		db:         &Archive{RWMutex: &sync.RWMutex{}},
	}
}

// Archive runs the archive action.
func (c *ArchiverController) Archive(ctx *app.ArchiveArchiverContext) error {
	doc := &Document{Status: ctx.Payload.Status, Body: ctx.Payload.Body}
	c.db.Store(doc)
	return ctx.OK(archiveMedia(doc))
}

// Read runs the read action.
func (c *ArchiverController) Read(ctx *app.ReadArchiverContext) error {
	doc := c.db.Read(ctx.ID)
	if doc == nil {
		return ctx.NotFound()
	}
	return ctx.OK(archiveMedia(doc))
}

// Store adds an archived document to the archive. Store takes care of
// initializing the document ID.
func (a *Archive) Store(doc *Document) {
	a.Lock()
	defer a.Unlock()
	id := len(a.docs) + 1
	doc.ID = id
	a.docs = append(a.docs, doc)
}

// Read retrieves an archived document by ID. It returns nil if there isn't one.
func (a *Archive) Read(id int) *Document {
	a.RLock()
	defer a.RUnlock()
	if id > len(a.docs) {
		return nil
	}
	return a.docs[id-1]
}

// archiveMedia converts a Document into a app.ArchiveMedia
func archiveMedia(doc *Document) *app.ArchiveMedia {
	return &app.ArchiveMedia{
		Href:   app.ArchiverHref(doc.ID),
		Status: doc.Status,
		Body:   doc.Body,
	}
}

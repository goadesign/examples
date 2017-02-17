package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/goadesign/examples/upload/app"
	"github.com/goadesign/goa"
)

// ImageData represents an uploaded image metadata.
// This would probably be persisted to a database in a real application.
type ImageData struct {
	// Unique image id
	ID int
	// Filename of image
	Filename string
	// UploadedAt is the upload timestamp
	UploadedAt time.Time
}

// ImageController implements the image resource.
type ImageController struct {
	*goa.Controller
	*sync.Mutex
	images []*ImageData
}

// NewImageController creates a image controller.
func NewImageController(service *goa.Service) *ImageController {
	return &ImageController{
		Controller: service.NewController("ImageController"),
		Mutex:      &sync.Mutex{},
	}
}

// Show runs the show action.
func (c *ImageController) Show(ctx *app.ShowImageContext) error {
	data := c.loadImage(ctx.ID)
	if data == nil {
		return ctx.NotFound()
	}
	return ctx.OK(&app.ImageMedia{ID: data.ID, Filename: data.Filename, UploadedAt: data.UploadedAt})
}

// Upload runs the upload action.
func (c *ImageController) Upload(ctx *app.UploadImageContext) error {
	// Assumes the image is under multipart section named "file"
	reader, err := ctx.MultipartReader()
	if err != nil {
		return goa.ErrBadRequest("failed to load multipart request: %s", err)
	}
	if reader == nil {
		return goa.ErrBadRequest("not a multipart request")
	}
	var images []*app.ImageMedia
	for {
		p, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return goa.ErrBadRequest("failed to load part: %s", err)
		}
		f, err := os.OpenFile("./images/"+p.FileName(), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return fmt.Errorf("failed to save file: %s", err) // causes a 500 response
		}
		defer f.Close()
		io.Copy(f, p)
		data := c.saveImage(p.FileName())
		images = append(images, &app.ImageMedia{ID: data.ID, Filename: data.Filename, UploadedAt: data.UploadedAt})
	}
	return ctx.OK(images)
}

// loadImage looks for the image with the given id.
// It returns nil if there isn't one.
func (c *ImageController) loadImage(id int) *ImageData {
	c.Lock()
	defer c.Unlock()
	id--
	if id >= 0 && id < len(c.images) {
		return c.images[id]
	}
	return nil
}

// saveImage builds the image data and "saves" it.
func (c *ImageController) saveImage(filename string) *ImageData {
	c.Lock()
	defer c.Unlock()
	data := &ImageData{
		ID:         len(c.images) + 1,
		Filename:   filename,
		UploadedAt: time.Now(),
	}
	c.images = append(c.images, data)
	return data
}

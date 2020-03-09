package uploaddownload

import (
	"context"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"

	updown "goa.design/examples/upload_download/gen/updown"
)

// updown service example implementation.
// The example methods log the requests and return zero values.
type updownsrvc struct {
	dir    string // Path to download and upload directory
	logger *log.Logger
}

// NewUpdown returns the updown service implementation.
func NewUpdown(dir string, logger *log.Logger) (updown.Service, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(abs); err != nil {
		return nil, err
	}

	return &updownsrvc{abs, logger}, nil
}

// Upload implements upload.
func (s *updownsrvc) Upload(ctx context.Context, p *updown.UploadPayload, req io.ReadCloser) error {
	// Don't forget to close the body reader!
	defer req.Close()

	// Make sure upload directory exists
	uploadDir := filepath.Join(s.dir, p.Dir)
	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		return updown.MakeInternalError(err)
	}

	// Createa multipart request reader
	_, params, err := mime.ParseMediaType(p.ContentType)
	if err != nil {
		return updown.MakeInvalidMediaType(err)
	}
	mr := multipart.NewReader(req, params["boundary"])

	// Go through each part and save the corresponding content to disk.
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			// We're done!
			return nil
		}
		if err != nil {
			return updown.MakeInvalidMultipartRequest(err)
		}

		// Create uploaded file, potentially overridding existing file.
		fpath := filepath.Join(uploadDir, part.FileName())
		f, err := os.Create(fpath)
		if err != nil {
			return updown.MakeInternalError(err)
		}
		defer f.Close()

		// Stream content to disk.
		n, err := io.Copy(f, part)
		if err != nil {
			return updown.MakeInternalError(err)
		}
		s.logger.Printf("Written %d bytes to %q", n, fpath)
	}
}

// Download implements download.
func (s *updownsrvc) Download(ctx context.Context, p string) (*updown.DownloadResult, io.ReadCloser, error) {
	// Locate file to download.
	abs := filepath.Join(s.dir, p)
	fi, err := os.Stat(abs)
	if err != nil {
		return nil, nil, updown.MakeInvalidFilePath(err)
	}

	// Open file for streaming
	f, err := os.Open(abs)
	if err != nil {
		return nil, nil, updown.MakeInternalError(err)
	}

	return &updown.DownloadResult{Length: fi.Size()}, f, nil
}

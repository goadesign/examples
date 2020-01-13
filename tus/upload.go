package tusupload

import (
	"context"
	"fmt"
	"log"

	upload "goa.design/examples/tus/gen/upload"
)

// upload service example implementation.
// The example methods log the requests and return zero values.
type uploadsrvc struct {
	uploadDir string
	maxSize   uint
	logger    *log.Logger
}

// NewUpload returns the upload service implementation.
func NewUpload(uploadDir string, maxSize uint, logger *log.Logger) upload.Service {
	return &uploadsrvc{uploadDir, maxSize, logger}
}

// Clients use the HEAD request to determine the offset at which the upload
// should be continued.
func (s *uploadsrvc) Head(ctx context.Context, p *upload.HeadPayload) (*upload.HeadResult, error) {
	up, ok := activeUploads[p.ID]
	if !ok {
		return nil, upload.MakeNotFound(fmt.Errorf("no ongoing upload with id %q", p.ID))
	}
	return &upload.HeadResult{
		UploadOffset:      up.offset,
		UploadLength:      up.length,
		UploadDeferLength: up.deferLength,
		UploadMetadata:    up.metadata,
	}, nil
}

// Clients use the PATCH method to start or resume an upload.
func (s *uploadsrvc) Patch(ctx context.Context, p *upload.PatchPayload) (res *upload.PatchResult, err error) {
	up, ok := activeUploads[p.ID]
	if !ok {
		return nil, upload.MakeNotFound(fmt.Errorf("no ongoing upload with id %q", p.ID))
	}
	res = &upload.PatchResult{}
	s.logger.Print("upload.patch")
	return
}

// Clients use the OPTIONS method to gather information about the Server’s
// current configuration.
func (s *uploadsrvc) Options(ctx context.Context) (res *upload.OptionsResult, err error) {
	res = &upload.OptionsResult{
		TusVersion:           []string{"1.0.0"},
		TusExtension:         "creation,creation-with-upload,creation-defer-length,expiration,checksum,termination",
		TusChecksumAlgorithm: "sha1,md5,crc32",
	}
	if s.maxSize > 0 {
		res.TusMaxSize = &s.maxSize
	}
	return
}

// Clients use the POST method against a known upload creation URL to request a
// new upload resource.
func (s *uploadsrvc) Post(ctx context.Context, p *upload.PostPayload) (res *upload.PostResult, err error) {
	res = &upload.PostResult{}
	s.logger.Print("upload.post")
	return
}

// Clients use the DELETE method to terminate completed and unfinished uploads
// allowing the Server to free up used resources.
func (s *uploadsrvc) Delete(ctx context.Context, p *upload.DeletePayload) (res *upload.DeleteResult, err error) {
	res = &upload.DeleteResult{}
	s.logger.Print("upload.delete")
	return
}

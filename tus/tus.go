package tus

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rs/xid"
	"goa.design/examples/tus/gen/http/tus/server"
	"goa.design/examples/tus/gen/tus"
	"goa.design/examples/tus/persist"
	"goa.design/examples/tus/upload"
	goahttp "goa.design/goa/v3/http"
)

// upload service
type tussvc struct {
	store         persist.Store
	newWriter     func(id string, len *int64) (io.WriteCloser, error)
	maxSize       int64
	uploadTimeout time.Duration
	logger        *log.Logger
}

// New creates a TUS service that accepts uploads up to maxSize bytes and that
// expire after uploadTimeout. If maxSize is 0 then there is no limit on the
// size of uploads. If uploadTimeout is 0 then uploads never expire. newWriter
// is called for each new tus. The unique upload identifier is given as argument
// to the function as well as the length of the upload if known (nil otherwise).
// The function returns the writer used to write the uploaded bytes. It must be
// possible to call Write concurrently on two writers returned by different
// invocation of the function.
func New(store persist.Store, newWriter func(string, *int64) (io.WriteCloser, error), maxSize int64, uploadTimeout time.Duration, logger *log.Logger) tus.Service {
	return HandleTUSResumable(&tussvc{
		store:         store,
		newWriter:     newWriter,
		maxSize:       maxSize,
		uploadTimeout: uploadTimeout,
		logger:        logger,
	})
}

// Clients use the HEAD request to determine the offset at which the upload
// should be continued.
func (s *tussvc) Head(ctx context.Context, p *tus.HeadPayload) (*tus.HeadResult, error) {
	up := s.getUpload(p.ID)

	// If the resource is not found, the Server SHOULD return either the 404 Not
	// Found, 410 Gone or 403 Forbidden status without the Upload-Offset header.
	if up == nil {
		return nil, tus.MakeNotFound(fmt.Errorf("no ongoing upload with id %q", p.ID))
	}

	pg := up.Progress()
	var df *int
	if pg.Length == 0 {
		v := 1
		df = &v
	}
	return &tus.HeadResult{
		UploadOffset:      pg.Offset,
		UploadLength:      &pg.Length,
		UploadDeferLength: df,
		UploadMetadata:    &pg.Metadata,
	}, nil
}

// Clients use the PATCH method to start or resume an tus.
func (s *tussvc) Patch(ctx context.Context, p *tus.PatchPayload, body io.ReadCloser) (res *tus.PatchResult, err error) {
	// All PATCH requests MUST use Content-Type:
	// application/offset+octet-stream, otherwise the server SHOULD return a 415
	// Unsupported Media Type status.
	ct := ctx.Value(goahttp.ContentTypeKey).(string)
	if ct != "application/offset+octet-stream" {
		return nil, tus.MakeInvalidContentType(fmt.Errorf("Content-Type header must be application/offset+octet-stream"))
	}

	up, err := s.store.Load(p.ID)
	if err != nil {
		return nil, tus.MakeInternal(err)
	}

	// If the server receives a PATCH request against a non-existent resource it
	// SHOULD return a 404 Not Found status.
	if up == nil {
		return nil, tus.MakeNotFound(fmt.Errorf("no ongoing upload with id %q or upload expired", p.ID))
	}

	offset, err := s.write(up, p.UploadOffset, body, p.UploadChecksum)
	if err != nil {
		return nil, err
	}

	res = &tus.PatchResult{UploadOffset: offset}
	if s.uploadTimeout > 0 {
		expires := up.Expiry().Format(http.TimeFormat)
		res.UploadExpires = &expires
	}
	return
}

// Clients use the OPTIONS method to gather information about the Server’s
// current configuration.
func (s *tussvc) Options(ctx context.Context) (res *tus.OptionsResult, err error) {
	res = &tus.OptionsResult{
		TusVersion:           TusResumable,
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
func (s *tussvc) Post(ctx context.Context, p *tus.PostPayload, body io.ReadCloser) (res *tus.PostResult, err error) {
	if p.UploadLength != nil && *p.UploadLength > s.maxSize {
		return nil, tus.MakeMaximumSizeExceeded(fmt.Errorf("upload length cannot exceed %v bytes", s.maxSize))
	}
	if p.UploadDeferLength != nil && *p.UploadDeferLength != 1 {
		return nil, tus.MakeInvalidDeferLength(fmt.Errorf("invalid Upload-Defer-Length header, got %v, expected 1", *p.UploadDeferLength))
	}
	if p.UploadLength == nil && p.UploadDeferLength == nil {
		return nil, tus.MakeMissingHeader(fmt.Errorf("missing Upload-Length or Upload-Defer-Length header"))
	}

	id := xid.New().String()
	w, err := s.newWriter(id, p.UploadLength)
	if err != nil {
		return nil, tus.MakeInternal(err)
	}
	up := upload.New(id, p.UploadLength, p.UploadMetadata, w, s.uploadTimeout)

	if err := s.store.Save(up); err != nil {
		return nil, tus.MakeInternal(err)
	}

	var offset int64
	offset, err = s.write(up, 0, body, p.UploadChecksum)
	if err != nil {
		return nil, err
	}
	if p.UploadLength != nil && offset >= *p.UploadLength {
		up.Complete()
	}

	res = &tus.PostResult{
		UploadOffset: offset,
		Location:     server.HeadTusPath(id),
	}
	if s.uploadTimeout > 0 {
		expires := up.Expiry().Format(http.TimeFormat)
		res.UploadExpires = &expires
	}
	return
}

// Clients use the DELETE method to terminate completed and unfinished uploads
// allowing the Server to free up used resources.
func (s *tussvc) Delete(ctx context.Context, p *tus.DeletePayload) (res *tus.DeleteResult, err error) {
	up, err := s.store.Load(p.ID)
	if err != nil {
		return nil, tus.MakeInternal(err)
	}

	// If the server receives a PATCH request against a non-existent resource it
	// SHOULD return a 404 Not Found status.
	if up == nil {
		return nil, tus.MakeNotFound(fmt.Errorf("no ongoing upload with id %q or upload expired", p.ID))
	}

	up.Cancel()

	return &tus.DeleteResult{}, nil
}

// write is a help method shared by POST and PATCH that takes care of mapping
// error from the upload package.
func (s *tussvc) write(up *upload.Upload, offset int64, content io.ReadCloser, checksum *string) (n int64, err error) {
	n, err = up.Write(offset, content, checksum)
	if err != nil {
		if _, ok := err.(upload.ErrBadChecksum); ok {
			return 0, tus.MakeChecksumMismatch(err)
		}
		if _, ok := err.(upload.ErrInvalidAlgo); ok {
			return 0, tus.MakeInvalidChecksumAlgorithm(err)
		}
		if _, ok := err.(upload.ErrInvalidOffset); ok {
			return 0, tus.MakeInvalidOffset(err)
		}
		return 0, tus.MakeInternal(fmt.Errorf("failed to write bytes: %s", err.Error()))
	}
	return
}

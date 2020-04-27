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
// is called for each upload. The unique upload identifier is given as argument
// to the function as well as the length of the upload if known (nil otherwise).
// The function returns the writer used to write the uploaded bytes. It must be
// possible to call Write concurrently on two writers returned by different
// invocation of the function. The given persist store is used to maintain state
// of ongoing uploads, it should persist the state so that uploads can be
// resumed across restarts.
func New(store persist.Store, newWriter func(string, *int64) (io.WriteCloser, error), maxSize int64, uploadTimeout time.Duration, logger *log.Logger) tus.Service {
	return handleTUSResumable(&tussvc{
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
	m, err := s.store.Load(p.ID)
	if err != nil {
		return nil, tus.MakeInternal(err)
	}

	// If the resource is not found, the Server SHOULD return either the 404 Not
	// Found, 410 Gone or 403 Forbidden status without the Upload-Offset header.
	if m == nil {
		return nil, tus.MakeNotFound(fmt.Errorf("no ongoing upload with id %q", p.ID))
	}

	var df *int
	if m.Length == nil {
		v := 1
		df = &v
	}
	return &tus.HeadResult{
		UploadOffset:      m.Offset,
		UploadLength:      m.Length,
		UploadDeferLength: df,
		UploadMetadata:    &m.Metadata,
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

	w, err := s.newWriter(p.ID, up.Length)
	if err != nil {
		return nil, tus.MakeInternal(err)
	}

	offset, err := Write(body, w, up, p.UploadOffset, p.UploadChecksum)
	if err != nil {
		return nil, err
	}
	if err := s.store.Save(p.ID, up); err != nil {
		return nil, tus.MakeInternal(err)
	}

	res = &tus.PatchResult{UploadOffset: offset}
	if !up.ExpiresAt.IsZero() {
		expires := up.ExpiresAt.Format(http.TimeFormat)
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
	startedAt := time.Now()
	var expiresAt time.Time
	if s.uploadTimeout > 0 {
		expiresAt = startedAt.Add(s.uploadTimeout)
	}
	var mdata string
	if p.UploadMetadata != nil {
		mdata = *p.UploadMetadata
	}
	up := persist.Upload{
		ID:        id,
		StartedAt: startedAt,
		ExpiresAt: expiresAt,
		Status:    persist.Started,
		Length:    p.UploadLength,
		Metadata:  mdata,
	}

	var offset int64
	offset, err = Write(body, w, &up, 0, p.UploadChecksum)
	if err != nil {
		return nil, err
	}
	if p.UploadLength != nil && offset >= *p.UploadLength {
		up.Status = persist.Completed
	}
	if err := s.store.Save(id, &up); err != nil {
		return nil, tus.MakeInternal(err)
	}

	res = &tus.PostResult{
		UploadOffset: offset,
		Location:     server.HeadTusPath(id),
	}
	if !expiresAt.IsZero() {
		expires := expiresAt.Format(http.TimeFormat)
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

	// For all future requests to this URL, the Server SHOULD respond with the
	// 404 Not Found or 410 Gone status.
	if up == nil {
		return nil, tus.MakeNotFound(fmt.Errorf("no ongoing upload with id %q or upload expired", p.ID))
	}
	if !up.Active() {
		return nil, tus.MakeGone(fmt.Errorf("upload with id %q %s", p.ID, up.Status.String()))
	}

	up.Status = persist.Completed
	if err := s.store.Save(p.ID, up); err != nil {
		return nil, tus.MakeInternal(err)
	}

	return &tus.DeleteResult{}, nil
}

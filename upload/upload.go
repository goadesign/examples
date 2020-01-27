package tusupload

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/rs/xid"
	"goa.design/examples/upload/gen/http/upload/server"
	genupload "goa.design/examples/upload/gen/upload"
	"goa.design/examples/upload/upload"
	goahttp "goa.design/goa/v3/http"
)

// upload service
type uploadsrvc struct {
	*sync.RWMutex
	activeUploads map[string]*upload.Upload

	newWriter         func(id string, len *uint) io.WriteCloser
	maxSize           uint
	uploadTimeout     time.Duration
	retentionDuration time.Duration
	logger            *log.Logger
}

// NewUploadSvc creates an upload service that accepts uploads up to maxSize bytes
// and that expire after uploadTimeout. If maxSize is 0 then there is no limit
// on the size of uploads. If uploadTimeout is 0 then uploads never expire.
// newWriter is called for each new upload. The unique upload identifier is
// given as argument to the function as well as the length of the upload if
// known (nil otherwise). The function returns the writer used to write the
// uploaded bytes. It must be possible to call Write concurrently on two writers
// returned by different invocation of the function.
func NewUploadSvc(newWriter func(string, *uint) io.WriteCloser, maxSize uint, uploadTimeout, retentionDuration time.Duration, logger *log.Logger) genupload.Service {
	return &uploadsrvc{
		RWMutex:           &sync.RWMutex{},
		activeUploads:     map[string]*upload.Upload{},
		newWriter:         newWriter,
		maxSize:           maxSize,
		uploadTimeout:     uploadTimeout,
		retentionDuration: retentionDuration,
		logger:            logger,
	}
}

// Clients use the HEAD request to determine the offset at which the upload
// should be continued.
func (s *uploadsrvc) Head(ctx context.Context, p *genupload.HeadPayload) (*genupload.HeadResult, error) {
	up := s.getUpload(p.ID)

	// If the resource is not found, the Server SHOULD return either the 404 Not
	// Found, 410 Gone or 403 Forbidden status without the Upload-Offset header.
	if up == nil {
		return nil, genupload.MakeNotFound(fmt.Errorf("no ongoing upload with id %q", p.ID))
	}

	pg := up.Progress()
	var df *int
	if pg.Length == 0 {
		v := 1
		df = &v
	}
	return &genupload.HeadResult{
		UploadOffset:      pg.Offset,
		UploadLength:      &pg.Length,
		UploadDeferLength: df,
		UploadMetadata:    &pg.Metadata,
	}, nil
}

// Clients use the PATCH method to start or resume an upload.
func (s *uploadsrvc) Patch(ctx context.Context, p *genupload.PatchPayload) (res *genupload.PatchResult, err error) {
	// All PATCH requests MUST use Content-Type:
	// application/offset+octet-stream, otherwise the server SHOULD return a 415
	// Unsupported Media Type status.
	ct := ctx.Value(goahttp.ContentTypeKey).(string)
	if ct != "application/offset+octet-stream" {
		return nil, genupload.MakeInvalidContentType(fmt.Errorf("Content-Type header must be application/offset+octet-stream"))
	}

	up := s.getUpload(p.ID)

	// If the server receives a PATCH request against a non-existent resource it
	// SHOULD return a 404 Not Found status.
	if up == nil {
		return nil, genupload.MakeNotFound(fmt.Errorf("no ongoing upload with id %q or upload expired", p.ID))
	}

	offset, err := s.write(up, p.UploadOffset, p.Content, p.UploadChecksum)
	if err != nil {
		return nil, err
	}

	res = &genupload.PatchResult{UploadOffset: offset}
	if s.uploadTimeout > 0 {
		expires := up.Expiry().Format(http.TimeFormat)
		res.UploadExpires = &expires
	}
	return
}

// Clients use the OPTIONS method to gather information about the Server’s
// current configuration.
func (s *uploadsrvc) Options(ctx context.Context) (res *genupload.OptionsResult, err error) {
	res = &genupload.OptionsResult{
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
func (s *uploadsrvc) Post(ctx context.Context, p *genupload.PostPayload) (res *genupload.PostResult, err error) {
	if p.UploadLength != nil && *p.UploadLength > s.maxSize || uint(len(p.Content)) > s.maxSize {
		return nil, genupload.MakeMaximumSizeExceeded(fmt.Errorf("upload length cannot exceed %v bytes", s.maxSize))
	}
	if p.UploadDeferLength != nil && *p.UploadDeferLength != 1 {
		return nil, genupload.MakeInvalidDeferLength(fmt.Errorf("invalid Upload-Defer-Length header, got %v, expected 1", *p.UploadDeferLength))
	}

	id := xid.New().String()
	up := upload.New(id, p.UploadLength, p.UploadMetadata, s.newWriter(id, p.UploadLength), s.uploadTimeout, s.onTerminate)

	s.Lock()
	s.activeUploads[id] = up
	s.Unlock()

	var offset uint
	if len(p.Content) > 0 {
		offset, err = s.write(up, 0, p.Content, p.UploadChecksum)
		if err != nil {
			return nil, err
		}
	}

	res = &genupload.PostResult{
		UploadOffset: offset,
		Location:     server.HeadUploadPath(id),
	}
	if s.uploadTimeout > 0 {
		expires := up.Expiry().Format(http.TimeFormat)
		res.UploadExpires = &expires
	}
	return
}

// Clients use the DELETE method to terminate completed and unfinished uploads
// allowing the Server to free up used resources.
func (s *uploadsrvc) Delete(ctx context.Context, p *genupload.DeletePayload) (res *genupload.DeleteResult, err error) {
	up := s.getUpload(p.ID)

	// If the server receives a PATCH request against a non-existent resource it
	// SHOULD return a 404 Not Found status.
	if up == nil {
		return nil, genupload.MakeNotFound(fmt.Errorf("no ongoing upload with id %q or upload expired", p.ID))
	}

	up.Cancel()
	res = &genupload.DeleteResult{}
	return
}

func (s *uploadsrvc) getUpload(id string) *upload.Upload {
	s.RLock()
	up, _ := s.activeUploads[id]
	s.RUnlock()
	return up
}

// write is a help method shared by POST and PATCH that takes care of mapping
// error from the upload package.
func (s *uploadsrvc) write(up *upload.Upload, offset uint, content []byte, checksum *string) (uint, error) {
	offset, err := up.Write(offset, content, checksum)
	if err != nil {
		if _, ok := err.(upload.ErrBadChecksum); ok {
			return 0, genupload.MakeChecksumMismatch(err)
		}
		if _, ok := err.(upload.ErrInvalidAlgo); ok {
			return 0, genupload.MakeInvalidChecksumAlgorithm(err)
		}
		if _, ok := err.(upload.ErrInvalidOffset); ok {
			return 0, genupload.MakeInvalidOffset(err)
		}
		return 0, genupload.MakeInternal(fmt.Errorf("failed to write bytes: %s", err.Error()))
	}
}

// onTerminate is called whenever an upload terminates, successfully or not.
func (s *uploadsrvc) onTerminate(id string, _ upload.State) {
	// Remove uploads after configured retention time.
	time.AfterFunc(s.retentionDuration, func() {
		s.Lock()
		delete(s.activeUploads, id)
		s.Unlock()
	})
}

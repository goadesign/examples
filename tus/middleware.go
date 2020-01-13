package tusupload

import (
	"context"

	upload "goa.design/examples/tus/gen/upload"
)

// uploadSvcWrapper wraps the given upload service instance and:
//
//   * checks the value of the request Tus-Resumable header.
//   * writes the response Tus-Resumable header.
type uploadSvcWrapper struct {
	svc *uploadsrvc
}

// TusResumable is the version of tus implemented by this service.
const TusResumable = "1.0.0"

// Wrap the given upload service instance.
func Wrap(s *uploadsrvc) upload.Service {
	w := uploadSvcWrapper{s}
	return &w
}

// Head checks and writes the Tus-Resumable header.
func (s *uploadSvcWrapper) Head(ctx context.Context, p *upload.HeadPayload) (res *upload.HeadResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, upload.MakeInvalidTusResumable(err)
	}
	res, err = s.svc.Head(ctx, p)
	res.TusResumable = TusResumable
	return
}

// Clients use the PATCH method to start or resume an upload.
func (s *uploadSvcWrapper) Patch(ctx context.Context, p *upload.PatchPayload) (res *upload.PatchResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, upload.MakeInvalidTusResumable(err)
	}
	res, err = s.svc.Patch(ctx, p)
	res.TusResumable = TusResumable
	return
}

// Clients use the OPTIONS method to gather information about the Server’s
// current configuration.
func (s *uploadSvcWrapper) Options(ctx context.Context) (res *upload.OptionsResult, err error) {
	res, err = s.svc.Options(ctx)
	res.TusResumable = TusResumable
	return
}

// Clients use the POST method against a known upload creation URL to request a
// new upload resource.
func (s *uploadSvcWrapper) Post(ctx context.Context, p *upload.PostPayload) (res *upload.PostResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, upload.MakeInvalidTusResumable(err)
	}
	res, err = s.svc.Post(ctx, p)
	res.TusResumable = TusResumable
	return
}

// Clients use the DELETE method to terminate completed and unfinished uploads
// allowing the Server to free up used resources.
func (s *uploadSvcWrapper) Delete(ctx context.Context, p *upload.DeletePayload) (res *upload.DeleteResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, upload.MakeInvalidTusResumable(err)
	}
	res, err = s.svc.Delete(ctx, p)
	res.TusResumable = TusResumable
	return
}

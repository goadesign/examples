package tus

import (
	"context"
	"io"

	"goa.design/examples/tus/gen/tus"
)

// tusWrapper wraps the given TUS service instance and:
//   - checks the value of the request Tus-Resumable header.
//   - writes the response Tus-Resumable header.
type tusWrapper struct {
	svc *tussvc
}

// TusResumable is the version of tus implemented by this service.
const TusResumable = "1.0.0"

// handleTUSResumable wraos the given upload service instance and adds
// Tus-Resumable request and response headers handling.
func handleTUSResumable(s *tussvc) tus.Service {
	w := tusWrapper{s}
	return &w
}

// Head checks and writes the Tus-Resumable header.
func (s *tusWrapper) Head(ctx context.Context, p *tus.HeadPayload) (res *tus.HeadResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, &tus.ErrInvalidTUSResumable{TusVersion: TusResumable}
	}
	res, err = s.svc.Head(ctx, p)
	if res != nil {
		res.TusResumable = TusResumable
	}
	return
}

// Clients use the PATCH method to start or resume an tus.
func (s *tusWrapper) Patch(ctx context.Context, p *tus.PatchPayload, r io.ReadCloser) (res *tus.PatchResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, &tus.ErrInvalidTUSResumable{TusVersion: TusResumable}
	}
	res, err = s.svc.Patch(ctx, p, r)
	if res != nil {
		res.TusResumable = TusResumable
	}
	return
}

// Clients use the OPTIONS method to gather information about the Server’s
// current configuration.
func (s *tusWrapper) Options(ctx context.Context) (res *tus.OptionsResult, err error) {
	res, err = s.svc.Options(ctx)
	if res != nil {
		res.TusResumable = TusResumable
	}
	return
}

// Clients use the POST method against a known upload creation URL to request a
// new upload resource.
func (s *tusWrapper) Post(ctx context.Context, p *tus.PostPayload, r io.ReadCloser) (res *tus.PostResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, &tus.ErrInvalidTUSResumable{TusVersion: TusResumable}
	}
	res, err = s.svc.Post(ctx, p, r)
	if res != nil {
		res.TusResumable = TusResumable
	}
	return
}

// Clients use the DELETE method to terminate completed and unfinished uploads
// allowing the Server to free up used resources.
func (s *tusWrapper) Delete(ctx context.Context, p *tus.DeletePayload) (res *tus.DeleteResult, err error) {
	if p.TusResumable != TusResumable {
		return nil, &tus.ErrInvalidTUSResumable{TusVersion: TusResumable}
	}
	res, err = s.svc.Delete(ctx, p)
	if res != nil {
		res.TusResumable = TusResumable
	}
	return
}

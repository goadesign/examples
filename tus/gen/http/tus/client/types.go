// Code generated by goa v3.21.5, DO NOT EDIT.
//
// tus HTTP client types
//
// Command:
// $ goa gen goa.design/examples/tus/design

package client

import (
	tus "goa.design/examples/tus/gen/tus"
	goa "goa.design/goa/v3/pkg"
)

// NewHeadResultOK builds a "tus" service "head" endpoint result from a HTTP
// "OK" response.
func NewHeadResultOK(tusResumable string, uploadOffset int64, uploadLength *int64, uploadDeferLength *int, uploadMetadata *string) *tus.HeadResult {
	v := &tus.HeadResult{}
	v.TusResumable = tusResumable
	v.UploadOffset = uploadOffset
	v.UploadLength = uploadLength
	v.UploadDeferLength = uploadDeferLength
	v.UploadMetadata = uploadMetadata

	return v
}

// NewHeadNotFound builds a tus service head endpoint NotFound error.
func NewHeadNotFound(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewHeadGone builds a tus service head endpoint Gone error.
func NewHeadGone(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewHeadInvalidTusResumable builds a tus service head endpoint
// InvalidTusResumable error.
func NewHeadInvalidTusResumable(tusVersion string) *tus.ErrInvalidTUSResumable {
	v := &tus.ErrInvalidTUSResumable{}
	v.TusVersion = tusVersion

	return v
}

// NewPatchResultNoContent builds a "tus" service "patch" endpoint result from
// a HTTP "NoContent" response.
func NewPatchResultNoContent(tusResumable string, uploadOffset int64, uploadExpires *string) *tus.PatchResult {
	v := &tus.PatchResult{}
	v.TusResumable = tusResumable
	v.UploadOffset = uploadOffset
	v.UploadExpires = uploadExpires

	return v
}

// NewPatchInvalidContentType builds a tus service patch endpoint
// InvalidContentType error.
func NewPatchInvalidContentType(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPatchInvalidOffset builds a tus service patch endpoint InvalidOffset
// error.
func NewPatchInvalidOffset(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPatchNotFound builds a tus service patch endpoint NotFound error.
func NewPatchNotFound(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPatchGone builds a tus service patch endpoint Gone error.
func NewPatchGone(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPatchInvalidChecksumAlgorithm builds a tus service patch endpoint
// InvalidChecksumAlgorithm error.
func NewPatchInvalidChecksumAlgorithm(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPatchChecksumMismatch builds a tus service patch endpoint
// ChecksumMismatch error.
func NewPatchChecksumMismatch(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPatchInternal builds a tus service patch endpoint Internal error.
func NewPatchInternal(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPatchInvalidTusResumable builds a tus service patch endpoint
// InvalidTusResumable error.
func NewPatchInvalidTusResumable(tusVersion string) *tus.ErrInvalidTUSResumable {
	v := &tus.ErrInvalidTUSResumable{}
	v.TusVersion = tusVersion

	return v
}

// NewOptionsResultNoContent builds a "tus" service "options" endpoint result
// from a HTTP "NoContent" response.
func NewOptionsResultNoContent(tusResumable string, tusVersion string, tusExtension string, tusMaxSize *int64, tusChecksumAlgorithm string) *tus.OptionsResult {
	v := &tus.OptionsResult{}
	v.TusResumable = tusResumable
	v.TusVersion = tusVersion
	v.TusExtension = tusExtension
	v.TusMaxSize = tusMaxSize
	v.TusChecksumAlgorithm = tusChecksumAlgorithm

	return v
}

// NewOptionsInvalidTusResumable builds a tus service options endpoint
// InvalidTusResumable error.
func NewOptionsInvalidTusResumable(tusVersion string) *tus.ErrInvalidTUSResumable {
	v := &tus.ErrInvalidTUSResumable{}
	v.TusVersion = tusVersion

	return v
}

// NewPostResultCreated builds a "tus" service "post" endpoint result from a
// HTTP "Created" response.
func NewPostResultCreated(location string, tusResumable string, uploadOffset int64, uploadExpires *string) *tus.PostResult {
	v := &tus.PostResult{}
	v.Location = location
	v.TusResumable = tusResumable
	v.UploadOffset = uploadOffset
	v.UploadExpires = uploadExpires

	return v
}

// NewPostMissingHeader builds a tus service post endpoint MissingHeader error.
func NewPostMissingHeader(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPostInvalidDeferLength builds a tus service post endpoint
// InvalidDeferLength error.
func NewPostInvalidDeferLength(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPostInvalidChecksumAlgorithm builds a tus service post endpoint
// InvalidChecksumAlgorithm error.
func NewPostInvalidChecksumAlgorithm(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPostMaximumSizeExceeded builds a tus service post endpoint
// MaximumSizeExceeded error.
func NewPostMaximumSizeExceeded(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPostChecksumMismatch builds a tus service post endpoint ChecksumMismatch
// error.
func NewPostChecksumMismatch(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewPostInvalidTusResumable builds a tus service post endpoint
// InvalidTusResumable error.
func NewPostInvalidTusResumable(tusVersion string) *tus.ErrInvalidTUSResumable {
	v := &tus.ErrInvalidTUSResumable{}
	v.TusVersion = tusVersion

	return v
}

// NewDeleteResultNoContent builds a "tus" service "delete" endpoint result
// from a HTTP "NoContent" response.
func NewDeleteResultNoContent(tusResumable string) *tus.DeleteResult {
	v := &tus.DeleteResult{}
	v.TusResumable = tusResumable

	return v
}

// NewDeleteNotFound builds a tus service delete endpoint NotFound error.
func NewDeleteNotFound(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewDeleteGone builds a tus service delete endpoint Gone error.
func NewDeleteGone(name string, id string, message string, temporary bool, timeout bool, fault bool) *goa.ServiceError {
	v := &goa.ServiceError{}
	v.Name = name
	v.ID = id
	v.Message = message
	v.Temporary = temporary
	v.Timeout = timeout
	v.Fault = fault

	return v
}

// NewDeleteInvalidTusResumable builds a tus service delete endpoint
// InvalidTusResumable error.
func NewDeleteInvalidTusResumable(tusVersion string) *tus.ErrInvalidTUSResumable {
	v := &tus.ErrInvalidTUSResumable{}
	v.TusVersion = tusVersion

	return v
}

package tusupload

// upload represents a single file upload.
type fileUpload struct {
	id          string
	offset      uint
	length      *uint
	deferLength *int
	metadata    *string
}

// activeUploads keeps track of all on-going uploads.
var activeUploads map[string]*fileUpload

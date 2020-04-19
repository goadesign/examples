package persist

import "goa.design/examples/tus/upload"

// Store is an opinionated data store interface suitable for use together with a
// TUS server implementation.
type Store interface {
	// Load retrieves the upload with the given ID. Load should not return an
	// error if there is not upload with the given ID.
	Load(id string) (*upload.Upload, error)

	// Save persists the given upload.
	Save(up *upload.Upload) error
}

package persist

// Store is an opinionated data store interface suitable for use together with a
// TUS server implementation.
type Store interface {
	// Load retrieves the upload with the given ID. Load should return nil, nil
	// if there is not upload with the given ID.
	Load(id string) (*Upload, error)

	// Save persists the given upload with the given ID.
	Save(id string, s *Upload) error
}

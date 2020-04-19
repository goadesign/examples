package persist

import (
	"sync"

	"goa.design/examples/tus/upload"
)

// memory map is a store implementation that keeps ongoing upload information in
// memory. This store can only be used when a single TUS server is used and does
// not tolerate process restarts. It is thus not suitable for production uses.
type memoryMap struct {
	m       *sync.RWMutex
	uploads map[string]*upload.Upload
}

// Load gets the upload from the map.
func (m memoryMap) Load(id string) (*upload.Upload, error) {
	m.m.RLock()
	defer m.m.RUnlock()
	return m.uploads[id], nil
}

// Save stores the upload in the map.
func (m memoryMap) Save(id string, up *upload.Upload) error {
	m.m.Lock()
	defer m.m.Unlock()
	m.uploads[id] = up
	return nil
}

package persist

import (
	"sync"
)

// memory map is a store implementation that keeps ongoing upload information in
// memory. This store can only be used when a single TUS server is used and does
// not tolerate process restarts. It is thus not suitable for production uses.
type memoryMap struct {
	m       *sync.RWMutex
	uploads map[string]*Upload
}

// NewInMemory returns a store implementation that relies on a in-memory map to
// keep the server state. This means only one process can be exposed to the
// upload HTTP endpoint and that uploads won't survive process restarts. This
// implementation is thus probably not suitable for prodcution use.
func NewInMemory() Store {
	var m sync.RWMutex
	uploads := make(map[string]*Upload)
	return &memoryMap{&m, uploads}
}

// Load gets the upload from the map.
func (m *memoryMap) Load(id string) (*Upload, error) {
	m.m.RLock()
	defer m.m.RUnlock()

	return m.uploads[id], nil
}

// Save stores the upload in the map.
func (m *memoryMap) Save(id string, up *Upload) error {
	m.m.Lock()
	defer m.m.Unlock()

	m.uploads[id] = up
	return nil
}

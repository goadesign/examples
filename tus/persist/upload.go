package persist

import (
	"time"
)

type (
	// Upload contains the information needed to manage an upload.
	Upload struct {
		// ID is a unique identifier for the upload.
		ID string
		// StartedAt is the creation timestamp for the upload.
		StartedAt time.Time
		// ExpiresAt is the upload expiration timestamp if any, the Zero
		// time.Time value otherwise.
		ExpiresAt time.Time
		// Status is the current upload status.
		Status Status
		// Length is the total length of the upload if known, nil otherwise.
		Length *int64
		// Offset is the count of bytes that have already been uploaded.
		Offset int64
		// Metadata is arbitrary pass-through data that can be used by clients to
		// attach information to the upload.
		Metadata string
	}

	// Status indicates the status of an upload.
	Status int
)

const (
	Started Status = iota
	Completed
	TimedOut
	Failed
)

// Active returns true if the upload has not completed, failed or timed out.
func (u *Upload) Active() bool {
	return u.Status == Started
}

var names = [...]string{
	"started",
	"completed",
	"timed out",
	"failed",
}

// String returns a human-friendly description of the upload state.
func (s Status) String() string {
	return names[s]
}

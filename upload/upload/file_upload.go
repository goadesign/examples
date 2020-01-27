package upload

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"strings"
	"sync"
	"time"
)

type (
	// Upload represents a single file upload.
	Upload struct {
		m         *sync.RWMutex
		id        string
		startedAt time.Time
		expiresAt time.Time
		metadata  *string
		state     State
		length    *uint
		offset    uint
		writer    io.Writer
		cb        func(string, State)
		timer     *time.Timer
		cancel    chan struct{}
	}

	// Info contains read-only information for a given upload.
	Info struct {
		// ID is the upload ID.
		ID string

		// State is the state of the upload.
		State State

		// Length is the total upload length in bytes, 0 if unknown.
		Length uint

		// Offset is the total number of bytes already uploaded.
		Offset uint

		// Metadata is an arbitrary value given when the upload is initially
		// created, empty string if none was provided.
		Metadata string

		// StartedAt is the upload creation timestamp.
		StartedAt time.Time
	}

	// ErrInvalidAlgo is the type of the error returned by Write when the
	// provided checksum algorithm is not supported.
	ErrInvalidAlgo string

	// ErrBadChecksum is the type of the error returned by Write when the
	// provided and computed checksums don't match.
	ErrBadChecksum string

	// ErrInvalidOffset is the type of the error returned by Write when the
	// provided offset does not match the upload internal state.
	ErrInvalidOffset string
)

func (e ErrInvalidAlgo) Error() string   { return string(e) }
func (e ErrBadChecksum) Error() string   { return string(e) }
func (e ErrInvalidOffset) Error() string { return string(e) }

// New instantiates a new upload with the given id that writes to the given
// writer. length indicates the expected total number of bytes if known. If ttl
// is set then a timer is created that changes the state of the upload from
// Started to TimedOut upon expiration. cb - if not nil - is called once when
// the state of the upload changes from Started to another value. It is given
// the ID of the upload and the new state as arguments.
func New(id string, length *uint, metadata *string, writer io.Writer, ttl time.Duration, cb func(string, State)) *Upload {
	up := Upload{
		id:        id,
		metadata:  metadata,
		startedAt: time.Now(),
		length:    length,
		writer:    writer,
		cb:        cb,
		cancel:    make(chan struct{}),
		m:         &sync.RWMutex{},
	}

	if ttl > 0 {
		up.expiresAt = up.startedAt.Add(ttl)
		up.timer = time.AfterFunc(ttl, func() {
			up.m.Lock()
			defer up.m.Unlock()

			up.terminate(TimedOut)
		})
	}

	return &up
}

// Progress returns information on the given upload.
func (up *Upload) Progress() Info {
	up.m.RLock()
	defer up.m.RUnlock()

	i := Info{
		ID:        up.id,
		State:     up.state,
		Offset:    up.offset,
		StartedAt: up.startedAt,
	}

	if up.length != nil {
		i.Length = *up.length
	}

	if up.metadata != nil {
		i.Metadata = *up.metadata
	}

	return i
}

// Expiry retrurns the upload expiration timestamp.
func (up *Upload) Expiry() time.Time {
	return up.expiresAt
}

// Write validates the checksum if not nil then writes the data to the upload
// writer and returns the updated upload offset value. checksum must start with
// "sha1 ", "md5 " or "crc32 " and be followed by the Base64 encoded checksum.
// It is an error to write to an upload whose state is not Started.
func (up *Upload) Write(offset uint, data []byte, checksum *string) (uint, error) {
	if checksum != nil {
		var err error
		switch {
		case strings.HasPrefix(*checksum, "sha1 "):
			err = check(data, (*checksum)[5:], sha1.New())
		case strings.HasPrefix(*checksum, "md5 "):
			err = check(data, (*checksum)[4:], md5.New())
		case strings.HasPrefix(*checksum, "crc32 "):
			err = check(data, (*checksum)[6:], crc32.New(crc32.IEEETable))
		default:
			err = ErrInvalidAlgo("invalid checksum algorithm %q, supported algorithms are sha1, md5 and crc32")
		}
		if err != nil {
			return 0, err
		}
	}

	up.m.Lock()
	defer up.m.Unlock()

	if up.state != Started {
		return 0, fmt.Errorf("cannot write to %s upload", up.state.String())
	}
	if up.offset != offset {
		return 0, ErrInvalidOffset(fmt.Sprintf("got offset %v, expected %v", offset, up.offset))
	}
	n, err := up.writer.Write(data)
	if err != nil {
		up.terminate(Failed)
		return 0, err
	}
	up.offset += uint(n)

	return up.offset, nil
}

// Complete completes the upload. It is idempotent.
func (up *Upload) Complete() {
	up.m.Lock()
	defer up.m.Unlock()

	up.terminate(Completed)
}

// Cancel cancels the upload. It is idempotent.
func (up *Upload) Cancel() {
	up.m.Lock()
	defer up.m.Unlock()

	up.terminate(Cancelled)
}

// terminate sets the state of the upload to the given value (which cannot be
// Started) and calls the termination callback if any. Calling terminate on a
// upload whose state is not Started does nothing. The caller must acquire a
// lock on the upload mutex before calling terminate.
func (up *Upload) terminate(s State) {
	if s == Started {
		panic("cannot terminate upload with state Started")
	}

	if up.state != Started {
		return
	}
	up.state = s

	// Stop timer cleanly if set.
	if up.timer != nil && !up.timer.Stop() {
		<-up.timer.C
	}

	// Invoke callback in separate goroutine.
	if up.cb != nil {
		go up.cb(up.id, s)
	}
}

func check(data []byte, checksum string, hasher hash.Hash) error {
	_, err := hasher.Write(data)
	if err != nil {
		return err
	}
	if base64.URLEncoding.EncodeToString(hasher.Sum(nil)) != checksum {
		return ErrBadChecksum("invalid checksum")
	}
	return nil
}

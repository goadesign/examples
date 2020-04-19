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
		length    *int64
		offset    int64
		writer    io.Writer
	}

	// Info contains read-only information for a given upload.
	Info struct {
		// ID is the upload ID.
		ID string

		// State is the state of the upload.
		State State

		// Length is the total upload length in bytes, 0 if unknown.
		Length int64

		// Offset is the total number of bytes already uploaded.
		Offset int64

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

	// ErrTimedOut is the type of the error returned by Write when the
	// upload has expired
	ErrTimedOut string
)

func (e ErrInvalidAlgo) Error() string   { return string(e) }
func (e ErrBadChecksum) Error() string   { return string(e) }
func (e ErrInvalidOffset) Error() string { return string(e) }
func (e ErrTimedOut) Error() string      { return string(e) }

// New instantiates a new upload with the given id that writes to the given
// writer. length indicates the expected total number of bytes if known. If ttl
// is set then a timer is created that changes the state of the upload from
// Started to TimedOut upon expiration. cb - if not nil - is called once when
// the state of the upload changes from Started to another value. It is given
// the ID of the upload and the new state as arguments.
func New(id string, length *int64, metadata *string, writer io.Writer, ttl time.Duration) *Upload {
	startedAt := time.Now()
	var expiresAt time.Time
	if ttl > 0 {
		expiresAt = startedAt.Add(ttl)
	}

	return Resume(id, length, metadata, startedAt, expiresAt, writer)
}

// Resume resumes a persisted upload. expiresAt - if not zero - must be after
// startedAt.
func Resume(id string, length *int64, metadata *string, startedAt, expiresAt time.Time, writer io.Writer) *Upload {
	return &Upload{
		id:        id,
		metadata:  metadata,
		startedAt: startedAt,
		expiresAt: expiresAt,
		length:    length,
		writer:    writer,
		m:         &sync.RWMutex{},
	}
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
func (up *Upload) Write(offset int64, data io.ReadCloser, checksum *string) (n int64, err error) {
	// Make sure that upload hasn't expired
	if !up.expiresAt.IsZero() && up.expiresAt.Before(time.Now()) {
		up.terminate(TimedOut)
		return 0, ErrTimedOut(fmt.Sprintf("upload expired %v", up.expiresAt))
	}

	// make sure we always close the reader
	defer func() {
		if err != nil {
			up.terminate(Failed)
			data.Close()
			return
		}
		err = data.Close()
	}()

	up.m.Lock()
	defer up.m.Unlock()

	if up.state != Started {
		return 0, fmt.Errorf("cannot write to %s upload", up.state.String())
	}
	if up.offset != offset {
		return 0, ErrInvalidOffset(fmt.Sprintf("got offset %v, expected %v", offset, up.offset))
	}

	// Thank you Go stdlib for making the below so elegant.
	var (
		h   hash.Hash
		chk string
	)
	if checksum != nil {
		switch {
		case strings.HasPrefix(*checksum, "sha1 "):
			h = sha1.New()
			chk = (*checksum)[5:]
		case strings.HasPrefix(*checksum, "md5 "):
			h = md5.New()
			chk = (*checksum)[4:]
		case strings.HasPrefix(*checksum, "crc32 "):
			h = crc32.New(crc32.IEEETable)
			chk = (*checksum)[6:]
		default:
			return 0, ErrInvalidAlgo("invalid checksum algorithm %q, supported algorithms are sha1, md5 and crc32")
		}
	}
	w := up.writer
	if h != nil {
		w = io.MultiWriter(w, h)
	}
	n, err = io.Copy(w, data)
	if err != nil {
		return 0, err
	}
	if h != nil && base64.URLEncoding.EncodeToString(h.Sum(nil)) != chk {
		return 0, ErrBadChecksum("invalid checksum")
	}
	up.offset += int64(n)

	return
}

// Complete completes the upload. It is idempotent.
func (up *Upload) Complete() { up.terminate(Completed) }

// Cancel cancels the upload. It is idempotent.
func (up *Upload) Cancel() { up.terminate(Cancelled) }

// terminate sets the state of the upload to the given value (which cannot be
// Started) and calls the termination callback if any. Calling terminate on a
// upload whose state is not Started does nothing. The caller must acquire a
// lock on the upload mutex before calling terminate.
func (up *Upload) terminate(s State) {
	up.m.Lock()
	defer up.m.Unlock()

	if s == Started {
		panic("cannot terminate upload with state Started")
	}

	if up.state != Started {
		return
	}
	up.state = s
}

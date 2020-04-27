package tus

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"strings"
	"time"

	"goa.design/examples/tus/gen/tus"
	"goa.design/examples/tus/persist"
)

// Write validates the checksum if not nil then writes the data to the given
// writer, updates the given upload state and returns the updated upload offset
// value. checksum must start with "sha1 ", "md5 " or "crc32 " and be followed
// by the Base64 encoded checksum. It is an error to write to an upload whose
// state is not Started.
func Write(r io.ReadCloser, w io.Writer, u *persist.Upload, offset int64, checksum *string) (n int64, err error) {
	// Make sure upload hasn't completed, failed or timed out.
	if !u.Active() {
		return 0, tus.MakeGone(fmt.Errorf("upload %s", u.Status.String()))
	}

	// Make sure upload hasn't expired.
	if !u.ExpiresAt.IsZero() && u.ExpiresAt.Before(time.Now()) {
		u.Status = persist.TimedOut
		return 0, tus.MakeGone(fmt.Errorf("upload expired %v", u.ExpiresAt))
	}

	// Make sure offset is correct.
	if u.Offset != offset {
		return 0, tus.MakeInvalidOffset(fmt.Errorf("got offset %v, expected %v", offset, u.Offset))
	}

	// Always close the reader.
	defer func() {
		if err != nil {
			u.Status = persist.Failed
			if err := r.Close(); err != nil {
				// log err
			}
			err = tus.MakeInternal(err)
		}
		if err := r.Close(); err != nil {
			err = tus.MakeInternal(err)
		}
	}()

	// Thank you Go stdlib for making the below so elegant!
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
			return 0, tus.MakeInvalidChecksumAlgorithm(fmt.Errorf("invalid checksum algorithm %q, supported algorithms are sha1, md5 and crc32", *checksum))
		}
	}
	if h != nil {
		w = io.MultiWriter(w, h)
	}
	n, err = io.Copy(w, r)
	if err != nil {
		return 0, tus.MakeInternal(err)
	}
	if h != nil && base64.URLEncoding.EncodeToString(h.Sum(nil)) != chk {
		return 0, tus.MakeChecksumMismatch(fmt.Errorf("invalid checksum"))
	}
	u.Offset += int64(n)

	return
}

// Package uploaddownload implements the updown service using an os.Root
// directory handle so uploaded and downloaded paths cannot escape the
// configured base directory.
package uploaddownload

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	genupdown "goa.design/examples/upload_download/gen/updown"
)

type (
	// updownsrvc streams multipart uploads into and file downloads from a
	// confined root directory.
	updownsrvc struct {
		// root confines all directory creation, file writes, and file reads to
		// the service's configured base directory.
		root *os.Root
		// logger records completed file writes and close failures.
		logger *log.Logger
	}
)

// NewUpdown returns the updown service implementation confined to root. The
// caller owns root and must close it when the service stops serving requests.
func NewUpdown(root *os.Root, logger *log.Logger) genupdown.Service {
	return &updownsrvc{
		root:   root,
		logger: logger,
	}
}

// Upload streams each multipart file part from req into p.Dir beneath s.root,
// creating missing subdirectories with mode 0750 and overwriting existing files
// with the same base name. It returns invalid_media_type when p.ContentType
// lacks a valid multipart boundary, invalid_file_path when p.Dir is invalid,
// names a non-directory, or escapes s.root, invalid_multipart_request when the
// multipart stream is malformed or a part lacks a valid file name, and
// internal_error on storage or permission failures.
func (s *updownsrvc) Upload(
	ctx context.Context,
	p *genupdown.UploadPayload,
	req io.ReadCloser,
) (err error) {
	defer func() {
		if closeErr := req.Close(); closeErr != nil && err == nil {
			err = genupdown.MakeInternalError(closeErr)
		}
	}()

	_, params, parseErr := mime.ParseMediaType(p.ContentType)
	if parseErr != nil {
		return genupdown.MakeInvalidMediaType(parseErr)
	}
	boundary := params["boundary"]
	if boundary == "" {
		return genupdown.MakeInvalidMediaType(fmt.Errorf("missing multipart boundary"))
	}

	dir, cleanDir, dirErr := localizeUploadDir(p.Dir)
	if dirErr != nil {
		return genupdown.MakeInvalidFilePath(dirErr)
	}
	if mkErr := s.root.MkdirAll(dir, 0o750); mkErr != nil {
		if isClientPathError(mkErr) {
			return genupdown.MakeInvalidFilePath(mkErr)
		}
		return genupdown.MakeInternalError(mkErr)
	}

	mr := multipart.NewReader(req, boundary)
	for {
		part, partErr := mr.NextPart()
		if errors.Is(partErr, io.EOF) {
			return nil
		}
		if partErr != nil {
			return genupdown.MakeInvalidMultipartRequest(partErr)
		}

		name := part.FileName()
		if !isValidPartFileName(name) {
			return genupdown.MakeInvalidMultipartRequest(fmt.Errorf("invalid part filename %q", name))
		}

		n, writeErr := writePart(s.root, filepath.Join(dir, name), part)
		if writeErr != nil {
			if errors.Is(writeErr, io.ErrUnexpectedEOF) {
				return genupdown.MakeInvalidMultipartRequest(writeErr)
			}
			if isClientPathError(writeErr) {
				return genupdown.MakeInvalidFilePath(writeErr)
			}
			return genupdown.MakeInternalError(writeErr)
		}
		s.logger.Printf("Written %d bytes to %q", n, path.Join(cleanDir, name))
	}
}

// Download opens the file at p beneath s.root and returns its byte length and
// an open reader for streaming the response body. It returns invalid_file_path
// when p is invalid, missing, names a directory, or escapes s.root, and
// internal_error on storage or permission failures.
func (s *updownsrvc) Download(
	ctx context.Context,
	p string,
) (*genupdown.DownloadResult, io.ReadCloser, error) {
	if p == "." {
		return nil, nil, genupdown.MakeInvalidFilePath(fmt.Errorf("%q is a directory", p))
	}
	localized, locErr := filepath.Localize(p)
	if locErr != nil {
		return nil, nil, genupdown.MakeInvalidFilePath(locErr)
	}

	f, err := s.root.Open(localized)
	if err != nil {
		if isClientPathError(err) {
			return nil, nil, genupdown.MakeInvalidFilePath(err)
		}
		return nil, nil, genupdown.MakeInternalError(err)
	}
	fi, err := f.Stat()
	if err != nil {
		if closeErr := f.Close(); closeErr != nil {
			s.logger.Printf("failed to close %q: %v", p, closeErr)
		}
		return nil, nil, genupdown.MakeInternalError(err)
	}
	if fi.IsDir() {
		if closeErr := f.Close(); closeErr != nil {
			s.logger.Printf("failed to close %q: %v", p, closeErr)
		}
		return nil, nil, genupdown.MakeInvalidFilePath(fmt.Errorf("%q is a directory", p))
	}

	return &genupdown.DownloadResult{Length: fi.Size()}, f, nil
}

// localizeUploadDir validates the slash-separated HTTP wildcard directory raw
// and converts it into an OS-local relative path alongside its clean slash form.
func localizeUploadDir(raw string) (string, string, error) {
	if raw == "" {
		return ".", ".", nil
	}
	if strings.HasPrefix(raw, "/") {
		return "", "", fmt.Errorf("upload directory %q must be relative", raw)
	}
	trimmed := strings.TrimRight(raw, "/")
	localized, err := filepath.Localize(trimmed)
	if err != nil {
		return "", "", err
	}
	return localized, trimmed, nil
}

// isClientPathError reports whether err indicates an invalid, missing,
// non-directory, or root-escaping client path rather than a server storage or
// permission failure.
func isClientPathError(err error) bool {
	if errors.Is(err, fs.ErrNotExist) ||
		errors.Is(err, fs.ErrExist) ||
		errors.Is(err, fs.ErrInvalid) ||
		errors.Is(err, syscall.EINVAL) ||
		errors.Is(err, syscall.ENOTDIR) ||
		errors.Is(err, syscall.EISDIR) ||
		errors.Is(err, syscall.ELOOP) ||
		errors.Is(err, syscall.ENAMETOOLONG) {
		return true
	}
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		leaf := error(pathErr)
		for {
			next := errors.Unwrap(leaf)
			if next == nil {
				break
			}
			leaf = next
		}
		msg := leaf.Error()
		return msg == "path escapes from parent" || msg == "empty path"
	}
	return false
}

// isValidPartFileName reports whether name is a single local file name suitable
// for creating inside an upload directory.
func isValidPartFileName(name string) bool {
	return filepath.IsLocal(name) && name != "." && !strings.ContainsAny(name, `/\`)
}

// writePart creates or truncates relPath inside root, copies src into it, and
// closes the file before returning.
func writePart(root *os.Root, relPath string, src io.Reader) (int64, error) {
	f, err := root.Create(relPath)
	if err != nil {
		return 0, err
	}
	n, err := io.Copy(f, src)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return n, err
}

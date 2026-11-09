// These tests verify that the updown HTTP server confines uploaded and
// downloaded paths to the configured os.Root directory and returns the declared
// Goa error responses for invalid paths, multipart inputs, and storage faults.
package uploaddownload_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	uploaddownload "goa.design/examples/upload_download"
	genupdownsvr "goa.design/examples/upload_download/gen/http/updown/server"
	genupdown "goa.design/examples/upload_download/gen/updown"
	goahttp "goa.design/goa/v3/http"
)

type (
	// multipartItem describes one raw Content-Disposition part in a test
	// multipart request body.
	multipartItem struct {
		// disposition is the complete Content-Disposition header value.
		disposition string
		// content is the part payload written after the part headers.
		content string
	}
)

// TestUpload verifies multipart streaming inside the root directory and
// rejection of escaping directories, escaping symlinks, non-directory paths,
// invalid part names, malformed Content-Type headers, and read-only directories.
func TestUpload(t *testing.T) {
	parentDir := t.TempDir()
	rootDir := filepath.Join(parentDir, "public")
	outsideDir := filepath.Join(parentDir, "outside")
	if err := os.MkdirAll(rootDir, 0o750); err != nil {
		t.Fatalf("create root dir: %v", err)
	}
	if err := os.MkdirAll(outsideDir, 0o750); err != nil {
		t.Fatalf("create outside dir: %v", err)
	}
	existingFile := filepath.Join(rootDir, "occupied.txt")
	if err := os.WriteFile(existingFile, []byte("keep"), 0o600); err != nil {
		t.Fatalf("write occupied.txt: %v", err)
	}
	readOnlyDir := filepath.Join(rootDir, "readonly")
	if err := os.MkdirAll(readOnlyDir, 0o500); err != nil {
		t.Fatalf("create readonly dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chmod(readOnlyDir, 0o750); err != nil {
			t.Errorf("restore readonly dir permissions: %v", err)
		}
	})

	symlinkSupported := true
	if err := os.Symlink(outsideDir, filepath.Join(rootDir, "escape_link")); err != nil {
		if runtime.GOOS == "windows" {
			symlinkSupported = false
		} else {
			t.Fatalf("create symlink: %v", err)
		}
	}

	mux := newTestServer(t, rootDir)
	boundary := "testboundary"
	validType := "multipart/form-data; boundary=" + boundary
	validParts := []multipartItem{
		{
			disposition: `form-data; name="file1"; filename="first.txt"`,
			content:     "hello world",
		},
		{
			disposition: `form-data; name="file2"; filename="second.txt"`,
			content:     "second payload",
		},
	}

	tests := []struct {
		name       string
		target     string
		headerType string
		parts      []multipartItem
		wantStatus int
		wantError  string
		skip       bool
		verify     func(t *testing.T)
	}{
		{
			name:       "uploads multiple parts into nested directory with restrictive permissions",
			target:     "/upload/nested/sub",
			headerType: validType,
			parts:      validParts,
			wantStatus: http.StatusNoContent,
			verify: func(t *testing.T) {
				firstPath := filepath.Join(rootDir, "nested", "sub", "first.txt")
				gotFirst, err := os.ReadFile(firstPath)
				if err != nil {
					t.Fatalf("read first.txt: %v", err)
				}
				if string(gotFirst) != "hello world" {
					t.Errorf("first.txt = %q, want %q", string(gotFirst), "hello world")
				}
				secondPath := filepath.Join(rootDir, "nested", "sub", "second.txt")
				gotSecond, err := os.ReadFile(secondPath)
				if err != nil {
					t.Fatalf("read second.txt: %v", err)
				}
				if string(gotSecond) != "second payload" {
					t.Errorf("second.txt = %q, want %q", string(gotSecond), "second payload")
				}
				if runtime.GOOS != "windows" {
					fi, statErr := os.Stat(filepath.Join(rootDir, "nested", "sub"))
					if statErr != nil {
						t.Fatalf("stat nested/sub: %v", statErr)
					}
					if perm := fi.Mode().Perm(); perm&0o027 != 0 {
						t.Errorf("directory mode = %04o, want no group-write or world bits", perm)
					}
				}
			},
		},
		{
			name:       "uploads when directory URL has trailing slash",
			target:     "/upload/nested/trailing/",
			headerType: validType,
			parts: []multipartItem{
				{
					disposition: `form-data; name="file"; filename="trail.txt"`,
					content:     "trailing ok",
				},
			},
			wantStatus: http.StatusNoContent,
			verify: func(t *testing.T) {
				trailPath := filepath.Join(rootDir, "nested", "trailing", "trail.txt")
				got, err := os.ReadFile(trailPath)
				if err != nil {
					t.Fatalf("read trail.txt: %v", err)
				}
				if string(got) != "trailing ok" {
					t.Errorf("trail.txt = %q, want %q", string(got), "trailing ok")
				}
			},
		},
		{
			name:       "uploads directly into root directory when wildcard is empty",
			target:     "/upload/",
			headerType: validType,
			parts: []multipartItem{
				{
					disposition: `form-data; name="file"; filename="root.txt"`,
					content:     "in root",
				},
			},
			wantStatus: http.StatusNoContent,
			verify: func(t *testing.T) {
				got, err := os.ReadFile(filepath.Join(rootDir, "root.txt"))
				if err != nil {
					t.Fatalf("read root.txt: %v", err)
				}
				if string(got) != "in root" {
					t.Errorf("root.txt = %q, want %q", string(got), "in root")
				}
			},
		},
		{
			name:       "rejects URL-encoded parent traversal in dir",
			target:     "/upload/..%2Foutside",
			headerType: validType,
			parts: []multipartItem{
				{
					disposition: `form-data; name="file"; filename="pwned.txt"`,
					content:     "escaped",
				},
			},
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_file_path",
			verify: func(t *testing.T) {
				escapedPath := filepath.Join(outsideDir, "pwned.txt")
				if _, err := os.Stat(escapedPath); !os.IsNotExist(err) {
					t.Errorf("outside pwned.txt stat err = %v, want not exist", err)
				}
			},
		},
		{
			name:       "rejects traversal after nonexistent prefix without creating prefix",
			target:     "/upload/transient/..%2F..%2Foutside",
			headerType: validType,
			parts:      validParts,
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_file_path",
			verify: func(t *testing.T) {
				transientPath := filepath.Join(rootDir, "transient")
				if _, err := os.Stat(transientPath); !os.IsNotExist(err) {
					t.Errorf("transient dir stat err = %v, want not exist", err)
				}
			},
		},
		{
			name:       "rejects symlink escaping root directory",
			target:     "/upload/escape_link",
			headerType: validType,
			parts: []multipartItem{
				{
					disposition: `form-data; name="file"; filename="via_link.txt"`,
					content:     "escaped",
				},
			},
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_file_path",
			skip:       !symlinkSupported,
			verify: func(t *testing.T) {
				viaLink := filepath.Join(outsideDir, "via_link.txt")
				if _, err := os.Stat(viaLink); !os.IsNotExist(err) {
					t.Errorf("outside via_link.txt stat err = %v, want not exist", err)
				}
			},
		},
		{
			name:       "rejects upload directory occupied by existing regular file",
			target:     "/upload/occupied.txt",
			headerType: validType,
			parts:      validParts,
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_file_path",
		},
		{
			name:       "returns internal_error when parent directory denies write permission",
			target:     "/upload/readonly/sub",
			headerType: validType,
			parts:      validParts,
			wantStatus: http.StatusInternalServerError,
			wantError:  "internal_error",
			skip:       runtime.GOOS == "windows" || os.Geteuid() == 0,
		},
		{
			name:       "rejects multipart part without filename",
			target:     "/upload/forms",
			headerType: validType,
			parts: []multipartItem{
				{
					disposition: `form-data; name="field"`,
					content:     "value",
				},
			},
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_multipart_request",
		},
		{
			name:       "rejects dot filename in multipart part",
			target:     "/upload/forms",
			headerType: validType,
			parts: []multipartItem{
				{
					disposition: `form-data; name="file"; filename="."`,
					content:     "value",
				},
			},
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_multipart_request",
		},
		{
			name:       "rejects dot-dot filename in multipart part",
			target:     "/upload/forms",
			headerType: validType,
			parts: []multipartItem{
				{
					disposition: `form-data; name="file"; filename=".."`,
					content:     "value",
				},
			},
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_multipart_request",
		},
		{
			name:       "rejects unparseable Content-Type parameter before creating directory",
			target:     "/upload/uncreated",
			headerType: `multipart/form-data; boundary="unterminated`,
			parts:      validParts,
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid_media_type",
			verify: func(t *testing.T) {
				uncreatedPath := filepath.Join(rootDir, "uncreated")
				if _, err := os.Stat(uncreatedPath); !os.IsNotExist(err) {
					t.Errorf("uncreated dir stat err = %v, want not exist", err)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.skip {
				t.Skip("precondition unavailable in test environment")
			}
			body := buildMultipartBody(boundary, tc.parts)
			req := httptest.NewRequest(http.MethodPost, tc.target, body)
			req.Header.Set("Content-Type", tc.headerType)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			if rec.Code != tc.wantStatus {
				t.Errorf(
					"status = %d, want %d (body: %s)",
					rec.Code,
					tc.wantStatus,
					rec.Body.String(),
				)
			}
			if tc.wantError != "" {
				if gotErr := rec.Header().Get("goa-error"); gotErr != tc.wantError {
					t.Errorf("goa-error header = %q, want %q", gotErr, tc.wantError)
				}
			}
			if tc.verify != nil {
				tc.verify(t)
			}
		})
	}
}

// TestDownload verifies streaming files inside the root directory and rejecting
// escaping paths, escaping symlinks, directories, and missing files.
func TestDownload(t *testing.T) {
	parentDir := t.TempDir()
	rootDir := filepath.Join(parentDir, "public")
	if err := os.MkdirAll(filepath.Join(rootDir, "sub"), 0o750); err != nil {
		t.Fatalf("create sub dir: %v", err)
	}
	secretPath := filepath.Join(parentDir, "secret.txt")
	if err := os.WriteFile(secretPath, []byte("top-secret"), 0o600); err != nil {
		t.Fatalf("write secret.txt: %v", err)
	}
	wantPayload := "streamed content"
	helloPath := filepath.Join(rootDir, "sub", "hello.txt")
	if err := os.WriteFile(helloPath, []byte(wantPayload), 0o600); err != nil {
		t.Fatalf("write hello.txt: %v", err)
	}
	symlinkSupported := true
	if err := os.Symlink(secretPath, filepath.Join(rootDir, "secret_link.txt")); err != nil {
		if runtime.GOOS == "windows" {
			symlinkSupported = false
		} else {
			t.Fatalf("create secret symlink: %v", err)
		}
	}

	mux := newTestServer(t, rootDir)

	tests := []struct {
		name       string
		target     string
		wantStatus int
		wantError  string
		wantBody   string
		skip       bool
	}{
		{
			name:       "downloads nested file with Content-Length",
			target:     "/download/sub/hello.txt",
			wantStatus: http.StatusOK,
			wantBody:   wantPayload,
		},
		{
			name:       "rejects URL-encoded parent traversal",
			target:     "/download/..%2Fsecret.txt",
			wantStatus: http.StatusNotFound,
			wantError:  "invalid_file_path",
		},
		{
			name:       "rejects symlink escaping root directory",
			target:     "/download/secret_link.txt",
			wantStatus: http.StatusNotFound,
			wantError:  "invalid_file_path",
			skip:       !symlinkSupported,
		},
		{
			name:       "rejects directory path",
			target:     "/download/sub",
			wantStatus: http.StatusNotFound,
			wantError:  "invalid_file_path",
		},
		{
			name:       "rejects missing file",
			target:     "/download/sub/missing.txt",
			wantStatus: http.StatusNotFound,
			wantError:  "invalid_file_path",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.skip {
				t.Skip("symlinks unavailable in test environment")
			}
			req := httptest.NewRequest(http.MethodGet, tc.target, nil)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			if rec.Code != tc.wantStatus {
				t.Errorf(
					"status = %d, want %d (body: %s)",
					rec.Code,
					tc.wantStatus,
					rec.Body.String(),
				)
			}
			if tc.wantError != "" {
				if gotErr := rec.Header().Get("goa-error"); gotErr != tc.wantError {
					t.Errorf("goa-error header = %q, want %q", gotErr, tc.wantError)
				}
			}
			if tc.wantStatus == http.StatusOK {
				if gotBody := rec.Body.String(); gotBody != tc.wantBody {
					t.Errorf("body = %q, want %q", gotBody, tc.wantBody)
				}
				wantLen := fmt.Sprintf("%d", len(tc.wantBody))
				if gotLen := rec.Header().Get("Content-Length"); gotLen != wantLen {
					t.Errorf("Content-Length = %q, want %q", gotLen, wantLen)
				}
			}
		})
	}
}

// newTestServer mounts an updown HTTP server backed by an os.Root at rootDir.
func newTestServer(t *testing.T, rootDir string) goahttp.Muxer {
	t.Helper()
	root, err := os.OpenRoot(rootDir)
	if err != nil {
		t.Fatalf("open root %q: %v", rootDir, err)
	}
	t.Cleanup(func() {
		if closeErr := root.Close(); closeErr != nil {
			t.Errorf("close root %q: %v", rootDir, closeErr)
		}
	})

	svc := uploaddownload.NewUpdown(root, log.New(io.Discard, "", 0))
	endpoints := genupdown.NewEndpoints(svc)
	mux := goahttp.NewMuxer()
	srv := genupdownsvr.New(
		endpoints,
		mux,
		goahttp.RequestDecoder,
		goahttp.ResponseEncoder,
		nil,
		nil,
	)
	genupdownsvr.Mount(mux, srv)
	return mux
}

// buildMultipartBody serializes raw Content-Disposition parts using boundary.
func buildMultipartBody(boundary string, parts []multipartItem) *bytes.Buffer {
	var buf bytes.Buffer
	for _, p := range parts {
		fmt.Fprintf(
			&buf,
			"--%s\r\nContent-Disposition: %s\r\n\r\n%s\r\n",
			boundary,
			p.disposition,
			p.content,
		)
	}
	fmt.Fprintf(&buf, "--%s--\r\n", boundary)
	return &buf
}

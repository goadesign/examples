package persist

import (
	"testing"
	"time"
)

var (
	simpleUpload = &Upload{
		ID:        "1",
		StartedAt: time.Now().Add(time.Duration(-1) * time.Hour),
		Status:    Started,
	}

	expiringUpload = &Upload{
		ID:        "2",
		StartedAt: time.Now().Add(time.Duration(-1) * time.Hour),
		ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour),
		Status:    Started,
	}

	expiredUpload = &Upload{
		ID:        "3",
		StartedAt: time.Now().Add(time.Duration(-1) * time.Hour),
		ExpiresAt: time.Now().Add(time.Duration(-10) * time.Minute),
		Status:    TimedOut,
	}

	completedUpload = &Upload{
		ID:        "4",
		StartedAt: time.Now().Add(time.Duration(-1) * time.Hour),
		ExpiresAt: time.Now().Add(time.Duration(-10) * time.Minute),
		Status:    Completed,
	}

	l = int64(100)

	knownLengthUpload = &Upload{
		ID:        "5",
		StartedAt: time.Now().Add(time.Duration(-1) * time.Hour),
		Status:    Started,
		Length:    &l,
	}

	metadataUpload = &Upload{
		ID:        "6",
		StartedAt: time.Now().Add(time.Duration(-1) * time.Hour),
		Status:    Started,
		Metadata:  "metadata",
	}
)

func TestInMemoryStore(t *testing.T) {
	StoreTest(NewInMemory(), t)
}

func StoreTest(s Store, t *testing.T) {
	cases := []struct {
		Name   string
		Upload *Upload
		ID     string
	}{
		{"simple", simpleUpload, "a"},
		{"expiring", expiringUpload, "b"},
		{"expired", expiredUpload, "c"},
		{"completed", completedUpload, "d"},
		{"knownLength", knownLengthUpload, "e"},
		{"metadata", metadataUpload, "f"},
		{"nil", nil, "g"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if err := s.Save(c.ID, c.Upload); err != nil {
				t.Errorf("failed to save: %s", err)
			}
			u, err := s.Load(c.ID)
			if err != nil {
				t.Errorf("failed to load: %s", err)
				return
			}
			if u != c.Upload {
				t.Errorf("got %v, expected %v", u, c.Upload)
			}
		})
	}
	if err := s.Save("a", completedUpload); err != nil {
		t.Errorf("failed to save over: %s", err)
		return
	}
	u, err := s.Load("a")
	if err != nil {
		t.Errorf("reload: got %v, expected %v", u, completedUpload)
	}
}

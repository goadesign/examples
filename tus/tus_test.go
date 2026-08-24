// These tests verify that terminating an upload removes its resumable state.
package tus

import (
	"context"
	"errors"
	"testing"

	gentus "goa.design/examples/tus/gen/tus"
	"goa.design/examples/tus/persist"
	goa "goa.design/goa/v3/pkg"
)

func TestDeleteRemovesUpload(t *testing.T) {
	for _, status := range []persist.Status{persist.Started, persist.Completed} {
		t.Run(status.String(), func(t *testing.T) {
			store := persist.NewInMemory()
			if err := store.Save("upload", &persist.Upload{ID: "upload", Status: status}); err != nil {
				t.Fatalf("save upload: %v", err)
			}
			svc := &tussvc{store: store}
			payload := &gentus.DeletePayload{ID: "upload", TusResumable: TusResumable}

			if _, err := svc.Delete(context.Background(), payload); err != nil {
				t.Fatalf("delete upload: %v", err)
			}
			_, err := svc.Head(context.Background(), &gentus.HeadPayload{ID: "upload", TusResumable: TusResumable})
			var serviceErr *goa.ServiceError
			if !errors.As(err, &serviceErr) || serviceErr.Name != "NotFound" {
				t.Fatalf("head deleted upload: got %v, expected NotFound", err)
			}
		})
	}
}

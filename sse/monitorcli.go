package monitorapi

import (
	"context"

	"goa.design/clue/log"
	genmonitor "goa.design/examples/sse/gen/monitor"
)

// InteractWithStreams interacts with the monitor service streams.
func InteractWithStreams(ctx context.Context, stream genmonitor.MonitorClientStream) {
	for {
		event, err := stream.RecvWithContext(ctx)
		if err != nil {
			log.Errorf(ctx, err, "error receiving event")
			return
		}
		log.Printf(ctx, "received event: %#v", event)
	}
}

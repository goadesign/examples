package monitorapi

import (
	"context"
	"embed"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"goa.design/clue/log"
	monitor "goa.design/examples/sse/gen/monitor"
)

// monitor service implementation.
type monitorsrvc struct{}

//go:embed public
var StaticFS embed.FS

// NewMonitor returns the monitor service implementation.
func NewMonitor() monitor.Service {
	return &monitorsrvc{}
}

// Monitor CPU and memory usage via Server-Sent Events
func (s *monitorsrvc) Monitor(ctx context.Context, stream monitor.MonitorServerStream) (err error) {
	log.Printf(ctx, "monitor.monitor started")

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf(ctx, "monitor.monitor stopped: client disconnected")
			return nil
		case <-ticker.C:
			cpuPercent, err := cpu.Percent(0, false)
			if err != nil {
				log.Printf(ctx, "CPU percent error: %v", err)
				continue
			}
			memStat, err := mem.VirtualMemory()
			if err != nil {
				log.Printf(ctx, "Memory info error: %v", err)
				continue
			}
			usage := &monitor.Usage{
				Timestamp: time.Now().Format(time.RFC3339),
				CPU:       int(cpuPercent[0]),
				Memory:    int(memStat.UsedPercent),
			}

			if err := stream.Send(usage); err != nil {
				return err
			}

			log.Printf(ctx, "%s CPU=%d%% Memory=%d%%", usage.Timestamp, usage.CPU, usage.Memory)
		}
	}
}

package tracer

import (
	"time"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func batchOptions() []sdktrace.BatchSpanProcessorOption {
	return []sdktrace.BatchSpanProcessorOption{
		sdktrace.WithMaxExportBatchSize(512),
		sdktrace.WithMaxQueueSize(2048),
		sdktrace.WithExportTimeout(30 * time.Second),
	}
}

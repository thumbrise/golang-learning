package tracer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

var ErrTraceExporterNew = errors.New("failed to create trace exporter")

func NewOTELExporter(ctx context.Context, cfg Config) (*otlptrace.Exporter, error) {
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(cfg.OTLPURL),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithURLPath("/v1/traces"),
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
		otlptracehttp.WithTimeout(5*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTraceExporterNew, err)
	}

	return exp, nil
}

func NewStdOutExporter() (*stdouttrace.Exporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

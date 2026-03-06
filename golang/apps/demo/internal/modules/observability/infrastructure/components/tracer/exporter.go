package tracer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

var ErrTraceExporterNew = errors.New("failed to create trace exporter")

func NewOTELExporter(ctx context.Context, cfg Config) (*otlptrace.Exporter, error) {
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(cfg.OTLPURL),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTraceExporterNew, err)
	}

	return exp, nil
}

func NewStdOutExporter() (*stdouttrace.Exporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

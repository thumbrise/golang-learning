package tracer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

var ErrTraceExporterNew = errors.New("failed to create trace exporter")

func NewOTELExporter(ctx context.Context, cfg components.OTLPConfig) (*otlptrace.Exporter, error) {
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(cfg.URL),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithTimeout(5*time.Second),
		otlptracegrpc.WithHeaders(map[string]string{
			cfg.TokenKey: cfg.TokenValue,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTraceExporterNew, err)
	}

	return exp, nil
}

func NewStdOutExporter() (*stdouttrace.Exporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

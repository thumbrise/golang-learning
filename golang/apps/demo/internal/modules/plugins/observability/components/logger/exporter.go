package logger

import (
	"context"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
)

func NewExporter(ctx context.Context, cfg components.OTLPConfig) (*otlploggrpc.Exporter, error) {
	return otlploggrpc.New(
		ctx,
		otlploggrpc.WithEndpoint(cfg.URL),
		otlploggrpc.WithInsecure(),
		otlploggrpc.WithTimeout(5*time.Second),
		otlploggrpc.WithHeaders(map[string]string{
			cfg.TokenKey: cfg.TokenValue,
		}),
	)
}

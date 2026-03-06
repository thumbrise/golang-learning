package meter

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
)

func NewExporter(ctx context.Context, cfg infrastructure.OTLPConfig) (*otlpmetricgrpc.Exporter, error) {
	return otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(cfg.URL),
		otlpmetricgrpc.WithCompressor("gzip"),
	)
}

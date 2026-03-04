package logger

import (
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/app"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

type CommonAttributesProvider struct {
	attrs []slog.Attr
}

func (c *CommonAttributesProvider) GetAttributes() []slog.Attr {
	return c.attrs
}

func NewCommonAttributesProvider(cfg app.Config) *CommonAttributesProvider {
	attrs := []slog.Attr{
		slog.String(string(semconv.ServiceNameKey), cfg.Name),
		slog.String(string(semconv.ServiceVersionKey), cfg.Version),
		slog.String(string(semconv.DeploymentEnvironmentKey), cfg.Environment),
	}

	return &CommonAttributesProvider{attrs: attrs}
}

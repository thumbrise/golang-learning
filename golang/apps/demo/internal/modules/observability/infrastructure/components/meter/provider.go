package meter

import (
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type Provider struct {
	p   *sdkmetric.MeterProvider
	cfg app.Config
}

func NewProvider(cfg app.Config, p *sdkmetric.MeterProvider) *Provider {
	return &Provider{cfg: cfg, p: p}
}

func (t *Provider) Meter() metric.Meter {
	return t.p.Meter(t.cfg.Name)
}

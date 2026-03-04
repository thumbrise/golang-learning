package tracer

import (
	"github.com/thumbrise/demo/golang-demo/internal/app"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type Provider struct {
	p   *sdktrace.TracerProvider
	cfg app.Config
}

func NewProvider(cfg app.Config, p *sdktrace.TracerProvider) *Provider {
	return &Provider{cfg: cfg, p: p}
}

func (t *Provider) Tracer() trace.Tracer {
	return t.p.Tracer(t.cfg.Name)
}

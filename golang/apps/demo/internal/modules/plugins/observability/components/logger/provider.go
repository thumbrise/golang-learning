package logger

import (
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type Provider struct {
	p   *sdklog.LoggerProvider
	cfg app.Config
}

func NewProvider(cfg app.Config, p *sdklog.LoggerProvider) *Provider {
	return &Provider{cfg: cfg, p: p}
}

func (t *Provider) Logger() log.Logger { //nolint:ireturn //specific case
	return t.p.Logger(t.cfg.Name)
}

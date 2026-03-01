package homepage

import (
	"context"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/endpoints/http"
	"go.uber.org/fx"
)

type FBootloader struct {
	router *http.HomePageRouter
}

func NewFBootloader(router *http.HomePageRouter) *FBootloader {
	return &FBootloader{router: router}
}

func (b *FBootloader) Name() string {
	return "homepage"
}
func (b *FBootloader) Bind() []fx.Option {
	slog.Debug("HIIII, im bind homepage")

	return []fx.Option{
		fx.Provide(NewFBootloader),
		fx.Provide(http.NewHomePageRouter),
	}
}
func (b *FBootloader) BeforeStart() interface{} {
	return func() {
		b.router.Register()
		slog.Debug("HIIII, im register homepage")
	}
}

func (b *FBootloader) OnStart(ctx context.Context) error {
	slog.Debug("HIIII, im boot homepage")

	return nil
}
func (b *FBootloader) Shutdown(ctx context.Context) error {
	slog.Debug("HIIII, im shutdown homepage")
	return nil
}

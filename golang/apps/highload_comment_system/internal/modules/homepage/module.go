package homepage

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/infrastucture/generator"
)

var Bindings = wire.NewSet(
	NewModule,
	http.NewHomePageRouter,
	generator.NewGenerator,
)

type Module struct {
	router *http.HomePageRouter
}

func NewModule(router *http.HomePageRouter) *Module {
	return &Module{router: router}
}

func (m *Module) Name() string {
	return "homepage"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	m.router.Register()

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}

func (m *Module) LongRun(ctx context.Context) error {
	return nil
}

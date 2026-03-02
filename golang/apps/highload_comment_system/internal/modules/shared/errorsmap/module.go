package errorsmap

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/endpoints/http"
)

var Bindings = wire.NewSet(
	NewModule,
	http.NewErrorsMapRouter,
	http.NewErrorsMapMiddleware,
)

type Module struct {
	router *http.ErrorsMapRouter
}

func NewModule(router *http.ErrorsMapRouter) *Module {
	return &Module{router: router}
}

func (m *Module) Name() string {
	return "errorsmap"
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

package errorsmap

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/components"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/endpoints/http"
)

var Bindings = wire.NewSet(
	NewModule,
	http.NewErrorsMapMiddleware,
)

type Module struct {
	mapErrorsMiddleware *http.ErrorsMapMiddleware
	kernel              *components.Kernel
}

func NewModule(kernel *components.Kernel, mapErrorsMiddleware *http.ErrorsMapMiddleware) *Module {
	return &Module{kernel: kernel, mapErrorsMiddleware: mapErrorsMiddleware}
}

func (m *Module) Name() string {
	return "errorsmap"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	m.kernel.Gin().Use(m.mapErrorsMiddleware.Handler())

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}

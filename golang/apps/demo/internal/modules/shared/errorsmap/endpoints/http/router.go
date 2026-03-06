package http

import (
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http/components"
)

type ErrorsMapRouter struct {
	mapErrorsMiddleware *ErrorsMapMiddleware
	kernel              *components.Kernel
}

func NewErrorsMapRouter(mapErrorsMiddleware *ErrorsMapMiddleware, kernel *components.Kernel) *ErrorsMapRouter {
	return &ErrorsMapRouter{mapErrorsMiddleware: mapErrorsMiddleware, kernel: kernel}
}

func (h *ErrorsMapRouter) Register() {
	h.kernel.Gin().Use(h.mapErrorsMiddleware.Handler())
}

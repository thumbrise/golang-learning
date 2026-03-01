package http

import (
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
)

type ErrorsMapRouter struct {
	mapErrorsMiddleware *ErrorsMapMiddleware
	kernel              *http.Kernel
}

func NewErrorsMapRouter(mapErrorsMiddleware *ErrorsMapMiddleware, kernel *http.Kernel) *ErrorsMapRouter {
	return &ErrorsMapRouter{mapErrorsMiddleware: mapErrorsMiddleware, kernel: kernel}
}

func (h *ErrorsMapRouter) Register() {
	h.kernel.Gin().Use(h.mapErrorsMiddleware.Handler())
}

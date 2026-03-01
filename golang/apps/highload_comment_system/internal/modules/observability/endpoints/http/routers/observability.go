package routers

import (
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/middlewares"
	http2 "github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type ObservabilityRouter struct {
	kernel                  *http2.Kernel
	observabilityMiddleware *middlewares.ObservabilityMiddleware
}

func NewObservabilityRouter(kernel *http2.Kernel, observabilityMiddleware *middlewares.ObservabilityMiddleware) *ObservabilityRouter {
	return &ObservabilityRouter{kernel: kernel, observabilityMiddleware: observabilityMiddleware}
}

func (h *ObservabilityRouter) Register() {
	h.kernel.Gin().Use(h.observabilityMiddleware.Handler())

	p := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem: "gin",
	})
	p.Use(h.kernel.Gin())
	// h.kernel.Gin().Use(otelgin.Middleware(h.cfg.Name))
}

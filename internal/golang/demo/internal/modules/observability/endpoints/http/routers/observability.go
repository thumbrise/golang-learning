package routers

import (
	ginprometheus "github.com/zsais/go-gin-prometheus"
	http2 "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/kernels/http"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/observability/endpoints/http/middlewares"
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

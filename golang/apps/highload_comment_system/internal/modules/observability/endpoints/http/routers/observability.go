package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/middlewares"
	http2 "github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
)

type ObservabilityRouter struct {
	kernel                  *http2.Kernel
	observabilityMiddleware *middlewares.ObservabilityMiddleware
	prometheusRegistry      *prometheus.Registry
}

func NewObservabilityRouter(kernel *http2.Kernel, observabilityMiddleware *middlewares.ObservabilityMiddleware, prometheusRegistry *prometheus.Registry) *ObservabilityRouter {
	return &ObservabilityRouter{kernel: kernel, observabilityMiddleware: observabilityMiddleware, prometheusRegistry: prometheusRegistry}
}

func (h *ObservabilityRouter) Register() {
	// TODO: В роутах оставить только те миддлвары которые цепляются к внутренним группам. Остальные мидлвары надо регистрировать в модуле??
	h.registerMetrics()
	h.kernel.Gin().Use(h.observabilityMiddleware.Handler())
}

func (h *ObservabilityRouter) registerMetrics() {
	h.prometheusRegistry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	opts := promhttp.HandlerOpts{}
	handler := gin.WrapH(promhttp.HandlerFor(h.prometheusRegistry, opts))
	h.kernel.Gin().GET("/metrics", handler)
}

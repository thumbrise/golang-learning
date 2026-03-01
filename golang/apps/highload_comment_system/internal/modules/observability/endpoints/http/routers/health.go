package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	http2 "github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
)

type HealthRouter struct {
	kernel *http2.Kernel
}

func NewHealthRouter(kernel *http2.Kernel) *HealthRouter {
	return &HealthRouter{kernel: kernel}
}

func (h *HealthRouter) Register() {
	h.kernel.Gin().HEAD("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	http2 "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/kernels/http"
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

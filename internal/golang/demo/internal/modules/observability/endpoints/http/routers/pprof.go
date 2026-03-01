package routers

import (
	"github.com/gin-contrib/pprof"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/kernels/http"
)

type PprofRouter struct {
	kernel *http.Kernel
}

func NewPprofRouter(kernel *http.Kernel) *PprofRouter {
	return &PprofRouter{kernel: kernel}
}

func (h *PprofRouter) Register() {
	pprof.Register(h.kernel.Gin())
}

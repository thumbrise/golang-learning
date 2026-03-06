package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
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

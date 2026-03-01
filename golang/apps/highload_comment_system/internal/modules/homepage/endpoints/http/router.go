package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpkernel "github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/infrastucture/generator"
)

type HomePageRouter struct {
	kernel    *httpkernel.Kernel
	generator *generator.Generator
}

func NewHomePageRouter(generator *generator.Generator, kernel *httpkernel.Kernel) *HomePageRouter {
	return &HomePageRouter{generator: generator, kernel: kernel}
}

func (h *HomePageRouter) Register() {
	h.kernel.Gin().GET("/", func(ctx *gin.Context) {
		routes := h.kernel.Gin().Routes()

		err := h.generator.Write(routes, ctx.Writer)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "error generating HTML: %v", err)

			return
		}

		ctx.Header("Content-Type", "text/html")
		ctx.Status(http.StatusOK)
	})
}

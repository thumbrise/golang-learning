package observability

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	otelhttp "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/otel"
)

type OTELMiddleware struct {
	recorder *OtelRecorder
}

func NewOTELMiddleware(recorder *OtelRecorder) *OTELMiddleware {
	return &OTELMiddleware{recorder: recorder}
}

func (m *OTELMiddleware) Handler(_ context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		spanName := fmt.Sprintf("%s %s", c.Request.Method, c.FullPath())

		ctx, span := otelhttp.Tracer.Start(c.Request.Context(), spanName)
		m.recorder.Start(ctx, c, span)
		c.Request = c.Request.WithContext(ctx)

		start := time.Now()

		c.Next()

		duration := time.Since(start)

		m.recorder.End(ctx, c, span, duration)
		span.End()
	}
}

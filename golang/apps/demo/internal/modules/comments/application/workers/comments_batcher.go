package workers

import (
	"context"
	"log/slog"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/tracer"
	"golang.org/x/time/rate"
)

type CommentsBatcher struct {
	logger         *slog.Logger
	tracerProvider *tracer.Provider
}

func NewCommentsBatcher(logger *slog.Logger, tracerProvider *tracer.Provider) *CommentsBatcher {
	return &CommentsBatcher{logger: logger, tracerProvider: tracerProvider}
}

func (b *CommentsBatcher) Run(ctx context.Context) error {
	// TODO: Вынести лимит в конфиг
	// TODO: Добавить возможность настройки количества горутин. Обязательно с shared limiter

	// temp
	limiter := rate.NewLimiter(rate.Every(3*time.Second), 1)
	trace := b.tracerProvider.Tracer()
	tempResult := 1
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			err := limiter.Wait(ctx)
			if err != nil {
				return ctx.Err()
			}
			ctx, span := trace.Start(ctx, "worker comments batcher")

			b.logger.InfoContext(ctx, "Processing comments",
				slog.Int("comments", tempResult),
			)
			tempResult++

			span.End()
		}
	}
}

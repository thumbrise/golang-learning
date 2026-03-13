package workers

import (
	"context"
	"log/slog"
	"time"

	"golang.org/x/time/rate"
)

type CommentsBatcher struct {
	logger *slog.Logger
}

func NewCommentsBatcher(logger *slog.Logger) *CommentsBatcher {
	return &CommentsBatcher{logger: logger}
}

func (b *CommentsBatcher) Start(ctx context.Context) error {
	// TODO: Вынести лимит в конфиг
	// TODO: Добавить возможность настройки количества горутин. Обязательно с shared limiter

	// temp
	limiter := rate.NewLimiter(rate.Every(3*time.Second), 1)
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
			b.logger.InfoContext(ctx, "Processing comments",
				slog.Int("comments", tempResult),
			)
			tempResult++
		}
	}
}

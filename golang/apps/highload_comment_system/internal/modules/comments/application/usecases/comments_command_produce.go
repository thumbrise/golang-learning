package usecases

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
)

type CommentsCommandProduce struct {
	logger *slog.Logger
	redis  *redis.Client
}

func NewCommentsCommandProduce(logger *slog.Logger, redis *redis.Client) *CommentsCommandProduce {
	return &CommentsCommandProduce{logger: logger, redis: redis}
}

type CommentsCommandProduceInput struct{}

type CommentsCommandProduceOutput struct{}

func (c *CommentsCommandProduce) Handle(ctx context.Context, input CommentsCommandProduceInput) (*CommentsCommandProduceOutput, error) {
	// TODO: Вынести лимит в конфиг
	// TODO: Добавить консюмера
	// TODO: Вынести общую логику стриминга в инфраструктуру
	// TODO: Добавить возможность настройки количества горутин. Обязательно с shared limiter
	limiter := rate.NewLimiter(rate.Limit(3000), 1)

	for {
		select {
		case <-ctx.Done():
			return &CommentsCommandProduceOutput{}, ctx.Err()
		default:
			err := limiter.Wait(ctx)
			if err != nil {
				continue
			}

			values := map[string]interface{}{
				"uuid":      uuid.New().String(),
				"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
				"content":   faker.Paragraph(),
			}

			result, err := c.redis.XAdd(ctx, &redis.XAddArgs{
				Stream:     "comments_unprocessed",
				NoMkStream: false,
				ID:         "*",
				Values:     values,
			}).Result()
			if err != nil {
				c.logger.Error("XAdd failed",
					slog.String("error", err.Error()),
					slog.Any("values", values),
				)
			} else {
				c.logger.Debug("XAdd succeeded",
					slog.String("result", result),
					slog.Any("values", values),
				)
			}
		}
	}
}

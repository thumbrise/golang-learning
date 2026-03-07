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

type CommentsCommandProduceInput struct {
	Comments []string // получаем комментарии из входных данных
}

type CommentsCommandProduceOutput struct{}

func (c *CommentsCommandProduce) Handle(ctx context.Context, input CommentsCommandProduceInput) (*CommentsCommandProduceOutput, error) {
	// TODO: Вынести лимит в конфиг
	// TODO: Добавить консюмера
	// TODO: Вынести общую логику стриминга в инфраструктуру
	// TODO: Добавить возможность настройки количества горутин. Обязательно с shared limiter

	limiter := rate.NewLimiter(rate.Limit(3000), 1)

	inputComments := input.Comments // получаем комментарии из входных данных

	for _, comment := range inputComments {
		values := map[string]interface{}{
			"uuid":      uuid.New().String(),
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
			"content":   faker.Paragraph(),
		}

		tx, err := c.redis.TxPipelined(ctx)
		if err != nil {
			return &CommentsCommandProduceOutput{}, err
		}
		defer tx.Close()

		result, err := tx.XAdd(ctx, &redis.XAddArgs{
			Stream:     "comments_unprocessed",
			NoMkStream: false,
			ID:         "*",
			Values:     values,
		}).Result()
		if err != nil {
			return &CommentsCommandProduceOutput{}, err
		}

		cacheKey := "comments:" + uuid.New().String()
		err = tx.HSet(ctx, cacheKey, values).Err()
		if err != nil {
			return &CommentsCommandProduceOutput{}, err
		}

		// Commit transaction
		_, err = tx.Exec(ctx)
		if err != nil {
			return &CommentsCommandProduceOutput{}, err
		}

		c.logger.DebugContext(ctx, "Comment saved successfully",
			slog.String("result", result),
			slog.Any("values", values),
		)

		return &CommentsCommandProduceOutput{}, nil
	}
}

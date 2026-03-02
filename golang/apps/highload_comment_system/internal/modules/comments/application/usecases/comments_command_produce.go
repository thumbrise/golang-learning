package usecases

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type CommentsCommandProduce struct {
	logger *slog.Logger
	redis  *redis.Client
}

func NewCommentsCommandProduce(logger *slog.Logger, redis *redis.Client) *CommentsCommandProduce {
	return &CommentsCommandProduce{logger: logger, redis: redis}
}

type CommentsCommandProduceInput struct {
}

type CommentsCommandProduceOutput struct {
}

func (c *CommentsCommandProduce) Handle(ctx context.Context, input CommentsCommandProduceInput) (*CommentsCommandProduceOutput, error) {
	for {
		select {
		case <-ctx.Done():
			return &CommentsCommandProduceOutput{}, ctx.Err()
		default:
			v := redis.KeyValue{
				Key:   "uuid",
				Value: uuid.New().String(),
			}
			c.redis.XAdd(ctx, &redis.XAddArgs{
				Stream:     "comments_unprocessed",
				NoMkStream: false,
				Limit:      0,
				Mode:       "",
				ID:         "*",
				Values:     v,
			})
			time.Sleep(1 * time.Second)
		}
	}
}

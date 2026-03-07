package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ErrFailedToSaveComments = errors.New("failed to save comments")

type CommentsCommandPublish struct {
	logger *slog.Logger
	redis  *redis.Client
}

func NewCommentsCommandProduce(logger *slog.Logger, redis *redis.Client) *CommentsCommandPublish {
	return &CommentsCommandPublish{logger: logger, redis: redis}
}

type CommentsCommandPublishInput struct {
	UserUUID string
	PostUUID string
	Content  string
}

type CommentsCommandPublishOutput struct {
	Comment *Comment
}

const stream = "comments_unprocessed"

type Comment struct {
	UUID      string    `json:"uuid,omitempty"`
	PostUUID  string    `json:"postUUID,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *CommentsCommandPublish) Handle(ctx context.Context, input CommentsCommandPublishInput) (*CommentsCommandPublishOutput, error) {
	// TODO: Вынести лимит в конфиг
	// TODO: Добавить консюмера
	// TODO: Вынести общую логику стриминга в инфраструктуру
	// TODO: Добавить возможность настройки количества горутин. Обязательно с shared limiter
	comment := &Comment{
		UUID:      uuid.New().String(),
		PostUUID:  input.PostUUID,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}
	// values := map[string]interface{}{
	//	"uuid":      "",
	//	"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	//	"post_uuid": input.PostUUID,
	//	"content":   input.Content,
	//}
	var err error

	values, err := json.Marshal(comment)
	if err != nil {
		c.logger.ErrorContext(ctx, "Error marshalling comment",
			slog.String("error", err.Error()),
			slog.Any("comment", comment),
		)

		return nil, err
	}

	_, err = c.redis.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		var err error

		_, err = c.redis.XAdd(ctx, &redis.XAddArgs{
			Stream: stream,
			ID:     "*",
			Values: values,
		}).Result()
		if err != nil {
			return fmt.Errorf("redis xAdd error: %w", err)
		}

		cacheKey := "comments:" + uuid.New().String()

		err = c.redis.HSet(ctx, cacheKey, comment).Err()
		if err != nil {
			return fmt.Errorf("redis hSet error: %w", err)
		}

		return nil
	})
	if err != nil {
		c.logger.ErrorContext(ctx, "TxPipelined error",
			slog.String("error", err.Error()),
			slog.Any("comment", comment),
		)

		return nil, ErrFailedToSaveComments
	}

	return &CommentsCommandPublishOutput{
		Comment: comment,
	}, nil
}

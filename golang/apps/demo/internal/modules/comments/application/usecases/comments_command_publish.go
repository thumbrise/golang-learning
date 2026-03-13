package usecases

import (
	"context"
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

func NewCommentsCommandPublish(logger *slog.Logger, redis *redis.Client) *CommentsCommandPublish {
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
	UUID      string
	PostUUID  string
	Content   string
	CreatedAt time.Time
}

func (c *CommentsCommandPublish) Handle(ctx context.Context, input CommentsCommandPublishInput) (*CommentsCommandPublishOutput, error) {
	// TODO: Вынести общую логику стриминга в инфраструктуру
	// TODO: Добавить валидацию
	comment := &Comment{
		UUID:      uuid.New().String(),
		PostUUID:  input.PostUUID,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}
	redisRecord := map[string]interface{}{
		"uuid":       comment.UUID,
		"post_uuid":  comment.PostUUID,
		"content":    comment.Content,
		"created_at": comment.CreatedAt.Unix(),
	}

	_, err := c.redis.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		var err error

		_, err = c.redis.XAdd(ctx, &redis.XAddArgs{
			Stream: stream,
			ID:     "*",
			Values: redisRecord,
		}).Result()
		if err != nil {
			return fmt.Errorf("redis xAdd error: %w", err)
		}

		cacheKey := "comments:" + uuid.New().String()

		err = c.redis.HSet(ctx, cacheKey, redisRecord).Err()
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

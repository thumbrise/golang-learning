package workers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/tracer"
	"golang.org/x/time/rate"
)

// TODO: Вынести в конфиг
const (
	idNext           = ">"
	idPending        = "0"
	stream           = "comments_unprocessed"
	group            = "test"
	workerName       = "worker comments batcher"
	rpsDatabaseLimit = 3000
)

// CommentsBatcher
//
// TODO: refactoring
type CommentsBatcher struct {
	logger         *slog.Logger
	tracerProvider *tracer.Provider
	redisClient    *redis.Client
	handledPending bool
}

func NewCommentsBatcher(logger *slog.Logger, redisClient *redis.Client, tracerProvider *tracer.Provider) *CommentsBatcher {
	return &CommentsBatcher{logger: logger, redisClient: redisClient, tracerProvider: tracerProvider}
}

func (b *CommentsBatcher) Run(ctx context.Context) error {
	// TODO: Добавить возможность настройки количества горутин. Обязательно с shared limiter
	err := b.ensureGroupExists(ctx)
	if err != nil {
		return err
	}

	return b.loop(ctx)
}

func (b *CommentsBatcher) loop(ctx context.Context) error {
	limiter := rate.NewLimiter(rate.Limit(rpsDatabaseLimit), 1)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			err := limiter.Wait(ctx)
			if err != nil {
				return ctx.Err()
			}

			ctx, span := b.tracerProvider.Tracer().Start(ctx, fmt.Sprintf("%s - %s", workerName, "process"))

			processed, err := b.process(ctx)
			if err != nil {
				b.logger.ErrorContext(ctx, "process()", slog.String("error", err.Error()))
				span.RecordError(err)
			}

			b.logger.InfoContext(ctx, "processed comments batcher", slog.Any("processed", processed))
			span.End()
		}
	}
}

func (b *CommentsBatcher) process(ctx context.Context) (int, error) {
	if !b.handledPending {
		processed, err := b.readBatch(ctx, idPending)
		if err != nil {
			return 0, err
		}

		if processed == 0 {
			b.handledPending = true
		}
	}

	return b.readBatch(ctx, idNext)
}

func (b *CommentsBatcher) ensureGroupExists(ctx context.Context) error {
	ctx, span := b.tracerProvider.Tracer().Start(ctx, fmt.Sprintf("%s - %s", workerName, "process"))
	defer span.End()

	res, err := b.redisClient.XGroupCreate(ctx, stream, group, "0").Result()
	if err != nil {
		if redis.HasErrorPrefix(err, "BUSYGROUP") {
			b.logger.DebugContext(ctx, "redisClient.XGroupCreate - group already exists", slog.String("res", res))

			return nil
		}

		b.logger.ErrorContext(ctx, "", slog.String("error", err.Error()))
		span.RecordError(err)

		return err
	}

	b.logger.DebugContext(ctx, "redisClient.XGroupCreate", slog.String("res", res))

	return nil
}

func (b *CommentsBatcher) readBatch(ctx context.Context, id string) (int, error) {
	streams, err := b.redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: "1",
		Streams:  []string{stream, id},
		Count:    100,
		Block:    0,
		NoAck:    false,
		Claim:    0,
	}).Result()
	if err != nil {
		return 0, err
	}

	result := 0

	for _, stream := range streams {
		for _, message := range stream.Messages {
			b.logger.DebugContext(ctx, "Processing comments batcher",
				slog.Any("message", message),
			)
			b.redisClient.XAckDel(ctx, stream.Stream, group, "ACKED", message.ID)

			result++
		}
	}

	return result, nil
}

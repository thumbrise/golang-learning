package workers

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/tracer"
	"golang.org/x/time/rate"
)

type CommentsBatcher struct {
	logger         *slog.Logger
	tracerProvider *tracer.Provider
	redisClient    *redis.Client
}

func NewCommentsBatcher(logger *slog.Logger, redisClient *redis.Client, tracerProvider *tracer.Provider) *CommentsBatcher {
	return &CommentsBatcher{logger: logger, redisClient: redisClient, tracerProvider: tracerProvider}
}

func (b *CommentsBatcher) Run(ctx context.Context) error {
	// TODO: Вынести лимит в конфиг
	// TODO: Добавить возможность настройки количества горутин. Обязательно с shared limiter

	// temp
	limiter := rate.NewLimiter(rate.Limit(3000), 1)
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

			if err := b.process(ctx); err != nil {
				b.logger.ErrorContext(ctx, "Error while processing comments batcher", slog.String("error", err.Error()))
				span.RecordError(err)
			}

			tempResult++

			span.End()
		}
	}
}

func (b *CommentsBatcher) process(ctx context.Context) error {
	// TODO: Вынести в конфиг
	// 	Сделать предсоздание группы
	//  Сделать дочитывание при первом поднятии и при падениях
	const group = "test"
	result, err := b.redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: "1",
		Streams:  []string{"comments_unprocessed", ">"},
		Count:    3,
		Block:    0,
		NoAck:    false,
		Claim:    0,
	}).Result()
	if err != nil {
		return err
	}

	for _, stream := range result {
		for _, message := range stream.Messages {
			b.logger.InfoContext(ctx, "Processing comments batcher",
				slog.Any("message", message),
			)
			b.redisClient.XAck(ctx, stream.Stream, group, message.ID)
		}
	}

	return nil
}

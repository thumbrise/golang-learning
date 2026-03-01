package otp

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/dal"
)

type OTPRedisRepository struct {
	redis *redis.Client
}

func NewOTPRedisRepository(redis *redis.Client) *OTPRedisRepository {
	return &OTPRedisRepository{redis: redis}
}

func (r *OTPRedisRepository) Create(ctx context.Context, otp *OTP) error {
	key := r.key(otp.Code, otp.UserID)

	resp := r.redis.Set(ctx, key, true, time.Until(otp.ExpiresAt))
	if err := resp.Err(); err != nil {
		return fmt.Errorf("%w: %w", dal.ErrFailedQuery, err)
	}

	return nil
}

func (r *OTPRedisRepository) ExistsValid(ctx context.Context, userID int, code string) (bool, error) {
	key := r.key(code, userID)

	resp, err := r.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("%w: %w", dal.ErrFailedQuery, err)
	}

	return resp > 0, nil
}

func (r *OTPRedisRepository) key(code string, userID int) string {
	return fmt.Sprintf("otp:%s:user:%d", code, userID)
}

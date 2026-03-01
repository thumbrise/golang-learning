package components

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/config"
)

const (
	defaultDialTimeout     = 2 * time.Second        // Быстрое установление соединения
	defaultReadTimeout     = 100 * time.Millisecond // Критично для высокой нагрузки
	defaultWriteTimeout    = 100 * time.Millisecond
	defaultPoolTimeout     = 1 * time.Second // Время ожидания свободного соединения
	defaultIdleTimeout     = 5 * time.Minute // Соединения живут дольше
	defaultPoolSize        = 200             // Важно для 2000 RPS
	defaultMinIdleConns    = 50              // Держим пул готовых соединений
	defaultMaxRetries      = 1               // Минимальные ретраи для снижения latency
	defaultMaxRetryBackoff = 100 * time.Millisecond
)

var ErrUrlParse = errors.New("parse url error")

func NewRedisClient(cfg config.Redis) (*redis.Client, error) {
	addr := cfg.URL
	if strings.Contains(addr, "://") {
		u, err := url.Parse(cfg.URL)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrUrlParse, err)
		}

		addr = u.Host
	}

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,

		DialTimeout:     defaultDialTimeout,
		ReadTimeout:     defaultReadTimeout,
		WriteTimeout:    defaultWriteTimeout,
		PoolTimeout:     defaultPoolTimeout,
		ConnMaxIdleTime: defaultIdleTimeout,
		PoolSize:        defaultPoolSize,
		MinIdleConns:    defaultMinIdleConns,
		MaxRetries:      defaultMaxRetries,
		MaxRetryBackoff: defaultMaxRetryBackoff,

		// Для мониторинга и отладки
		PoolFIFO:        true,             // FIFO очередь соединений (лучше для однородной нагрузки)
		ConnMaxLifetime: 30 * time.Minute, // Переподключаемся периодически
	}), nil
}

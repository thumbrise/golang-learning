package components

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thumbrise/demo/golang-demo/internal/config"
	"go.uber.org/fx"
)

type DB struct {
	pool   *pgxpool.Pool
	config config.DB
}

func NewDB(lc fx.Lifecycle, config config.DB) *DB {
	db := &DB{config: config}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			db.pool.Close()
			return nil
		},
	})
	return db
}
func (db *DB) Connect(ctx context.Context) error {
	if db.pool != nil {
		log.Fatal(ErrPoolAlreadyOpen)
	}

	pool, err := pgxpool.New(ctx, db.dsn())
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantBeCreated, err)
	}

	db.pool = pool
	slog.Debug("DB connected")
	return nil
}

func (db *DB) dsn() string {
	hostPort := net.JoinHostPort(db.config.Host, db.config.Port)

	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		db.config.Username,
		db.config.Password,
		hostPort,
		db.config.Database,
	)
}

var (
	ErrPoolAlreadyOpen = errors.New("pool already open")
	ErrPoolIsClosed    = errors.New("pool is closed")
	ErrCantBeCreated   = errors.New("pool can't be created")
)

func (db *DB) Pool() *pgxpool.Pool {
	if db.pool == nil {
		panic(ErrPoolIsClosed)
	}

	return db.pool
}

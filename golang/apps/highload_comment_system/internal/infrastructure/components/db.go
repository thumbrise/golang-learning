package components

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thumbrise/demo/golang-demo/internal/config"
)

type DB struct {
	pool   *pgxpool.Pool
	config config.DB
}

func MustConnect(ctx context.Context, config config.DB) *DB {
	db := &DB{config: config}
	if db.pool != nil {
		log.Fatal(ErrPoolAlreadyOpen)
	}

	pool, err := pgxpool.New(ctx, db.dsn())
	if err != nil {
		log.Fatalf("%s: %s", ErrCantBeCreated, err)
	}

	db.pool = pool

	return db
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
		log.Fatal(ErrPoolIsClosed)
	}

	return db.pool
}

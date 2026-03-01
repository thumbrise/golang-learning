package database

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/components"
)

type Bootloader struct {
	db *components.DB
}

func NewBootloader(db *components.DB) *Bootloader {
	return &Bootloader{db: db}
}

func (b *Bootloader) Boot(ctx context.Context) error {
	return b.db.Connect(ctx)
}

func (b *Bootloader) Shutdown(context.Context) error {
	b.db.Pool().Close()

	return nil
}

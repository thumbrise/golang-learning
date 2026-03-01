package redis

import (
	"context"
)

type Bootloader struct{}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}

func NewBootloader() *Bootloader {
	return &Bootloader{}
}

func (b *Bootloader) Boot(ctx context.Context) error {
	return nil
}

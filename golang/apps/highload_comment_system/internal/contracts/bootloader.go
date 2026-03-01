package contracts

import "context"

type Bootloader interface {
	Boot(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

package database

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(
		NewDB,
		NewConfig,
	),
	fx.Invoke(func(lc fx.Lifecycle, db *DB) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return db.Connect(ctx)
			},
			OnStop: func(ctx context.Context) error {
				db.Pool().Close()

				return nil
			},
		})
	}),
)

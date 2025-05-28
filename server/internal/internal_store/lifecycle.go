package internal_store

import (
	"context"
	"github.com/dgraph-io/badger/v4"
	"go.etcd.io/bbolt"
	"go.uber.org/fx"
)

var lifecycle = fx.Module("internal_store_lifecycle",
	fx.Invoke(badgerLifecycle, bblotLifecycle),
)

func badgerLifecycle(lc fx.Lifecycle, db *badger.DB) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
}

func bblotLifecycle(lc fx.Lifecycle, db *bbolt.DB) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
}

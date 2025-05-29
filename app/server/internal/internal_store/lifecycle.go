package internal_store

import (
	"context"
	"github.com/dgraph-io/badger/v4"
	"go.etcd.io/bbolt"
	"go.uber.org/fx"
)

var lifecycle = fx.Module("internal_store_lifecycle",
	fx.Invoke(diskBadgerLifecycle, memoryBadgerLifecycle, bblotLifecycle),
)

type DiskBadgerLifecycleOption struct {
	fx.In
	DB        *badger.DB `name:"disk"`
	Lifecycle fx.Lifecycle
}

func diskBadgerLifecycle(opt DiskBadgerLifecycleOption) {
	opt.Lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return opt.DB.Close()
		},
	})
}

type MemoryBadgerLifecycleOption struct {
	fx.In
	DB        *badger.DB `name:"memory"`
	Lifecycle fx.Lifecycle
}

func memoryBadgerLifecycle(opt MemoryBadgerLifecycleOption) {
	opt.Lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return opt.DB.Close()
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

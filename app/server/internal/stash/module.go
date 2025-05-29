package stash

import (
	"github.com/dgraph-io/ristretto/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("stash", fx.Provide(newMemStash))

func newMemStash() (*ristretto.Cache[string, string], error) {
	return ristretto.NewCache(&ristretto.Config[string, string]{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
}

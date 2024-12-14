package cache

import (
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	ristrettostore "github.com/eko/gocache/store/ristretto/v4"
	"go.uber.org/fx"
)

var Module = fx.Module("cache", fx.Provide(globalCache))

func globalCache() *cache.Cache[string] {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := ristrettostore.NewRistretto(ristrettoCache)

	cacheManager := cache.New[string](ristrettoStore)
	return cacheManager
}

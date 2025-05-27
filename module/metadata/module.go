package metadata

import (
	"go.etcd.io/bbolt"
	"go.uber.org/fx"
)

var Module = fx.Module("metadata", fx.Provide(newBbolt))

func newBbolt() (*bbolt.DB, error) {
	return bbolt.Open("my.db", 0600, nil)
}

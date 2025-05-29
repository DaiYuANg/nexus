package auth

import (
	"go.etcd.io/bbolt"
	"go.uber.org/fx"
)

var Module = fx.Module("auth",
	fx.Provide(
		fx.Annotate(
			newAuthStorage,
			fx.ResultTags(`name:"authStore"`),
		),
	),
)

func newAuthStorage() (*bbolt.DB, error) {
	return bbolt.Open("auth.db", 0600, nil)
}

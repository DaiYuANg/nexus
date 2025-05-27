package namespace

import (
	"go.uber.org/fx"
)

var Module = fx.Module("namespace",
	fx.Provide(
		NewService,
	),
)

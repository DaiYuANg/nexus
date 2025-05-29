package bucket

import (
	"go.uber.org/fx"
)

var Module = fx.Module("bucket",
	fx.Provide(
		NewService,
	),
)

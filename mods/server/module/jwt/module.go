package jwt

import (
	"go.uber.org/fx"
)

var Module = fx.Module("jwt",
	fx.Provide(
		fx.Annotate(
			getSigningKey,
			fx.ResultTags(`name:"jwtKey"`),
		),
	),
)

package runtime

import (
	"go.uber.org/fx"
	"os"
)

var Module = fx.Module("runtime",
	fx.Provide(
		fx.Annotate(
			pwd,
			fx.ResultTags(`name:"pwd"`),
		),
	),
)

func pwd() (string, error) {
	return os.Getwd()
}

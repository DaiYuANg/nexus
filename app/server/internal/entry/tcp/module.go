package tcp

import (
	"go.uber.org/fx"
)

var Module = fx.Module("tcp", fx.Provide(newStorixServer), fx.Invoke(startServer))

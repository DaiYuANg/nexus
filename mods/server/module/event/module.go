package event

import (
	"github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
)

var Module = fx.Module("event_bus_module",
	fx.Provide(goeventbus.NewEventBus),
)

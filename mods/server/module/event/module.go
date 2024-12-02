package event

import (
	"github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
	"nexus/internal/event"
)

var Module = fx.Module("event_bus_module",
	fx.Provide(goeventbus.NewEventBus),
	fx.Invoke(event.UserRegisteredConsumer),
)

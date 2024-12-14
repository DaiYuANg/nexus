package server

import "go.uber.org/fx"

var middleware = fx.Module("middleware",
	fx.Invoke(
		registerCompress,
		registerEnvvars,
		registerMonitor,
		registerRequestId,
		registerJwt,
		registerPprof,
		registerPrometheus,
		registerLogger,
		registerZap,
		registerRecover,
		registerUaRecorder,
		registerSwagger,
	),
)

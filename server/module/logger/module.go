package logger

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Module = fx.Module("logger_module", fx.Provide(newLogger, sugaredLogger), fx.Invoke(deferLogger))

func newLogger() *zap.Logger {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	logger := zap.New(core)
	return logger
}

func sugaredLogger(log *zap.Logger) *zap.SugaredLogger {
	return log.Sugar()
}

func deferLogger(lc fx.Lifecycle, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})
}

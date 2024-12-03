package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"nexus/internal/conf"
	"os"
)

var Module = fx.Module("logger", fx.Provide(newLogger))

type Params struct {
	fx.In
	Config *conf.LoggingConfig
}

func newLogger(params Params) *zap.Logger {
	loggerConfig := params.Config
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   loggerConfig.Path,
		MaxSize:    3, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeLevel = lowerCaseLevelEncoder
	productionCfg.StacktraceKey = "stack"

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	developmentCfg.CallerKey = "caller"
	developmentCfg.EncodeCaller = zapcore.FullCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)

	fileEncoder := zapcore.NewJSONEncoder(productionCfg)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
	return zap.New(core)
}

func lowerCaseLevelEncoder(
	level zapcore.Level,
	enc zapcore.PrimitiveArrayEncoder,
) {
	if level == zap.PanicLevel || level == zap.DPanicLevel {
		enc.AppendString("error")
		return
	}

	zapcore.LowercaseLevelEncoder(level, enc)
}

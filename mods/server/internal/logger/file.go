package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func file() zapcore.Encoder {
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeLevel = lowerCaseLevelEncoder
	productionCfg.StacktraceKey = "stack"
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)
	return fileEncoder
}

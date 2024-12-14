package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func consoleConfig() zapcore.Encoder {
	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	developmentCfg.CallerKey = "caller"
	developmentCfg.EncodeCaller = zapcore.FullCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)

	return consoleEncoder
}

func consoleCore() {

}

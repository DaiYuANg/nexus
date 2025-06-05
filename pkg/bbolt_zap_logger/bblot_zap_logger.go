package bblot_zap_logger

import (
	"fmt"
	"go.uber.org/zap"
)

const bblotLogPrefix = "[internal_store:bblot]"

type BblotZapLogger struct {
	Logger *zap.SugaredLogger
}

func (b BblotZapLogger) Debug(v ...interface{}) {
	b.Logger.Debug(v...)
}

func (b BblotZapLogger) Debugf(format string, v ...interface{}) {
	b.Logger.Debugf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b BblotZapLogger) Error(v ...interface{}) {
	b.Logger.Error(v...)
}

func (b BblotZapLogger) Errorf(format string, v ...interface{}) {
	b.Logger.Errorf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b BblotZapLogger) Info(v ...interface{}) {
	b.Logger.Info(v...)
}

func (b BblotZapLogger) Infof(format string, v ...interface{}) {
	b.Logger.Infof(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b BblotZapLogger) Warning(v ...interface{}) {
	b.Logger.Warn(v...)
}

func (b BblotZapLogger) Warningf(format string, v ...interface{}) {
	b.Logger.Warnf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b BblotZapLogger) Fatal(v ...interface{}) {
	b.Logger.Fatal(v...)
}

func (b BblotZapLogger) Fatalf(format string, v ...interface{}) {
	b.Logger.Fatalf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b BblotZapLogger) Panic(v ...interface{}) {
	b.Logger.Panic(v...)
}

func (b BblotZapLogger) Panicf(format string, v ...interface{}) {
	b.Logger.Panicf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

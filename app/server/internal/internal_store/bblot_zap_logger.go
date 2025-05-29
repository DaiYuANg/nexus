package internal_store

import (
	"fmt"
	"go.uber.org/zap"
)

const bblotLogPrefix = "[internal_store:bblot]"

type bblotZapLogger struct {
	suger *zap.SugaredLogger
}

func (b bblotZapLogger) Debug(v ...interface{}) {
	b.suger.Debug(v...)
}

func (b bblotZapLogger) Debugf(format string, v ...interface{}) {
	b.suger.Debugf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b bblotZapLogger) Error(v ...interface{}) {
	b.suger.Error(v...)
}

func (b bblotZapLogger) Errorf(format string, v ...interface{}) {
	b.suger.Errorf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b bblotZapLogger) Info(v ...interface{}) {
	b.suger.Info(v...)
}

func (b bblotZapLogger) Infof(format string, v ...interface{}) {
	b.suger.Infof(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b bblotZapLogger) Warning(v ...interface{}) {
	b.suger.Warn(v...)
}

func (b bblotZapLogger) Warningf(format string, v ...interface{}) {
	b.suger.Warnf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b bblotZapLogger) Fatal(v ...interface{}) {
	b.suger.Fatal(v...)
}

func (b bblotZapLogger) Fatalf(format string, v ...interface{}) {
	b.suger.Fatalf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

func (b bblotZapLogger) Panic(v ...interface{}) {
	b.suger.Panic(v...)
}

func (b bblotZapLogger) Panicf(format string, v ...interface{}) {
	b.suger.Panicf(fmt.Sprintf("%s %s", bblotLogPrefix, format), v...)
}

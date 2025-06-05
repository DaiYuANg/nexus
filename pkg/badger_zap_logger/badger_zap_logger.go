package badger_zap_logger

import (
	"fmt"
	"go.uber.org/zap"
)

const badgerLogPrefix = "[internal_store:badger] "

type BadgerZapLogger struct {
	Logger *zap.SugaredLogger
}

func (l *BadgerZapLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

func (l *BadgerZapLogger) Warningf(format string, args ...interface{}) {
	l.Logger.Warnf(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

func (l *BadgerZapLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

func (l *BadgerZapLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

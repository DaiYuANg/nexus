package internal_store

import (
	"fmt"
	"go.uber.org/zap"
)

const badgerLogPrefix = "[internal_store:badger] "

type badgerLogger struct {
	sugar *zap.SugaredLogger
}

func (l *badgerLogger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

func (l *badgerLogger) Warningf(format string, args ...interface{}) {
	l.sugar.Warnf(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

func (l *badgerLogger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

func (l *badgerLogger) Debugf(format string, args ...interface{}) {
	l.sugar.Debugf(fmt.Sprintf("%s %s", badgerLogPrefix, format), args...)
}

package internal_store

import "go.uber.org/zap"

type badgerLogger struct {
	sugar *zap.SugaredLogger
}

func (l *badgerLogger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(format, args...)
}

func (l *badgerLogger) Warningf(format string, args ...interface{}) {
	l.sugar.Warnf(format, args...)
}

func (l *badgerLogger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(format, args...)
}

func (l *badgerLogger) Debugf(format string, args ...interface{}) {
	l.sugar.Debugf(format, args...)
}

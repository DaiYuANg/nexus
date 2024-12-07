package local

import (
	"go.uber.org/zap"
)

type badgerZapLoggerAdapter struct {
	zlogger *zap.SugaredLogger
}

func (l *badgerZapLoggerAdapter) Errorf(s string, i ...interface{}) {
	l.zlogger.Errorf(s, i...)
}

func (l *badgerZapLoggerAdapter) Warningf(s string, i ...interface{}) {
	l.zlogger.Warnf(s, i...)
}

func (l *badgerZapLoggerAdapter) Infof(s string, i ...interface{}) {
	l.zlogger.Infof(s, i...)
}

func (l *badgerZapLoggerAdapter) Debugf(s string, i ...interface{}) {
	l.zlogger.Debugf(s, i...)
}

func newBadgerZapLoggerAdapter(zapLogger *zap.SugaredLogger) *badgerZapLoggerAdapter {
	return &badgerZapLoggerAdapter{zapLogger}
}

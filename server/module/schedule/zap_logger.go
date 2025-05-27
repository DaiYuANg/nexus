package schedule

import "go.uber.org/zap"

type schedulerZapLogger struct {
	zapLogger *zap.SugaredLogger
}

func (s schedulerZapLogger) Debug(msg string, args ...any) {
	s.zapLogger.Debugf(msg, args...)
}

func (s schedulerZapLogger) Error(msg string, args ...any) {
	s.zapLogger.Errorf(msg, args...)
}

func (s schedulerZapLogger) Info(msg string, args ...any) {
	s.zapLogger.Infof(msg, args...)
}

func (s schedulerZapLogger) Warn(msg string, args ...any) {
	s.zapLogger.Warnf(msg, args...)
}

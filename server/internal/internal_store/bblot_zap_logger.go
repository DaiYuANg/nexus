package internal_store

import "go.uber.org/zap"

type bblotZapLogger struct {
	suger *zap.SugaredLogger
}

func (b bblotZapLogger) Debug(v ...interface{}) {
	b.suger.Debug(v...)
}

func (b bblotZapLogger) Debugf(format string, v ...interface{}) {
	b.suger.Debugf(format, v...)
}

func (b bblotZapLogger) Error(v ...interface{}) {
	b.suger.Error(v...)
}

func (b bblotZapLogger) Errorf(format string, v ...interface{}) {
	b.suger.Errorf(format, v...)
}

func (b bblotZapLogger) Info(v ...interface{}) {
	b.suger.Info(v...)
}

func (b bblotZapLogger) Infof(format string, v ...interface{}) {
	b.suger.Infof(format, v...)
}

func (b bblotZapLogger) Warning(v ...interface{}) {
	b.suger.Warn(v...)
}

func (b bblotZapLogger) Warningf(format string, v ...interface{}) {
	b.suger.Warnf(format, v...)
}

func (b bblotZapLogger) Fatal(v ...interface{}) {
	b.suger.Fatal(v...)
}

func (b bblotZapLogger) Fatalf(format string, v ...interface{}) {
	b.suger.Fatalf(format, v...)
}

func (b bblotZapLogger) Panic(v ...interface{}) {
	b.suger.Panic(v...)
}

func (b bblotZapLogger) Panicf(format string, v ...interface{}) {
	b.suger.Panicf(format, v...)
}

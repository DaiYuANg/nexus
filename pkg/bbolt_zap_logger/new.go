package bblot_zap_logger

import (
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func NewBblotDb(path string, logger *zap.Logger, options *bbolt.Options) (*bbolt.DB, error) {
	zapLogger := BblotZapLogger{Logger: logger.Sugar()}
	options.Logger = zapLogger
	return bbolt.Open(path, 0666, options)
}

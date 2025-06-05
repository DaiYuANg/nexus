package badger_zap_logger

import (
	"github.com/dgraph-io/badger/v4"
	"go.uber.org/zap"
)

func NewBadger(logger *zap.Logger, options badger.Options) (*badger.DB, error) {
	zapLogger := &BadgerZapLogger{Logger: logger.Sugar()}
	options.
		WithLogger(zapLogger)
	return badger.Open(options)
}

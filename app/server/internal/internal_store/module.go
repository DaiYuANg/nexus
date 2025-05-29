package internal_store

import (
	"github.com/adrg/xdg"
	"github.com/dgraph-io/badger/v4"
	"github.com/influxdata/influxdb/pkg/snowflake"
	"go.etcd.io/bbolt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"path"
)

var Module = fx.Module("internal_store",
	fx.Provide(
		storePath,
		snowflakeId,
		newBblot,
		fx.Annotate(
			newDiskBadger,
			fx.ResultTags(`name:"disk"`),
		),
		fx.Annotate(
			newMemBadger,
			fx.ResultTags(`name:"memory"`),
		),
	),
	lifecycle,
)

const (
	dataDir = "maxio_data"
)

func storePath(logger *zap.SugaredLogger) (string, error) {
	dataPath := path.Join(xdg.DataHome, dataDir)
	_, err := os.Stat(dataPath)
	if err != nil {
		err := os.MkdirAll(dataPath, 0755)
		if err != nil {
			return "", err
		}
	}
	logger.Infof("data path:%s", dataPath)
	return dataPath, nil
}

func snowflakeId() *snowflake.Generator {
	return snowflake.New(0)
}

func newMemBadger(suger *zap.SugaredLogger) (*badger.DB, error) {
	logger := &badgerLogger{suger}
	options := badger.
		DefaultOptions("").
		WithZSTDCompressionLevel(1).
		WithLoggingLevel(badger.DEBUG).
		WithLogger(logger).
		WithInMemory(true)

	return badger.Open(options)
}

func newDiskBadger(storePath string, suger *zap.SugaredLogger) (*badger.DB, error) {
	badgerPath := path.Join(storePath, "metadata")
	logger := &badgerLogger{suger}
	options := badger.
		DefaultOptions(badgerPath).
		WithZSTDCompressionLevel(1).
		WithLoggingLevel(badger.DEBUG).WithLogger(logger)
	//options.Logger = logger
	return badger.Open(options)
}

func newBblot(storePath string, suger *zap.SugaredLogger) (*bbolt.DB, error) {
	bblotPath := path.Join(storePath, "storix.db")
	suger.Debugf("opening bblot: %s", bblotPath)
	logger := bblotZapLogger{suger: suger}
	options := bbolt.DefaultOptions
	options.Logger = logger
	return bbolt.Open(bblotPath, 0666, options)
}

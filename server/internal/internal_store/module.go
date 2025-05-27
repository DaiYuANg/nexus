package internal_store

import (
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
		newBadger,
	),
)

const (
	dataDir = "storix_data"
)

func storePath() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dataPath := path.Join(pwd, dataDir)
	_, err = os.Stat(dataPath)
	if err != nil {
		err := os.MkdirAll(dataPath, 0755)
		if err != nil {
			return "", err
		}
	}
	return dataPath, nil
}

func snowflakeId() *snowflake.Generator {
	return snowflake.New(0)
}

func newBadger(storePath string, suger *zap.SugaredLogger) (*badger.DB, error) {
	badgerPath := path.Join(storePath, "metadata")
	logger := &badgerLogger{suger}
	options := badger.DefaultOptions(badgerPath).WithLogger(logger).WithLoggingLevel(badger.DEBUG)
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

package storage

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/influxdata/influxdb/pkg/snowflake"
	"github.com/spf13/afero"
	"go.uber.org/fx"
)

var Module = fx.Module("storage", fx.Provide(newMemfs, newFileChunker))

func newMemfs() afero.Fs {
	return afero.NewMemMapFs()
}

type FileChuckerOption struct {
	fx.In
	DB        *badger.DB `name:"disk"`
	Generator *snowflake.Generator
}

func newFileChunker(options FileChuckerOption) *FileChunker {
	return &FileChunker{store: options.DB, idGenerator: options.Generator}
}

package storage

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/spf13/afero"
	"go.uber.org/fx"
)

var Module = fx.Module("storage", fx.Provide(newMemfs, newFileChunker))

func newMemfs() afero.Fs {
	return afero.NewMemMapFs()
}

func newFileChunker(db *badger.DB) *FileChunker {
	return &FileChunker{store: db}
}

package local

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"log"
)

type VfsConfig struct {
	BasePath string
	*zap.Logger
}

func New(config VfsConfig) *VFS {
	var localFs = afero.NewBasePathFs(afero.NewOsFs(), config.BasePath)
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return &VFS{
		config.Logger,
		watcher,
		localFs,
		db,
	}
}

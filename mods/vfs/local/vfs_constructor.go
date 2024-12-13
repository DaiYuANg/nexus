package local

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"io/fs"
	"path"
	"path/filepath"
)

type VfsConfig struct {
	BasePath      string
	IgnoreWatch   []string
	ChangeHandler ChangeHandler
	*zap.SugaredLogger
}

func New(config VfsConfig) (*VFS, error) {
	var localFs = afero.NewBasePathFs(afero.NewOsFs(), config.BasePath)
	databasePath := path.Join(config.BasePath, meta)
	var dbOption = badger.DefaultOptions(databasePath)
	dbOption.Logger = newBadgerZapLoggerAdapter(config.SugaredLogger)
	db, err := badger.Open(dbOption)
	if err != nil {
		return nil, err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = watcher.Add(config.BasePath)
	if err != nil {
		return nil, err
	}

	err = filepath.WalkDir(config.BasePath, func(path string, d fs.DirEntry, err error) error {
		if path == config.BasePath {
			return nil
		}
		if err != nil {
			return err
		}
		if d.Name() == meta {
			return nil
		}
		if lo.Contains(config.IgnoreWatch, d.Name()) {
			return nil
		}
		if d.IsDir() {
			config.Info("Watch", zap.String("path", path))
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var handler ChangeHandler
	if config.ChangeHandler != nil {
		handler = config.ChangeHandler
	} else {
		handler = func(op fsnotify.Op, filename string) {
			config.SugaredLogger.Infof("File %s, %s", op, filename)
		}
	}
	return &VFS{
		config.SugaredLogger,
		watcher,
		localFs,
		db,
		handler,
	}, nil
}

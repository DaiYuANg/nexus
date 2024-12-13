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
	// Step 1: Create the file system
	localFs := createFs(config)

	// Step 2: Open Badger DB
	db, err := createDatabase(config.BasePath, config.SugaredLogger)
	if err != nil {
		return nil, err
	}

	// Step 3: Create and configure file system watcher
	watcher, err := createWatcher(config.BasePath)
	if err != nil {
		return nil, err
	}

	// Step 4: Add directories to watcher
	err = addDirectoriesToWatcher(config.BasePath, config.IgnoreWatch, watcher, config.SugaredLogger)
	if err != nil {
		return nil, err
	}

	// Step 5: Set up a change handler
	handler := setupChangeHandler(config)

	return &VFS{
		config.SugaredLogger,
		watcher,
		localFs,
		db,
		handler,
	}, nil
}

// createFs creates and returns a new file system with the given base path.
func createFs(config VfsConfig) afero.Fs {
	return afero.NewBasePathFs(afero.NewOsFs(), config.BasePath)
}

// createDatabase initializes and opens the Badger database.
func createDatabase(basePath string, logger *zap.SugaredLogger) (*badger.DB, error) {
	databasePath := path.Join(basePath, meta)
	options := badger.DefaultOptions(databasePath)
	options.Logger = newBadgerZapLoggerAdapter(logger)
	return badger.Open(options)
}

// createWatcher initializes and returns a file system watcher.
func createWatcher(basePath string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = watcher.Add(basePath)
	if err != nil {
		return nil, err
	}
	return watcher, nil
}

// addDirectoriesToWatcher walks through the base path and adds directories to the watcher.
func addDirectoriesToWatcher(
	basePath string,
	ignoreWatch []string,
	watcher *fsnotify.Watcher,
	logger *zap.SugaredLogger,
) error {
	return filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if path == basePath || err != nil || d.Name() == meta || lo.Contains(ignoreWatch, d.Name()) {
			return nil
		}

		if d.IsDir() {
			logger.Info("Watch", zap.String("path", path))
			return watcher.Add(path)
		}
		return nil
	})
}

// setupChangeHandler sets up a default or custom change handler.
func setupChangeHandler(config VfsConfig) ChangeHandler {
	if config.ChangeHandler != nil {
		return config.ChangeHandler
	}
	return func(op fsnotify.Op, filename string) {
		config.SugaredLogger.Infof("File %s, %s", op, filename)
	}
}

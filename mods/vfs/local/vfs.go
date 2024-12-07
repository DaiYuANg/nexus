package local

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

type ChangeHandler func(op fsnotify.Op, filename string)

type VFS struct {
	*zap.SugaredLogger
	*fsnotify.Watcher
	afero.Fs
	*badger.DB
	changeHandler ChangeHandler
}

func (v *VFS) Watch() {
	for {
		select {
		case event, ok := <-v.Events:
			if !ok {
				return
			}
			v.changeHandler(event.Op, event.Name)
		case err, ok := <-v.Errors:
			v.Error("error:", zap.Error(err))
			if !ok {
				return
			}
		}
	}
}

func (v *VFS) Close() error {
	err := v.DB.Close()
	if err != nil {
		return err
	}

	err = v.Watcher.Close()
	if err != nil {
		return err
	}
	return nil
}

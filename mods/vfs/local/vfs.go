package local

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"io/fs"
	"log"
)

type VFS struct {
	*zap.Logger
	*fsnotify.Watcher
	afero.Fs
	*badger.DB
}

func (v *VFS) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Stat(name string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Watch(name string) {
	for {
		select {
		case event, ok := <-v.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			if event.Has(fsnotify.Write) {
				log.Println("modified file:", event.Name)
			}
		case err, ok := <-v.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
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

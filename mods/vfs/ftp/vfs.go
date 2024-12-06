package sftp

import (
	"go.uber.org/zap"
	"io/fs"
)

type NFS struct {
	*zap.Logger
}

func (NFS) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (NFS) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (NFS) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (NFS) Stat(name string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (NFS) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

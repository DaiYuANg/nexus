package sftp

import (
	"github.com/studio-b12/gowebdav"
	"go.uber.org/zap"
	"io/fs"
)

type Webdav struct {
	*zap.Logger
	*gowebdav.Client
}

func (Webdav) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (Webdav) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (Webdav) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (Webdav) Stat(name string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (Webdav) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

package sftp

import (
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"io/fs"
)

type Sftp struct {
	*zap.Logger
	*sftp.Client
}

func (Sftp) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (Sftp) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (Sftp) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (Sftp) Stat(name string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (Sftp) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

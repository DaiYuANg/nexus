package sftp

import (
	"github.com/jfjallid/go-smb/smb"
	"go.uber.org/zap"
	"io/fs"
)

type SMB struct {
	*zap.Logger
	*smb.Connection
}

func (SMB) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (SMB) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (SMB) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (SMB) Stat(name string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (SMB) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

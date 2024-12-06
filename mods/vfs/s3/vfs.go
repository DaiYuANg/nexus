package s3

import (
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"io/fs"
	"nexus/vfs"
)

type VFS struct {
	*minio.Client
	*zap.Logger
}

func (VFS) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (VFS) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (VFS) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (VFS) Stat(name string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (VFS) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

func (VFS) ListDir(path string) ([]vfs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (VFS) Mkdir(path string, perm fs.FileMode) error {
	//TODO implement me
	panic("implement me")
}

func (VFS) Exists(path string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (VFS) Copy(srcPath, destPath string) error {
	//TODO implement me
	panic("implement me")
}

func (VFS) Move(srcPath, destPath string) error {
	//TODO implement me
	panic("implement me")
}

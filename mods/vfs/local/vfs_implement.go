package local

import (
	"go.uber.org/zap"
	"io"
	"io/fs"
	"nexus/vfs"
)

func (v *VFS) Open(name string) (fs.File, error) {
	file, err := v.Fs.Open(name)
	if err != nil {
		v.SugaredLogger.Error("Failed to open file", zap.String("path", name), zap.Error(err))
		return nil, err
	}
	return file, nil
}

func (v *VFS) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) ListDir(path string) ([]vfs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Exists(path string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Copy(srcPath, destPath string) error {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Move(srcPath, destPath string) error {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Create(path string) (io.WriteCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Truncate(path string, size int64) error {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Chown(path string, user, group string) error {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) GetFileInfo(path string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) Symlink(target, link string) error {
	//TODO implement me
	panic("implement me")
}

func (v *VFS) ReadSymlink(path string) (string, error) {
	//TODO implement me
	panic("implement me")
}

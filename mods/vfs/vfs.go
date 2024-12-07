package vfs

import (
	"io"
	"io/fs"
)

type VFS interface {
	fs.FS
	fs.ReadDirFS
	fs.ReadFileFS
	fs.StatFS
	fs.SubFS
	io.Closer
	ListDir(path string) ([]File, error)
	Mkdir(path string, perm fs.FileMode) error
	Exists(path string) (bool, error)
	Copy(srcPath, destPath string) error
	Move(srcPath, destPath string) error
	Create(path string) (io.WriteCloser, error)
	Remove(path string) error
	Rename(oldPath, newPath string) error
	Truncate(path string, size int64) error
	WriteFile(path string, data []byte, perm fs.FileMode) error
	Chmod(path string, perm fs.FileMode) error
	Chown(path string, user, group string) error
	GetFileInfo(path string) (fs.FileInfo, error)
	RemoveAll(path string) error
	Symlink(target, link string) error
	ReadSymlink(path string) (string, error)
}

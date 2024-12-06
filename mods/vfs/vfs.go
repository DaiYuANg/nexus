package vfs

import (
	"io/fs"
)

type VFS interface {
	fs.FS
	fs.ReadDirFS
	fs.ReadFileFS
	fs.StatFS
	fs.SubFS
	ListDir(path string) ([]File, error)       // 列出目录内容
	Mkdir(path string, perm fs.FileMode) error // 创建目录
	Exists(path string) (bool, error)          // 检查文件或目录是否存在
	Copy(srcPath, destPath string) error       // 复制文件
	Move(srcPath, destPath string) error       // 移动文件
	Close() error
}

package vfs

import "io/fs"

type File interface {
	fs.FileInfo
	Md5() string
}

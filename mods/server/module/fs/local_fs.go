package fs

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/vfs/local"
)

var LocalFSModule = fx.Module("local_fs")

func NewLocalFs(logger *zap.Logger) *local.VFS {
	return local.New(local.VfsConfig{
		BasePath: "",
		Logger:   nil,
	})
}

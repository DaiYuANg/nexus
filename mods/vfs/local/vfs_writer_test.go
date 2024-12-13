package local

import (
	"github.com/adrg/xdg"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"path"
	"testing"
)

func TestVFS_WriteFile(t *testing.T) {
	t.Run("WriteFile", func(t *testing.T) {
		vfs, err := New(VfsConfig{
			BasePath:    path.Join(xdg.CacheHome, "vfs_test"),
			IgnoreWatch: []string{".git"},
			ChangeHandler: func(op fsnotify.Op, filename string) {
				println(filename)
			},
			SugaredLogger: zap.NewNop().Sugar(),
		})
		if err != nil {
			t.Errorf("WriteFile() error = %v", err)
		}
		err = vfs.WriteFile("test", []byte("hello world"))
		if err != nil {
			t.Errorf("WriteFile() error = %v", err)
		}
	})
}

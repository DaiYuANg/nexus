package memory

import "github.com/spf13/afero"

func New() *VFS {
	return &VFS{
		Fs: afero.NewMemMapFs(),
	}
}

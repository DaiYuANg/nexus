package temp

import (
	"github.com/spf13/afero"
	"os"
)

func New() *VFS {
	return &VFS{
		Fs: afero.NewBasePathFs(afero.NewOsFs(), os.TempDir()),
	}
}

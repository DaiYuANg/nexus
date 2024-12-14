package fs

import (
	"context"
	"github.com/spf13/afero"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/config"
	"nexus/vfs/local"
	"os"
)

var Module = fx.Module("fs",
	fx.Provide(
		fx.Annotate(
			newMemfs,
			fx.ResultTags(`name:"mem"`),
		),
		fx.Annotate(
			newTempFs,
			fx.ResultTags(`name:"temp"`),
		),
		newLocalFs,
	),
	fx.Invoke(startWatch),
)

func newTempFs() afero.Fs {
	return afero.NewBasePathFs(afero.NewOsFs(), os.TempDir())
}

func newMemfs() afero.Fs {
	return afero.NewMemMapFs()
}

type LocalFsParam struct {
	fx.In
	*zap.SugaredLogger
	*config.FileConfig
}

func newLocalFs(param LocalFsParam) (*local.VFS, error) {
	return local.New(local.VfsConfig{
		BasePath:      param.Data,
		SugaredLogger: param.SugaredLogger,
		IgnoreWatch:   []string{"log"},
	})
}

func startWatch(lc fx.Lifecycle, fs *local.VFS) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go fs.Watch()
			return nil
		},
	})
}

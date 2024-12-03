package fs

import (
	"github.com/spf13/afero"
	"go.uber.org/fx"
	"nexus/internal/conf"
	"os"
)

var Module = fx.Module("fs", fx.Provide(
	fx.Annotate(
		newMemfs,
		fx.ResultTags(`name:"mem"`),
	),
	fx.Annotate(
		newTempFs,
		fx.ResultTags(`name:"temp"`),
	),
	fx.Annotate(
		newLocalFs,
		fx.ResultTags(`name:"local"`),
	),
))

func newTempFs() afero.Fs {
	return afero.NewBasePathFs(afero.NewOsFs(), os.TempDir())
}

func newMemfs() afero.Fs {
	return afero.NewMemMapFs()
}

func newLocalFs(config *conf.FileConfig) afero.Fs {
	return afero.NewBasePathFs(afero.NewOsFs(), config.Path)
}

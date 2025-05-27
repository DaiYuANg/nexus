package storage

import (
	"github.com/spf13/afero"
	"go.uber.org/fx"
)

var Module = fx.Module("storage", fx.Provide(newMemfs))

func newMemfs() afero.Fs {
	return afero.NewMemMapFs()
}

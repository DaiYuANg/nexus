package print

import (
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("print", fx.Invoke(configPrint))

func configPrint(conf *koanf.Koanf) {
	conf.Print()
}

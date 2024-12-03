package print

import (
	"github.com/knadh/koanf/v2"
	"github.com/pterm/pterm"
	"go.uber.org/fx"
	"nexus/internal/conf"
)

var Module = fx.Module("print", fx.Invoke(configPrint, printHttp))

func configPrint(conf *koanf.Koanf) {
	conf.Print()
}

func printHttp(config *conf.HttpConfig) {
	pterm.DefaultBasicText.Println("Monitor:" + pterm.LightMagenta("http://127.0.0.1"+config.GetPort()))
}

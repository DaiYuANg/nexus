package print

import (
	"context"
	"github.com/knadh/koanf/v2"
	"github.com/pterm/pterm"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"nexus/internal/config"
	"os"
	"strings"
)

var Module = fx.Module("print",
	fx.Invoke(
		configPrint,
		printHttp,
		printEnv,
	),
)

func configPrint(lc fx.Lifecycle, conf *koanf.Koanf) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var text = pterm.DefaultBox.WithTitle("Config").WithTitleTopCenter(true).Print(conf.Sprint())
			pterm.DefaultCenter.WithCenterEachLineSeparately(true).Print(text)
			return nil
		},
	})
}

func printHttp(lc fx.Lifecycle, config *config.HttpConfig) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			text := pterm.FgLightCyan.Sprintf("Monitor:%s%s/%s", "http://127.0.0.1", config.GetPort(), "metrics")
			pterm.DefaultBox.WithTitle("Access URL").WithTitleTopCenter().Println(text)
			return nil
		},
	})

}

func printEnv(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var tableData = pterm.TableData{
				{"Environment", "Value"},
			}
			o := lo.Map(os.Environ(), func(item string, index int) []string {
				return strings.Split(item, "=")
			})
			lo.ForEach(o, func(item []string, index int) {
				tableData = append(tableData, item)
			})
			return pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(tableData).Render()
		},
	})
}

package cmd

import "go.uber.org/fx"

func container() *fx.App {
	return fx.New()
}

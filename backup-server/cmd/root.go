package cmd

import "go.uber.org/fx"

func Run() {
	fx.New().Run()
}

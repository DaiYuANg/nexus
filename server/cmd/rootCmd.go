package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var app *fx.App

var rootCmd = cobra.Command{
	Use: "maxio",
	PreRun: func(cmd *cobra.Command, args []string) {
		app = container()
	},
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

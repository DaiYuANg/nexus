package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var app *fx.App

var rootCmd = &cobra.Command{
	Use:   "storix",
	Short: "storix is a CLI tool for interacting with storix",
	Long:  `storix is a CLI tool for interacting with storix`,
	PreRun: func(cmd *cobra.Command, args []string) {
		app = container()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Start(cmd.Context())
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		return app.Stop(cmd.Context())
	},
}

func Execute() error {
	return rootCmd.Execute()
}

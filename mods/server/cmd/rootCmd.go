package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var container *fx.App

var rootCmd = &cobra.Command{
	Use:           "nexus",
	Long:          `Cloud Workstation`,
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		container = newDiContainer()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Start(context.Background())
	},
}

func Execute() error {
	return rootCmd.Execute()
}

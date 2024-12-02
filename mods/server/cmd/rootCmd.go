package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"os"
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

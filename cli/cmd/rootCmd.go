package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "storix",
	Short: "storix is a CLI tool for interacting with storix",
	Long:  `storix is a CLI tool for interacting with storix`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

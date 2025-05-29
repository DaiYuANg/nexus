package main

import (
	"github.com/DaiYuANg/maxio/cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}

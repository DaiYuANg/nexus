package main

import (
	"github.com/DaiYuANg/storix/server/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}

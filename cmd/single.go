package cmd

import (
	"github.com/deifyed/eeyore/cmd/single"
	"github.com/spf13/cobra"
)

// singleCmd represents the ask command
var singleCmd = &cobra.Command{
	Use:     "single",
	Aliases: []string{"s", "si"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "Ask a single question to OpenAI",
	RunE:    single.RunE(),
}

func init() {
	rootCmd.AddCommand(singleCmd)
}

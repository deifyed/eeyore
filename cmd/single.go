package cmd

import (
	"github.com/deifyed/eeyore/cmd/single"
	"github.com/spf13/cobra"
)

// singleCmd represents the ask command
var singleCmd = &cobra.Command{
	Use:  "single",
	Args: cobra.ExactArgs(1),
	RunE: single.RunE(),
}

func init() {
	rootCmd.AddCommand(singleCmd)
}

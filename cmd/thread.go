package cmd

import (
	"github.com/deifyed/eeyore/cmd/thread"
	"github.com/spf13/cobra"
)

var threadCmd = &cobra.Command{
	Use:  "thread",
	Args: cobra.NoArgs,
	RunE: thread.RunE(),
}

func init() {
	rootCmd.AddCommand(threadCmd)
}

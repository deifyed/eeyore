package cmd

import (
	"github.com/deifyed/eeyore/cmd/thread"
	"github.com/spf13/cobra"
)

var threadCmd = &cobra.Command{
	Use:     "thread",
	Aliases: []string{"t", "th"},
	Args:    cobra.NoArgs,
	Short:   "Start a conversation with OpenAI",
	RunE:    thread.RunE(),
}

func init() {
	rootCmd.AddCommand(threadCmd)
}

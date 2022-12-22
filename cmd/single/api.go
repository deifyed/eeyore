package single

import (
	"fmt"
	"io"

	"github.com/deifyed/eeyore/pkg/openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		token := viper.GetString("token")
		var question string

		if len(args) == 0 {
			rawQuestion, err := io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}

			question = string(rawQuestion)
		} else {
			question = args[0]
		}

		response, err := openai.Query(cmd.Context(), openai.QueryOptions{
			Token:     token,
			MaxTokens: 1024,
			Message:   question,
		})
		if err != nil {
			return fmt.Errorf("asking question: %w", err)
		}

		fmt.Println(response)

		return nil
	}
}

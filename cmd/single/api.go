package single

import (
	"fmt"
	"io"

	"github.com/deifyed/eeyore/pkg/config"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		token := viper.GetString(config.OpenAIToken)
		maxTokens := viper.GetInt(config.MaxTokens)
		temperature := float32(viper.GetFloat64(config.Temperature))
		model := viper.GetString(config.Model)

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

		client := gogpt.NewClient(token)

		request := gogpt.CompletionRequest{
			Model:       model,
			MaxTokens:   maxTokens,
			Temperature: temperature,
			Prompt:      question,
		}

		response, err := client.CreateCompletion(cmd.Context(), request)
		if err != nil {
			return fmt.Errorf("creating completion: %w", err)
		}

		fmt.Println(response.Choices[0].Text)

		return nil
	}
}

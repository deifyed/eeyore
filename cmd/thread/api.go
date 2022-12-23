package thread

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/deifyed/eeyore/pkg/config"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		client := gogpt.NewClient(viper.GetString(config.OpenAIToken))
		inputBuffer := bufio.NewScanner(cmd.InOrStdin())

		var maxTokens = viper.GetInt(config.MaxTokens)
		var temperature = float32(viper.GetFloat64(config.Temperature))

		for {
			fmt.Printf("> ")
			inputBuffer.Scan()

			if contains([]string{"quit", "q"}, inputBuffer.Text()) {
				break
			}

			request := gogpt.CompletionRequest{
				Model:       "text-davinci-003",
				MaxTokens:   maxTokens,
				Temperature: temperature,
				Prompt:      inputBuffer.Text(),
			}

			response, err := client.CreateCompletion(cmd.Context(), request)
			if err != nil {
				return fmt.Errorf("creating completion: %w", err)
			}

			output := fmt.Sprintf("\n%s", wordWrap(response.Choices[0].Text, 120))

			fmt.Fprintln(cmd.OutOrStdout(), output)
		}

		return nil
	}
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}

func wordWrap(text string, lineWidth int) (wrapped string) {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped = words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return
}

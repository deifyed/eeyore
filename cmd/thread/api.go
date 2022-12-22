package thread

import (
	"bufio"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		client := gogpt.NewClient(viper.GetString("token"))

		inputBuffer := bufio.NewScanner(cmd.InOrStdin())

		for {
			fmt.Printf("> ")
			inputBuffer.Scan()

			if contains([]string{"quit", "q"}, inputBuffer.Text()) {
				break
			}

			request := gogpt.CompletionRequest{
				Model:       "text-davinci-003",
				MaxTokens:   1024,
				Temperature: 0.5,
				Prompt:      inputBuffer.Text(),
			}

			response, err := client.CreateCompletion(cmd.Context(), request)
			if err != nil {
				return fmt.Errorf("creating completion: %w", err)
			}

			fmt.Println(response.Choices[0].Text)
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

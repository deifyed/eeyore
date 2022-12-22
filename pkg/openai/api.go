package openai

import (
	"context"
	"fmt"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func Query(ctx context.Context, opts QueryOptions) (string, error) {
	client := gogpt.NewClient(opts.Token)

	req := gogpt.CompletionRequest{
		Model:       model,
		MaxTokens:   opts.MaxTokens,
		Temperature: 0.5,
		Prompt:      opts.Message,
	}

	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("creating completion: %w", err)
	}

	return strings.TrimSpace(resp.Choices[0].Text), nil
}

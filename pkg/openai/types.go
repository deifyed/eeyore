package openai

type QueryOptions struct {
	Token     string
	MaxTokens int
	Message   string
}

const model = "text-davinci-003"

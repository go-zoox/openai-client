package openaiclient

import "github.com/go-zoox/core-utils/fmt"

// CreateCompletionRequest ...
type CreateCompletionRequest struct {
	// Model is ID of the model to use.
	// You can use the List models API to see all of your available models,
	// or see our Model overview for descriptions of them.
	// model is required.
	Model string `json:"model"`

	// prompt is the prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.
	// Note that <|endoftext|> is the document separator that the model sees during training,
	// so if a prompt is not specified the model will generate as if from the beginning of a new document.
	Prompt string `json:"prompt"`

	// MaxTokens is the maximum number of tokens to generate in the completion.
	// The token count of your prompt plus max_tokens cannot exceed the model's context length.
	// Most models have a context length of 2048 tokens (except for the newest models, which support 4096).
	MaxTokens int `json:"max_tokens"`

	// Temperature means What sampling temperature to use, between 0 and 2.
	// Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature"`

	// Stream whether to stream back partial progress.
	// If set, tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a data: [DONE] message.
	Stream bool `json:"stream"`

	// Stop is up to 4 sequences where the API will stop generating further tokens.
	// The returned text will not contain the stop sequence.
	Stop string `json:"stop"`

	// Suffix      string `json:"suffix"`
	// TopP        int    `json:"top_p"`
	// N           int    `json:"n"`
	// Logprobs    bool   `json:"logprobs"`
}

// CreateCompletionResponse ...
type CreateCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     int    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_token"`
		CompletionTokens int `json:"completion_token"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// reference: https://platform.openai.com/docs/api-reference/completions/create
func (c *client) CreateCompletion(cfg *CreateCompletionRequest) (*CreateCompletionResponse, error) {
	if cfg.Temperature == 0 {
		cfg.Temperature = 0.8
	}

	var apiPath string
	switch c.cfg.APIType {
	case APITypeOpenAI:
		// /v1/completions
		apiPath = fmt.Sprintf("/%s/%s", c.cfg.APIVersion, ResourceCompletion)
	case APITypeAzure:
		// openai/deployments/{deployment_id}/completions
		apiPath = fmt.Sprintf("/openai/deployments/%s/%s", c.cfg.AzureDeployment, ResourceCompletion)
	}

	resp, err := c.post(apiPath, cfg)
	if err != nil {
		return nil, err
	}

	var response CreateCompletionResponse
	if err := resp.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

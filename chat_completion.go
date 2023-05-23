package openaiclient

import "fmt"

// CreateChatCompletionRequest ...
type CreateChatCompletionRequest struct {
	// Model is ID of the model to use.
	// You can use the List models API to see all of your available models,
	// or see our Model overview for descriptions of them.
	// model is required.
	Model string `json:"model"`

	// messages is the messages to generate chat completions for, in the chat format.
	Messages []CreateChatCompletionMessage `json:"messages"`

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

	// User is a unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.
	User string `json:"user"`
}

// CreateChatCompletionResponse ...
type CreateChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Model   string `json:"model"`
	Choices []struct {
		Message      CreateChatCompletionMessage `json:"message"`
		Index        int                         `json:"index"`
		Logprobs     int                         `json:"logprobs"`
		FinishReason string                      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// CreateChatCompletionMessage ...
type CreateChatCompletionMessage struct {
	// Role is the mssage role, available: system | user | assistant
	Role    string `json:"role"`
	Content string `json:"content"`
}

// reference: https://platform.openai.com/docs/api-reference/completions/create
func (c *client) CreateChatCompletion(cfg *CreateChatCompletionRequest) (*CreateChatCompletionResponse, error) {
	if cfg.Temperature == 0 {
		cfg.Temperature = 0.8
	}

	var apiPath string
	switch c.cfg.APIType {
	case APITypeOpenAI:
		// /v1/completions
		apiPath = fmt.Sprintf("/%s/%s", c.cfg.APIVersion, ResourceChatCompletion)
	case APITypeAzure:
		// openai/deployments/{deployment_id}/completions
		apiPath = fmt.Sprintf("/openai/deployments/%s/%s", c.cfg.AzureDeployment, ResourceChatCompletion)
	}

	resp, err := c.post(apiPath, cfg)
	if err != nil {
		return nil, err
	}

	var response CreateChatCompletionResponse
	if err := resp.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

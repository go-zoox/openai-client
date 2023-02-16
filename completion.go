package openaiclient

// CreateCompletionRequest ...
type CreateCompletionRequest struct {
	// Model is ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
	Stream      bool   `json:"stream"`
	Stop        string `json:"stop"`
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
	resp, err := c.post("/v1/completions", cfg)
	if err != nil {
		return nil, err
	}

	var response CreateCompletionResponse
	if err := resp.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

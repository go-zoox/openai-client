package openaiclient

import "fmt"

// ListModelRequest ...
type ListModelRequest struct{}

// ListModelResponse ...
type ListModelResponse struct {
	Data   []RetrieveModelResponse `json:"data"`
	Object string                  `json:"object"`
}

// RetrieveModelRequest ...
type RetrieveModelRequest struct {
}

// RetrieveModelResponse ...
type RetrieveModelResponse struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	OwnedBy    string `json:"owned_by"`
	Permission []struct {
		ID                 string `json:"id"`
		Object             string `json:"object"`
		Created            int64  `json:"created"`
		AllowCreateEngine  bool   `json:"allow_create_engine"`
		AllowSampling      bool   `json:"allow_sampling"`
		AllowLogprobs      bool   `json:"allow_logprobs"`
		ALlowSearchIndices bool   `json:"allow_search_indices"`
		AllowView          bool   `json:"allow_view"`
		AllowFineTuning    bool   `json:"allow_fine_tuning"`
		Organization       string `json:"organization"`
		// Group string `json:"group"`
		IsBlocking bool `json:"is_blocking"`
	} `json:"permission"`
	Root string `json:"root"`
	// Parent string `json:"parent"`
}

func (c *client) ListModels() (*ListModelResponse, error) {
	resp, err := c.get("/models", nil)
	if err != nil {
		return nil, err
	}

	var response ListModelResponse
	if err := resp.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) RetrieveModel(model string) (*RetrieveModelResponse, error) {
	resp, err := c.get(fmt.Sprintf("/models/%s", model), nil)
	if err != nil {
		return nil, err
	}

	var response RetrieveModelResponse
	if err := resp.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

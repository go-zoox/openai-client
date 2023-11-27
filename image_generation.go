package openaiclient

import "fmt"

// Image ...
type Image struct {
	// URL is the URL of the generated image, if `response_format` is `url`.
	URL string `json:"url"`

	// B64JSON is the JSON representation of the image, if `response_format` is `b64_json`.
	B64JSON string `json:"b64_json"`

	// RevisedPrompt is the prompt that was used to generate the image,
	// if there was any revision to the prompt.
	RevisedPrompt string `json:"revised_prompt"`
}

// ImageGenerationRequest ...
type ImageGenerationRequest struct {
	// Prompt is a text description of the desired image(s).
	// The maximum length is 1000 characters for `dall-e-2` and
	// 4000 characters for `dall-e-3`.
	// Required.
	Prompt string `json:"prompt"`

	// Model is the model to use for image generation
	// Optional, default is dall-e-2.
	Model string `json:"model"`

	// N is the number of images to generate.
	// Must be between 1 and 1-10.
	// For `dall-e-3`, only `n=1` is supported.
	// Optional, default is 1.
	N int `json:"n"`

	// Quality is the qualityof the image that will be generated.
	// `hd` creates images with finer details and greater consistency
	// across the image. This param is only supported for `dall-e-3`.
	// Optional, default is `standard`.
	Quality string `json:"quality"`

	// ResponseFormat is the format in which the generated images are returned.
	// Must be one of `url` or `b64_json`.
	// Optional, default is `url`.
	ResponseFormat string `json:"response_format"`

	// Size is the size of the generated images.
	// Must be one of `256x256`, `512x512`, or `1024x1024` for `dall-e-2`.
	// Must be one of `1024x1024`, `1792x1024`, or `1024x1792` for `dall-e-3`.
	// Optional, default is `1024x1024`.
	Size string `json:"size"`

	// Style is the style of the generated images.
	// Must be one of `vivid` or `normal`.
	// Vivid causes the model to lean towards generating hyper-real and dramatic images.
	// Natural causes the model to produce more natural, less hyper-real looking images.
	// This param is only supported for `dall-e-3`.
	// Optional, default is `vivid`.
	Style string `json:"style"`

	// User is a unique identifier representing your end-user,
	// which can help OpenAI to monitor and detect abuse.
	// Optional.
	User string `json:"user"`
}

// ImageGenerationResponse ...
type ImageGenerationResponse struct {
	Data    []Image `json:"data"`
	Created int64   `json:"created"`
}

func (c *client) ImageGeneration(cfg *ImageGenerationRequest) (*ImageGenerationResponse, error) {
	if cfg.N == 0 {
		cfg.N = 1
	}
	if cfg.Size == "" {
		cfg.Size = "1792x1024"
	}
	if cfg.Quality == "" {
		cfg.Quality = "standard"
	}
	if cfg.ResponseFormat == "" {
		cfg.ResponseFormat = "url"
	}
	if cfg.Model == "" {
		cfg.Model = "dall-e-3"
	}
	if cfg.Style == "" {
		cfg.Style = "vivid"
	}

	var apiPath string
	switch c.cfg.APIType {
	case APITypeOpenAI:
		// /completions
		apiPath = fmt.Sprintf("/%s", ResourceImageGeneration)
	case APITypeAzure:
		// openai/deployments/{deployment_id}/completions
		apiPath = fmt.Sprintf("/openai/deployments/%s/%s", c.cfg.AzureDeployment, ResourceImageGeneration)
	}

	resp, err := c.post(apiPath, cfg)
	if err != nil {
		return nil, err
	}

	var response ImageGenerationResponse
	if err := resp.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

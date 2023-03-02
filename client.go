package openaiclient

import (
	"fmt"

	"github.com/go-zoox/fetch"
)

// DefaultAPIServer is OpenAI API Server.
const DefaultAPIServer = "https://api.openai.com"

// Client is a OpenAI API Client.
type Client interface {
	// CreateCompletion creates a completion.
	CreateCompletion(cfg *CreateCompletionRequest) (*CreateCompletionResponse, error)

	// CreateChatCompletion creates a chat completion.
	CreateChatCompletion(cfg *CreateChatCompletionRequest) (*CreateChatCompletionResponse, error)

	// ListModels lists all models.
	ListModels() (*ListModelResponse, error)

	// RetrieveModels retrieves a model.
	RetrieveModel(model string) (*RetrieveModelResponse, error)
}

type client struct {
	cfg *Config
}

// Config is the OpenAI Client configuration.
type Config struct {
	APIKey    string `json:"api_key"`
	APIServer string `json:"api_server"`
}

// New creates a OpenAI Client.
func New(cfg *Config) (Client, error) {
	if cfg.APIServer == "" {
		cfg.APIServer = DefaultAPIServer
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("api key is required")
	}

	return &client{
		cfg: cfg,
	}, nil
}

func (c *client) post(path string, body fetch.Body) (*fetch.Response, error) {
	response, err := fetch.Post(path, &fetch.Config{
		BaseURL: c.cfg.APIServer,
		Headers: fetch.Headers{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", c.cfg.APIKey),
		},
		Body: body,
	})
	if err != nil {
		return nil, err
	}

	if !response.Ok() {
		return nil, fmt.Errorf("failed to create completion: %s", response.Get("error").String())
	}

	return response, nil
}

func (c *client) get(path string, query fetch.Query) (*fetch.Response, error) {
	response, err := fetch.Get(path, &fetch.Config{
		BaseURL: c.cfg.APIServer,
		Headers: fetch.Headers{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", c.cfg.APIKey),
		},
		Query: query,
	})
	if err != nil {
		return nil, err
	}

	if !response.Ok() {
		return nil, fmt.Errorf("failed to create completion: %s", response.Get("error").String())
	}

	return response, nil
}

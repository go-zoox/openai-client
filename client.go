package openaiclient

import (
	"fmt"
	"time"

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
	// APIKey is the OpenAI API Key.
	APIKey string `json:"api_key"`

	// APIServer customs the OpenAI API Server.
	APIServer string `json:"api_server"`

	// APIType specify the OpenAI API Type, available: azure, default: empty (openai).
	APIType string `json:"api_type"`

	// APIVersion specify the OpenAI API Version, available: v1, default: empty (v1).
	// if APIType is azure, APIVersion should not be empty.
	APIVersion string `json:"api_version"`

	// AzureDeployment is the Azure Deployment.
	AzureDeployment string `json:"azure_deployment"`

	// Proxy sets the request proxy.
	//
	//	support http, https, socks5
	//	example:
	//		http://127.0.0.1:17890
	//		https://127.0.0.1:17890
	//		socks5://127.0.0.1:17890
	Proxy string `json:"proxy"`

	// Timeout sets the request timeout.
	// default: 300s
	Timeout time.Duration `json:"timeout"`
}

// New creates a OpenAI Client.
func New(cfg *Config) (Client, error) {
	if cfg.APIServer == "" {
		cfg.APIServer = DefaultAPIServer
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("api key is required")
	}

	if cfg.APIType == "" {
		cfg.APIType = APITypeOpenAI
	}

	if cfg.APIVersion == "" {
		switch cfg.APIType {
		case APITypeOpenAI:
			cfg.APIVersion = "v1"
		case APITypeAzure:
			return nil, fmt.Errorf("azure api version is required")
		default:
			return nil, fmt.Errorf("unknown api type: %s", cfg.APIType)
		}
	}

	if cfg.APIType == APITypeAzure && cfg.AzureDeployment == "" {
		return nil, fmt.Errorf("azure deployment is required")
	}

	if cfg.Timeout == 0 {
		cfg.Timeout = 300 * time.Second
	}

	return &client{
		cfg: cfg,
	}, nil
}

func (c *client) post(path string, body fetch.Body) (*fetch.Response, error) {
	headers := c.buildHeaders()
	query := c.buildQuery()

	response, err := fetch.Post(path, &fetch.Config{
		BaseURL: c.cfg.APIServer,
		Headers: headers,
		Query:   query,
		Body:    body,
		//
		Proxy: c.cfg.Proxy,
		//
		Timeout: c.cfg.Timeout,
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
	headers := c.buildHeaders()

	response, err := fetch.Get(path, &fetch.Config{
		BaseURL: c.cfg.APIServer,
		Headers: headers,
		Query:   query,
		//
		Proxy: c.cfg.Proxy,
		//
		Timeout: c.cfg.Timeout,
	})
	if err != nil {
		return nil, err
	}

	if !response.Ok() {
		return nil, fmt.Errorf("failed to create completion: %s", response.Get("error").String())
	}

	return response, nil
}

func (c *client) buildHeaders() fetch.Headers {
	headers := fetch.Headers{
		"Content-Type": "application/json",
	}

	switch c.cfg.APIType {
	case APITypeOpenAI:
		headers["Authorization"] = fmt.Sprintf("Bearer %s", c.cfg.APIKey)
	case APITypeAzure:
		headers["api-key"] = c.cfg.APIKey
	}

	return headers
}

func (c *client) buildQuery() fetch.Query {
	query := fetch.Query{}

	switch c.cfg.APIType {
	case APITypeAzure:
		query["api-version"] = c.cfg.APIVersion
	}

	return query
}

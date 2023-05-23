package openaiclient

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"

	_ "github.com/go-zoox/dotenv"
)

func TestCreateChatCompletion(t *testing.T) {
	client, _ := New(&Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		// Proxy:  "socks5://127.0.0.1:17890",
	})

	completion, err := client.CreateChatCompletion(&CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []CreateChatCompletionMessage{
			{Role: "user", Content: "你好用英语怎么说？"},
		},
		MaxTokens: 3000, // 4097
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(completion)
}

func TestCreateChatCompletionAzure(t *testing.T) {
	client, _ := New(&Config{
		APIKey:          os.Getenv("OPENAI_API_KEY"),
		APIType:         APITypeAzure,
		APIServer:       os.Getenv("AZURE_API_SERVER"),
		APIVersion:      os.Getenv("AZURE_API_VERSION"),
		AzureDeployment: os.Getenv("AZURE_DEPLOYMENT"),
	})

	completion, err := client.CreateChatCompletion(&CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []CreateChatCompletionMessage{
			{Role: "user", Content: "你好用英语怎么说？"},
		},
		MaxTokens: 3000, // 4097
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(completion)
}

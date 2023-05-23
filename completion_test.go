package openaiclient

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"

	_ "github.com/go-zoox/dotenv"
)

func TestCreateCompletion(t *testing.T) {
	client, _ := New(&Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
	})

	completion, err := client.CreateCompletion(&CreateCompletionRequest{
		Model:       "text-davinci-003",
		Prompt:      "你好用英语怎么说？",
		MaxTokens:   4000, // 4097
		Temperature: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(completion)
}

func TestCreateCompletionAzure(t *testing.T) {
	client, err := New(&Config{
		APIKey:          os.Getenv("OPENAI_API_KEY"),
		APIType:         APITypeAzure,
		APIServer:       os.Getenv("AZURE_API_SERVER"),
		APIVersion:      os.Getenv("AZURE_API_VERSION"),
		AzureDeployment: os.Getenv("AZURE_DEPLOYMENT"),
	})
	if err != nil {
		t.Fatal(err)
	}

	completion, err := client.CreateCompletion(&CreateCompletionRequest{
		Model:       "text-davinci-003",
		Prompt:      "你好用英语怎么说？",
		MaxTokens:   4000, // 4097
		Temperature: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(completion)
}

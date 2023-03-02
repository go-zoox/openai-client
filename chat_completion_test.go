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

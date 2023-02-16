package openaiclient

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"

	_ "github.com/go-zoox/dotenv"
)

func TestListModel(t *testing.T) {
	client, _ := New(&Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
	})

	models, err := client.ListModels()
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(models)
}

func TestRetrieveModel(t *testing.T) {
	client, _ := New(&Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
	})

	model, err := client.RetrieveModel("text-davinci-003")
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(model)
}

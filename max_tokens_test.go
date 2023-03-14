package openaiclient

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestGetMaxTokens(t *testing.T) {
	testify.Equal(t, GetMaxTokens(ModelGPT_4), 8192)
	testify.Equal(t, GetMaxTokens(ModelGPT_4_0314), 8192)
	testify.Equal(t, GetMaxTokens(ModelGPT_4_32K), 32768)
	testify.Equal(t, GetMaxTokens(ModelGPT_4_32K_0314), 32768)

	testify.Equal(t, GetMaxTokens(ModelGPT3_5Turbo), 4096)
	testify.Equal(t, GetMaxTokens(ModelTextDavinci003), 4097)

	testify.Equal(t, GetMaxTokens("unknown_model"), MaxTokensDefault)
}

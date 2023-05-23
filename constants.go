package openaiclient

const (
	// APITypeOpenAI means OpenAI API
	APITypeOpenAI = "openai"

	// APITypeAzure means Azure OpenAI API
	APITypeAzure = "azure"
)

const (
	// ResourceCompletion ...
	ResourceCompletion = "completions"
	// ResourceChatCompletion ...
	ResourceChatCompletion = "chat/completions"
	// ResourceEmbedding ...
	ResourceEmbedding = "embedding"
)

const (
	// ModelAda ...
	ModelAda = "ada"
	// ModelBabbage ...
	ModelBabbage = "babbage"
	// ModelCurie ...
	ModelCurie = "curie"
	// ModelDavinci ...
	ModelDavinci = "davinci"
)

// Text Models => https://platform.openai.com/docs/models/gpt-3
const (
	// ModelTextDavinci003 ...
	ModelTextDavinci003 = "text-davinci-003"
	// ModelTextDavinci002 ...
	ModelTextDavinci002 = "text-davinci-002"
	// ModelTextDavinci001 ...
	ModelTextDavinci001 = "text-davinci-001"
	// ModelTextCurie001 ...
	ModelTextCurie001 = "text-curie-001"
	// ModelTextBabbage001 ...
	ModelTextBabbage001 = "text-babbage-001"
	// ModelTextAda001 ...
	ModelTextAda001 = "text-ada-001"
)

// Chat Models => https://platform.openai.com/docs/api-reference/chat
const (
	ModelGPT3_5Turbo     = "gpt-3.5-turbo"
	ModelGPT3_5Turbo0301 = "gpt-3.5-turbo-0301"
	// GPT-4 https://platform.openai.com/docs/models/gpt-4
	ModelGPT_4          = "gpt-4"
	ModelGPT_4_0314     = "gpt-4-0314"
	ModelGPT_4_32K      = "gpt-4-32k"
	ModelGPT_4_32K_0314 = "gpt-4-32k-0314"
)

// Codex => https://platform.openai.com/docs/models/codex
const (
	// ModelTextDavinci002 ...
	ModelCodeDavinci002 = "code-davinci-002"
	// ModelCodeCushman001 ...
	ModelCodeCushman001 = "code-cushman-001"
)

// MaxTokensDefault ...
const MaxTokensDefault = 4096

// MaxTokensMap ...
var MaxTokensMap = map[string]int64{
	//
	ModelGPT_4:          8192,
	ModelGPT_4_0314:     8192,
	ModelGPT_4_32K:      32768,
	ModelGPT_4_32K_0314: 32768,
	//
	ModelGPT3_5Turbo:     4096,
	ModelGPT3_5Turbo0301: 4096,
	//
	ModelTextDavinci003: 4097,
	ModelTextDavinci002: 4097,
	//
	ModelCodeDavinci002: 8001,
}

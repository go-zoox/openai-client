package openaiclient

// GetMaxTokens gets the max tokens for specific model.
func GetMaxTokens(modelName string) int64 {
	if v, ok := MaxTokensMap[modelName]; ok {
		return v
	}

	return MaxTokensDefault
}

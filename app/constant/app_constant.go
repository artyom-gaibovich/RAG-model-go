package constant

import (
	"os"
)

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}

type ApiKeyGPT string
type ActionPrompt int

const (
	PromptConnectText ActionPrompt = iota + 1
	PromptDelAdvText
	PromptChangeText
)

func GetApiKeyGPT() string {
	apiKey := ApiKeyGPT("")
	return apiKey.getApiKeyGPT()
}

func (a ApiKeyGPT) getApiKeyGPT() string {
	return os.Getenv("API_KEY_GPT")
}

func GetTextPrompt(prompt ActionPrompt) string {
	return prompt.getTextPrompt()
}

func (a ActionPrompt) getTextPrompt() string {
	switch a {
	case PromptConnectText:
		return os.Getenv("PROMPT_CONNECT_TEXT")
	case PromptDelAdvText:
		return os.Getenv("PROMPT_DEL_ADV_TEXT")
	case PromptChangeText:
		return os.Getenv("PROMPT_CHANGE_TEXT")
	default:
		return ""
	}
}

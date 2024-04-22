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

type Proxy string
type ProxyUsername string
type ProxyPassword string

type LimitCharacters uint16

const (
	PromptConnectText ActionPrompt = iota + 1
	PromptDelAdvText
	PromptChangeText
)

func GetApiKeyGPT() string {
	apiKey := ApiKeyGPT("")
	return apiKey.getApiKeyGPT()
}

func GetProxy() string {
	proxy := Proxy("")
	return proxy.getProxy()
}

func GetProxyUsername() string {
	proxyUsername := ProxyUsername("")
	return proxyUsername.getProxyUsername()
}

func GetProxyPassword() string {
	proxyPassword := ProxyPassword("")
	return proxyPassword.getProxyPassword()
}

func (a ApiKeyGPT) getApiKeyGPT() string {
	return os.Getenv("API_KEY_GPT")
}

func (p Proxy) getProxy() string {
	return os.Getenv("PROXY")
}

func (p ProxyUsername) getProxyUsername() string {
	return os.Getenv("PROXY_USERNAME")
}

func (p ProxyPassword) getProxyPassword() string {
	return os.Getenv("PROXY_PASSWORD")
}

func GetTextPrompt(prompt ActionPrompt) string {
	return prompt.getTextPrompt()
}

func GetLimitCharacters() string {
	limitCharacters := LimitCharacters(4096)
	return limitCharacters.getLimitCharacters()
}

func (l LimitCharacters) getLimitCharacters() string {
	return os.Getenv("LIMIT_CHARACTERS")
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

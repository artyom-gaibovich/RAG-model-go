package constant

import "os"

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

type ApiKeyGPT string

func (a ApiKeyGPT) getApiKeyGPT() string {
	return os.Getenv("API_KEY_GPT")
}

func GetApiKeyGPT() string {
	apiKey := ApiKeyGPT("")
	return apiKey.getApiKeyGPT()
}

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}

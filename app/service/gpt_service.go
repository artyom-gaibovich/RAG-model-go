// Package service предоставляет сервис для генерации текстов с использованием ChatGPT.
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/APSN4/RAG-model-go/app/constant"
	"github.com/APSN4/RAG-model-go/app/domain/dto"
	"github.com/APSN4/RAG-model-go/app/pkg"
	"github.com/APSN4/RAG-model-go/config/proxy"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type GPTService interface {
	GenerateText(c *gin.Context)
}

type GPTServiceImpl struct{}

// GenerateText генерирует текст с использованием ChatGPT.
func (g GPTServiceImpl) GenerateText(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request dto.GPTModel
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	configOpenai := openai.DefaultConfig(constant.GetApiKeyGPT())
	configOpenai.HTTPClient = proxy.CreateProxy()
	client := openai.NewClientWithConfig(configOpenai)
	textToAI, errL := promptAgentTexts(request.RequestTexts, request.ModeGen)
	if errL != nil {
		pkg.PanicException(constant.UnknownError)
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: textToAI,
				},
			},
		},
	)

	if err != nil {
		log.Error(err)
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, resp.Choices[0].Message.Content))
}

func promptAgentTexts(textSlice []string, modeGen string) (string, error) {
	readyArticle := createOneFullText(textSlice)
	if !CheckLimitTexts(textSlice) {
		log.Error("Error len text > ", constant.GetLimitCharacters())
		return "err", errors.New("error len text")
	}
	modeArticle := []string{modeGen, readyArticle}
	modeArticleReady := createOneFullText(modeArticle)
	fmt.Println(modeArticleReady)
	return modeArticleReady, nil
}

func createOneFullText(textSlice []string) string {
	readyText := strings.Join(textSlice, "\n")
	return readyText
}

func GPTServiceInit() *GPTServiceImpl {
	return &GPTServiceImpl{}
}

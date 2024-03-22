package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"net/http"
	"rag-model/app/constant"
	"rag-model/app/domain/dto"
	"rag-model/app/pkg"
	"strings"
)

type GPTService interface {
	GenerateText(c *gin.Context)
}

type GPTServiceImpl struct{}

func (g GPTServiceImpl) GenerateText(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request dto.GPTModel
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	client := openai.NewClient(constant.GetApiKeyGPT())
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: createOneText(request.RequestTexts),
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

func createOneText(textSlice []string) string {
	readyText := strings.Join(textSlice, "\n")
	return readyText
}

func GPTServiceInit() *GPTServiceImpl {
	return &GPTServiceImpl{}
}

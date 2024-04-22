package service

import (
	"github.com/APSN4/RAG-model-go/app/constant"
	"github.com/gin-gonic/gin"
	"strconv"
)

type LimitService interface {
	CheckLimitTextsNetwork(c *gin.Context)
}

type LimitServiceImpl struct{}

func (l LimitServiceImpl) CheckLimitTextsNetwork(c *gin.Context) {}

func CheckLimitTexts(texts []string) bool {
	localText := createOneFullText(texts)
	countCharacters := len(localText)
	limitCharacters, _ := strconv.Atoi(constant.GetLimitCharacters())
	if countCharacters > limitCharacters {
		return false
	}
	return true
}

func LimitServiceInit() *LimitServiceImpl {
	return &LimitServiceImpl{}
}

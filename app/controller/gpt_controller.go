package controller

import (
	"github.com/gin-gonic/gin"
	"rag-model/app/service"
)

type GPTController interface {
	GenerateText(c *gin.Context)
}

type GPTControllerImpl struct {
	svc service.GPTService
}

func (g GPTControllerImpl) GenerateText(c *gin.Context) {
	g.svc.GenerateText(c)
}

func GPTControllerInit(GPTService service.GPTService) *GPTControllerImpl {
	return &GPTControllerImpl{
		svc: GPTService,
	}
}

package controller

import (
	"github.com/APSN4/RAG-model-go/app/service"
	"github.com/gin-gonic/gin"
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

package controller

import (
	"github.com/APSN4/RAG-model-go/app/service"
	"github.com/gin-gonic/gin"
)

type LimitController interface {
	CheckLimitTextsNetwork(c *gin.Context)
}

type LimitControllerImpl struct {
	svc service.LimitService
}

func (l LimitControllerImpl) CheckLimitTextsNetwork(c *gin.Context) {
	l.svc.CheckLimitTextsNetwork(c)
}

func LimitControllerInit(LimitService service.LimitService) *LimitControllerImpl {
	return &LimitControllerImpl{
		svc: LimitService,
	}
}

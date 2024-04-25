package controller

import (
	"github.com/APSN4/RAG-model-go/app/service"
	"github.com/gin-gonic/gin"
)

type ChangeAPIController interface {
	ChangeAPI(c *gin.Context)
}

type ChangeAPIControllerImpl struct {
	svc service.ChangeAPIService
}

func (ch ChangeAPIControllerImpl) ChangeAPI(c *gin.Context) {
	ch.svc.ChangeAPI(c)
}

func ChangeAPIControllerInit(ChangeAPIService service.ChangeAPIService) *ChangeAPIControllerImpl {
	return &ChangeAPIControllerImpl{
		svc: ChangeAPIService,
	}
}

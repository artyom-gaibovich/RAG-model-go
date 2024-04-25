package service

import (
	"github.com/gin-gonic/gin"
)

type ChangeAPIService interface {
	ChangeAPI(c *gin.Context)
}

type ChangeAPIServiceImpl struct{}

func (ch ChangeAPIServiceImpl) ChangeAPI(c *gin.Context) {

}

func ChangeAPIServiceInit() *ChangeAPIServiceImpl {
	return &ChangeAPIServiceImpl{}
}

package service

import (
	"github.com/gin-gonic/gin"
)

type GPTService interface {
	GenerateText(c *gin.Context)
	ChangeGeneratedText(c *gin.Context)
}

//type UserServiceImpl struct {
//	userRepository repository.UserRepository
//}

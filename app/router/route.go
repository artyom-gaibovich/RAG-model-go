package router

import (
	"github.com/APSN4/RAG-model-go/config"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		gpt := api.Group("/gpt")
		gpt.POST("/generate", init.GPTCtrl.GenerateText)
		service := api.Group("/service")
		service.POST("/api", init.ChangeAPICtrl.ChangeAPI)

		userComponents := api.Group("/component")
		userComponents.GET("/:userID", init.UserCtrl.UpdateComponentUserData)
	}

	return router
}

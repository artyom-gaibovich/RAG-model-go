package router

import (
	"github.com/gin-gonic/gin"
	"rag-model/config"
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

		userPhoto := api.Group("/photo")
		userPhoto.POST("/:userID", init.UserCtrl.UploadPhotoUser)

		userComponents := api.Group("/component")
		userComponents.GET("/:userID", init.UserCtrl.UpdateComponentUserData)
	}

	return router
}

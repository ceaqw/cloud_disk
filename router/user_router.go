package router

import (
	"CouldDisk/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userController := new(controller.User)
	dashBoardRouter := router.Group("/file")

	{
		dashBoardRouter.POST("/login", userController.Login)
		dashBoardRouter.GET("/checkuserlogininfo", userController.Login)
	}
}

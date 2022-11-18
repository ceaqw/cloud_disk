package router

import (
	"CouldDisk/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userController := new(controller.User)
	dashBoardRouter := router.Group("/file")

	{
		dashBoardRouter.POST("/getfile", userController.GetUser)
	}
}

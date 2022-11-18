package router

import (
	"CouldDisk/controller"

	"github.com/gin-gonic/gin"
)

func InitFileRouter(router *gin.RouterGroup) {
	fileController := new(controller.File)
	dashBoardRouter := router.Group("/file")

	{
		dashBoardRouter.POST("/getfile", fileController.GetFile)
	}
}

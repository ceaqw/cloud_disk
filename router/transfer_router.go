package router

import (
	"CouldDisk/controller"

	"github.com/gin-gonic/gin"
)

func InitTransferRouter(router *gin.RouterGroup) {
	transferController := new(controller.Transfer)
	dashBoardRouter := router.Group("/filetransfer")

	{
		dashBoardRouter.GET("/getstorage", transferController.GetStorage)
	}
}

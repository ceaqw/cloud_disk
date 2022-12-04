package router

import (
	"CouldDisk/controller"

	"github.com/gin-gonic/gin"
)

func InitFileRouter(router *gin.RouterGroup) {
	fileController := new(controller.File)
	dashBoardRouter := router.Group("/file")
	recoveryFile := router.Group("/recoveryfile")
	shareFile := router.Group("/share")
	{
		dashBoardRouter.GET("/getfilelist", fileController.GetFileList)
		dashBoardRouter.GET("/selectfilebyfiletype", fileController.GetFileListByType)
		dashBoardRouter.GET("/getfiletree", fileController.GetFileTree)
	}
	{
		recoveryFile.POST("list", fileController.GetRecoveryFileList)
	}
	{
		shareFile.GET("shareList", fileController.GetShareFileList)
	}
}

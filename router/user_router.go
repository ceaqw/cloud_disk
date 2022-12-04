package router

import (
	"CouldDisk/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userController := new(controller.User)
	dashBoardRouter := router.Group("/user")

	{
		dashBoardRouter.POST("/login", userController.Login)
		dashBoardRouter.POST("/register", userController.Register)
		dashBoardRouter.GET("/checkuserlogininfo", userController.CheckUserLoginInfo)
		dashBoardRouter.GET("/sendcode", userController.SendCode)
		dashBoardRouter.GET("/updatepwdbyemail", userController.UpdatePwdByEmail)
	}
}

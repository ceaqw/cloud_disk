package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	router := gin.New()
	// 身份验证
	// router.Use(jwt.JWTAuth())
	// 日志记录
	// router.Use(logger.LoggerToFile())
	baseRouter := router.Group("/controller")
	{
		InitFileRouter(baseRouter)
		InitUserRouter(baseRouter)
	}
	return router
}
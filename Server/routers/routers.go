package routers

import (
	"Server/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 50 << 20 //限制文件上传大小
	router.Static("/storage", "./storage")
	rd := router.Group("/rd")
	user := rd.Group("/user")
	{
		user.GET("/history", controller.HistHandler)  // 历史查询
		user.POST("/register", controller.RegHandler) // 注册
		user.POST("/login", controller.LoginHandler)  // 登录
	}

	return router
}

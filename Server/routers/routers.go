package routers

import (
	"Server/controller"
	"Server/middleware"
	"Server/websocket"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 50 << 20 //限制文件上传大小
	router.Static("/storage", "./storage")
	rd := router.Group("/rd")
	user := rd.Group("/user")
	{
		user.GET("/history", middleware.AuthByToken(), controller.HistHandler) // 历史查询
		user.POST("/register", controller.RegHandler)                          // 注册
		user.POST("/login", controller.LoginHandler)                           // 登录
	}

	detect := router.Group("/detect")
	{
		detect.POST("/upload_text", middleware.AuthByToken(), controller.TextHandler)
		detect.POST("/upload_file", middleware.AuthByToken(), func(c *gin.Context) {
			websocket.ServeWebSocket(c.Writer, c.Request)
		})
	}

	return router
}

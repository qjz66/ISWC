package routers

import (
	"Server/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 50 << 20 //限制文件上传大小
	router.Static("/videos", "./videos")
	dy := router.Group("/douyin")
	dy.GET("/feed", controller.FeedHandler) // 视频流接口
	user := dy.Group("/user")
	{
		user.GET("/")          // 用户信息
		user.POST("/register") // 注册
		user.POST("/login")    // 登录
	}

	return router
}

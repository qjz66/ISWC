package routers

import (
	"Server/controller"
	"Server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.MaxMultipartMemory = 50 << 20 //限制文件上传大小
	router.Static("/storage", "./storage")
	rd := router.Group("/rd")
	user := rd.Group("/user")
	{
		user.GET("/history", middleware.AuthByToken(), controller.HistHandler)            // 历史查询
		user.POST("/register", controller.RegHandler)                                     // 注册
		user.POST("/login", controller.LoginHandler)                                      // 登录
		user.PUT("/update_password", middleware.AuthByToken(), controller.UpdatePassword) //修改密码
		user.PUT("/update_username", middleware.AuthByToken(), controller.UpdateUsername) //修改用户名
	}
	ground := rd.Group("/ground")
	{
		ground.PUT("/favorite", middleware.AuthByToken(), controller.FavoriteHandler)
		ground.POST("/comment", middleware.AuthByToken(), controller.CommentHandler)
		ground.POST("/update", middleware.AuthByToken(), controller.UpdateHandler)
		ground.GET("/blacklist", middleware.AuthByToken(), controller.BlackListHandler)
		ground.GET("/get_comments", middleware.AuthByToken(), controller.GetCommentsHandler)
		ground.GET("/get_updates", middleware.AuthByToken(), controller.GetUpdateListHandler)
	}
	detect := rd.Group("/detect")
	{
		detect.POST("/upload_text", middleware.AuthByToken(), controller.TextHandler)
		detect.POST("/upload_file", middleware.AuthByToken(), controller.FileHandler)
	}
	push := rd.Group("/push")
	{
		push.GET("get_rumors", middleware.AuthByToken(), controller.PushHandler)
	}

	return router
}

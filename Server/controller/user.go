package controller

import (
	"Server/service"
	"Server/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserRegReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserRegResp struct {
	ID uint `json:"id"`
}
type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginResp struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}
type UserHisReq struct{}
type UserHisResp struct{}

// 用户查询历史记录
func HistHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// 用户注册
func RegHandler(c *gin.Context) {
	var id uint
	var err error
	req := UserRegReq{}
	resp := UserRegResp{}
	svc := service.NewService(c)

	//绑定参数至结构体
	req.Username = c.Query("username")
	req.Password = c.Query("password")

	//将用户信息插入数据库
	id, err = svc.InsertUser(req.Username, req.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
	}

	resp.ID = id
	c.JSON(http.StatusOK, resp)
}

// 用户登录
func LoginHandler(c *gin.Context) {
	req := UserLoginReq{}
	resp := UserLoginResp{}
	svc := service.NewService(c)
	var id uint
	var token string

	req.Username = c.Query("username")
	req.Password = c.Query("password")

	id = svc.GetIDByName(req.Username, req.Password)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username or Password error",
		})
		return
	}
	token, err := util.GenerateToken(int64(id), req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	resp.Token = token
	resp.ID = id
	c.JSON(http.StatusOK, resp)
}

package controller

import (
	"Server/model"
	"Server/service"
	"Server/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserRegReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserRegResp struct {
	ID int64 `json:"id"`
}
type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginResp struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}
type UserHisReq struct{}
type UserHisResp struct{}

type UpdatePasswordReq struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
}
type UpdatePasswordResp struct {
	Message string `json:"message"`
}
type UpdateUsernameReq struct {
	NewUsername string `json:"new_username"`
	ID          int64  `json:"id"`
}
type UpdateUsernameResp struct {
	Message string `json:"message"`
}

type HistoryResp struct {
	Message string       `json:"message"`
	Rumors  []model.Info `json:"rumors"`
}

// HistHandler 用户查询历史记录
func HistHandler(c *gin.Context) {
	resp := HistoryResp{}
	svc := service.NewService(c)
	var rumors []model.Info

	id, _ := strconv.Atoi(c.Query("id"))
	userid := int64(id)
	err := svc.History(&rumors, userid)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = "success"
	resp.Rumors = rumors
	c.JSON(http.StatusOK, resp)
}

// RegHandler 用户注册
func RegHandler(c *gin.Context) {
	var id int64
	var err error
	req := UserRegReq{}
	resp := UserRegResp{}
	svc := service.NewService(c)

	//绑定参数至结构体
	req.Username = c.PostForm("username")
	req.Password = c.PostForm("password")

	//将用户信息插入数据库
	id, err = svc.InsertUser(req.Username, req.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	resp.ID = id
	c.JSON(http.StatusOK, resp)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	req := UserLoginReq{}
	resp := UserLoginResp{}
	svc := service.NewService(c)
	var id int64
	var token string

	req.Username = c.PostForm("username")
	req.Password = c.PostForm("password")

	id = svc.GetIDByName(req.Username, req.Password)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username or Password error",
		})
		return
	}
	token, err := util.GenerateToken(id, req.Username)
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

func UpdatePassword(c *gin.Context) {
	req := UpdatePasswordReq{}
	resp := UpdatePasswordResp{}
	svc := service.NewService(c)
	req.Username = c.PostForm("username")
	req.NewPassword = c.PostForm("new_password")
	err := svc.UpdatePassword(req.Username, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	resp.Message = "update password success"
	c.JSON(http.StatusOK, resp)
}

func UpdateUsername(c *gin.Context) {
	req := UpdateUsernameReq{}
	resp := UpdateUsernameResp{}
	svc := service.NewService(c)

	userId, _ := strconv.Atoi(c.Query("id"))
	req.ID = int64(userId)
	req.NewUsername = c.PostForm("new_username")

	err := svc.UpdateUsername(req.ID, req.NewUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	resp.Message = "update password success"
	c.JSON(http.StatusOK, resp)
}

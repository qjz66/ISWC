package controller

import (
	"Server/model"
	"Server/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserRegReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserRegResp struct {
	ID string `json:"id"`
}
type UserLoginReq struct{}
type UserLoginResp struct{}
type UserHisReq struct{}
type UserHisResp struct{}

// 用户查询历史记录
func HistHandler(c *gin.Context) {

}

// 用户注册
func RegHandler(c *gin.Context) error {
	var id string
	req := UserRegReq{}
	resp := UserRegResp{}
	user := model.User{}
	svc := service.NewService(c)

	//绑定参数至结构体
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return err
	}

	//将用户信息插入数据库
	id, err = svc.InsertUser(user.Username, user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return err
	}

	resp.ID = id
	c.JSON(http.StatusOK, resp)
	return nil
}

// 用户登录
func LoginHandler(c *gin.Context) {

}

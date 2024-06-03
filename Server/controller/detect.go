package controller

import (
	"Server/M3FD"
	"Server/model"
	"Server/service"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DetectReq struct {
	id      int64
	content string `form:"content" binding:"required"`
}
type DetectResp struct {
	Message string `json:"message"`
}

// TextHandler 用户使用文本查询
func TextHandler(c *gin.Context) {
	// 初始化随机数生成器的种子
	// 种子应从当前时间获取，以保证每次运行时生成的随机数序列不同
	rand.Seed(time.Now().UnixNano())

	var info model.Info
	req := DetectReq{}
	resp := DetectResp{}
	svc := service.NewService(c)
	//绑定参数
	userid, _ := strconv.Atoi(c.Query("id"))
	req.id = int64(userid)
	req.content = c.PostForm("content")
	info.ID = req.id
	info.Content = req.content
	//判断是否为谣言
	if M3FD.IsRumor(req.content) {
		//是谣言
		info.Authenticity = "false"
		resp.Message = info.Authenticity
	} else {
		info.Authenticity = "true"
		resp.Message = info.Authenticity
	}

	info.Topic = "tech"
	err := svc.InsertIntoInfo(info, info.Topic)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Message = strconv.FormatFloat(rand.Float64()+0.8, 'f', -1, 64)
	c.JSON(http.StatusOK, resp)
}

// FileHandler 用户使用文件形式查询
func FileHandler(c *gin.Context) {
	// 初始化随机数生成器的种子
	// 种子应从当前时间获取，以保证每次运行时生成的随机数序列不同
	rand.Seed(time.Now().UnixNano())

	var info model.Info
	req := DetectReq{}
	resp := DetectResp{}
	svc := service.NewService(c)
	//绑定参数
	userid, _ := strconv.Atoi(c.Query("id"))
	req.id = int64(userid)
	// 从表单中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		// 处理错误
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// 指定文件应保存的路径
	dst := "./files/" + file.Filename

	// 保存文件到指定路径
	if err = c.SaveUploadedFile(file, dst); err != nil {
		// 处理错误
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// 使用os.ReadFile读取文件内容
	content, err := os.ReadFile(dst)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	info.ID = req.id
	info.Content = string(content)
	//判断是否为谣言
	if M3FD.IsRumor(req.content) {
		//是谣言
		info.Authenticity = "false"
		resp.Message = info.Authenticity
	} else {
		info.Authenticity = "true"
		resp.Message = info.Authenticity
	}

	err = svc.InsertIntoInfo(info, info.Topic)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Message = strconv.FormatFloat(rand.Float64()+0.8, 'f', -1, 64)
	c.JSON(http.StatusOK, resp)
}

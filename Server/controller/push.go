package controller

import (
	"Server/model"
	"Server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PushReq struct {
	Id int64 `json:"id"`
}

type PushResp struct {
	Rumors  []model.Rumor `json:"rumors"`
	Message string        `json:"message"`
}

func PushHandler(c *gin.Context) {
	req := PushReq{}
	resp := PushResp{}
	svc := service.NewService(c)
	var rumors []model.Rumor
	id, _ := strconv.Atoi(c.Query("id"))
	req.Id = int64(id)
	topic, err := svc.GetTopic(req.Id)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	//用户没有检测过信息
	if topic == "" {
		err = svc.Push(&rumors)
	} else {
		err = svc.Push(&rumors)
	}
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.Rumors = rumors
	resp.Message = "push success!"
	c.JSON(http.StatusOK, resp)
}

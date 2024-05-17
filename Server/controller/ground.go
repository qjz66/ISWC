package controller

import (
	"Server/model"
	"Server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FavorReq struct {
	UpdateID int64 `json:"update_id"`
	ID       int64 `json:"id"`
}
type FavorResp struct {
	Message  string `json:"message"`
	Favorite int64  `json:"favorite"`
}

type UpdateReq struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}
type UpdateResp struct {
	Message  string `json:"message"`
	UpdateID int64  `json:"update_id"`
}

type CommentReq struct {
	ID       int64  `json:"id"`
	UpdateID int64  `json:"update_id"`
	Content  string `json:"content"`
}
type CommentResp struct {
	Message string `json:"message"`
}

type GetCommentsReq struct {
	ID       int64 `json:"id"`
	UpdateID int64 `json:"update_id"`
}
type GetCommentsResp struct {
	Message  string          `json:"message"`
	Comments []model.Comment `json:"comments"`
}

type GetBlackListReq struct {
	ID int64 `json:"id"`
}
type GetBlackListResp struct {
	Message   string            `json:"message"`
	Blacklist []model.BlackList `json:"blacklist"`
}

type GetUpdatesResp struct {
	Message string         `json:"message"`
	Updates []model.Update `json:"updates"`
}

func FavoriteHandler(c *gin.Context) {
	req := FavorReq{}
	resp := FavorResp{}
	svc := service.NewService(c)
	var favorite int64

	err := c.ShouldBind(&req)
	updateId, _ := strconv.Atoi(c.Query("update_id"))
	Id, _ := strconv.Atoi(c.Query("id"))
	req.UpdateID = int64(updateId)
	req.ID = int64(Id)

	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	favorite, err = svc.Favorite(req.ID, req.UpdateID)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.Favorite = favorite
	resp.Message = "favorite success!"
	c.JSON(http.StatusOK, resp)
}

func CommentHandler(c *gin.Context) {
	req := CommentReq{}
	resp := CommentResp{}
	comment := model.Comment{ID: req.ID, Content: req.Content, UpdateID: req.UpdateID, Time: time.Now()}
	svc := service.NewService(c)
	err := c.ShouldBind(&req)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	updateId, _ := strconv.Atoi(c.Query("update_id"))
	Id, _ := strconv.Atoi(c.Query("id"))
	comment.UpdateID = int64(updateId)
	comment.ID = int64(Id)
	comment.Content = req.Content
	err = svc.Comment(comment)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = "comment success!"
	c.JSON(http.StatusOK, resp)
}

func GetCommentsHandler(c *gin.Context) {
	var comments []model.Comment
	req := GetCommentsReq{}
	resp := GetCommentsResp{}
	svc := service.NewService(c)
	err := c.ShouldBind(&req)
	if err != nil {
		resp.Message = err.Error()
		resp.Comments = nil
		c.JSON(http.StatusBadRequest, resp)
	}
	updateId, _ := strconv.Atoi(c.Query("update_id"))
	req.UpdateID = int64(updateId)
	err = svc.GetComments(req.UpdateID, &comments)
	if err != nil {
		resp.Message = err.Error()
		resp.Comments = nil
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = "comments success!"
	resp.Comments = comments
	c.JSON(http.StatusOK, resp)
}

func BlackListHandler(c *gin.Context) {
	var blackList []model.BlackList
	req := GetBlackListReq{}
	resp := GetBlackListResp{}
	svc := service.NewService(c)
	err := c.ShouldBind(&req)
	if err != nil {
		resp.Message = err.Error()
		resp.Blacklist = nil
		c.JSON(http.StatusBadRequest, resp)
	}
	err = svc.GetBlackList(&blackList)
	if err != nil {
		resp.Message = err.Error()
		resp.Blacklist = nil
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = "blacklist success!"
	resp.Blacklist = blackList
	c.JSON(http.StatusOK, resp)
}

func UpdateHandler(c *gin.Context) {
	req := UpdateReq{}
	resp := UpdateResp{}
	svc := service.NewService(c)

	err := c.ShouldBind(&req)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	Id, _ := strconv.Atoi(c.Query("id"))
	req.ID = int64(Id)
	updateId, err := svc.Update(req.Content, req.ID)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.UpdateID = updateId
	resp.Message = "update success!"
	c.JSON(http.StatusOK, resp)
}

// GetUpdateListHandler 得到动态列表
func GetUpdateListHandler(c *gin.Context) {
	resp := GetUpdatesResp{}
	svc := service.NewService(c)
	var updateList []model.Update
	err := svc.GetUpdateList(&updateList)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = "update success!"
	resp.Updates = updateList
	c.JSON(http.StatusOK, resp)
}

package controller

import (
	"Server/model"
	"Server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	Id       int64  `json:"id"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	FromName string `json:"fromName"`
}
type UpdateResp struct {
	Message  string `json:"message"`
	UpdateID int64  `json:"updateId"`
}

type CommentReq struct {
	Id       int64  `json:"Id"`
	UpdateId int64  `json:"updateId"`
	FromName string `json:"fromName"`
	ToId     int64  `json:"toId"`
	ToName   string `json:"toName"`
	Content  string `json:"content"`
	Date     string `json:"date"`
}
type CommentResp struct {
	Message   string `json:"message"`
	CommentId int64  `json:"commentId"`
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
	updateId, _ := strconv.Atoi(c.PostForm("updateId"))
	Id, _ := strconv.Atoi(c.Query("id"))
	req.UpdateID = int64(updateId)
	req.ID = int64(Id)

	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	favorite, err = svc.Favorite(req.UpdateID)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Favorite = favorite
	resp.Message = "favorite success!"
	c.JSON(http.StatusOK, resp)
}

func CommentHandler(c *gin.Context) {
	req := CommentReq{}
	resp := CommentResp{}
	comment := model.Comment{}
	svc := service.NewService(c)
	updateId, _ := strconv.Atoi(c.PostForm("updateId"))
	Id, _ := strconv.Atoi(c.Query("id"))
	toId, _ := strconv.Atoi(c.PostForm("toId"))
	req.Id = int64(Id)
	req.UpdateId = int64(updateId)
	req.FromName = c.PostForm("fromName")
	req.ToName = c.PostForm("toName")
	req.Content = c.PostForm("content")
	req.Date = c.PostForm("date")
	req.ToId = int64(toId)
	req.Content = c.PostForm("content")

	comment.UpdateID = req.UpdateId
	comment.AuthorID = req.Id
	comment.AuthorName = req.FromName
	comment.UpdateAuthorID = req.ToId
	comment.AuthorName = req.ToName
	comment.Date = req.Date
	comment.Content = req.Content

	commentId, err := svc.Comment(comment)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Message = "comment success!"
	resp.CommentId = commentId
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
		return
	}
	updateId, _ := strconv.Atoi(c.Query("update_id"))
	req.UpdateID = int64(updateId)
	err = svc.GetComments(req.UpdateID, &comments)
	if err != nil {
		resp.Message = err.Error()
		resp.Comments = nil
		c.JSON(http.StatusBadRequest, resp)
		return
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
		return
	}
	err = svc.GetBlackList(&blackList)
	if err != nil {
		resp.Message = err.Error()
		resp.Blacklist = nil
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Message = "blacklist success!"
	resp.Blacklist = blackList
	c.JSON(http.StatusOK, resp)
}

func UpdateHandler(c *gin.Context) {
	req := UpdateReq{}
	resp := UpdateResp{}
	svc := service.NewService(c)
	var update model.Update
	favorite := new(int64)
	*favorite = 0
	LikeStatus := new(int64)
	*LikeStatus = 0
	showStatus := true
	agree := 2

	Id, _ := strconv.Atoi(c.Query("id"))
	req.Id = int64(Id)
	req.Date = c.PostForm("date")
	req.FromName = c.PostForm("fromName")
	req.Content = c.PostForm("content")
	update.AuthorID = req.Id
	update.AuthorName = req.FromName
	update.Content = req.Content
	update.Date = req.Date
	update.ShowStatus = showStatus
	update.Agree = int64(agree)
	update.LikeStatus = LikeStatus
	update.Favorite = favorite
	updateId, err := svc.Update(update)
	if err != nil {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
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
		return
	}
	resp.Message = "update success!"
	resp.Updates = updateList
	c.JSON(http.StatusOK, resp)
}

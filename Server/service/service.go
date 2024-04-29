package service

import (
	"Server/dao"
	"Server/db"
	"github.com/gin-gonic/gin"
)

type Service struct {
	d *dao.Dao
	c *gin.Context
}

// NewService 创建一个新的Service实例
func NewService(c *gin.Context) *Service {
	return &Service{
		c: c,
		d: dao.NewDao(db.MySqlDB),
	}
}

package middleware

import (
	"Server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AuthByToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		ID, _ := strconv.Atoi(c.Query("id"))
		var userID uint

		//token解析失败
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		//用户不存在
		userID = uint(claims.Id)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		//token解析的ID与用户ID不符
		if userID != uint(ID) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		c.Next()
	}
}

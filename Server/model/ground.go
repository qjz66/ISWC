package model

import "time"

// Comment 评论结构体
type Comment struct {
	ID        int64     `json:"id"`                           //评论作者ID
	CommentID int64     `json:"comment_id" gorm:"primaryKey"` // 评论ID
	UpdateID  int64     `json:"update_id"`                    //所属动态ID
	Content   string    `json:"content"`                      // 评论内容
	Time      time.Time `json:"time"`                         // 评论时间
}

// Update 动态消息结构体
type Update struct {
	UpdateID int64     `json:"update_id" gorm:"primaryKey"` // 动态
	ID       int64     `json:"id"`                          //动态作者ID
	Content  string    `json:"content"`                     // 动态内容
	Favorite *int64    `json:"favorite"`                    // 点赞数
	Time     time.Time `json:"time"`                        // 动态时间
}

// BlackList 黑名单结构体
type BlackList struct {
	ID       int64  `gorm:"primaryKey" json:"id"` //黑名单人员ID
	Platform string `json:"platform"`             //平台
	Account  string `json:"account"`              //账号
	Username string `json:"username"`             //用户名
	Number   int    `json:"number"`               //发送谣言数量
	Icon     string `json:"icon"`                 //头像
}

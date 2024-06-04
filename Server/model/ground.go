package model

// Comment 评论结构体
type Comment struct {
	ID               int64  `json:"id" gorm:"primaryKey"`
	UpdateID         int64  `json:"updateId" gorm:"column:updateId"` // 评论ID
	AuthorID         int64  `json:"fromId" gorm:"column:fromId"`     //评论作者ID
	AuthorName       string `json:"fromName" gorm:"column:fromName"`
	UpdateAuthorID   int64  `json:"toId" gorm:"column:toId"` //所属动态ID
	UpdateAuthorName int64  `json:"toName" gorm:"column:toName"`
	Content          string `json:"content" gorm:"column:content"` // 评论内容
	Date             string `json:"date" gorm:"column:date"`       // 评论时间
}

// Update 动态消息结构体
type Update struct {
	UpdateID   int64     `json:"id"  gorm:"column:id"`        // 动态
	Date       string    `json:"date" gorm:"column:date"`     // 动态时间
	AuthorID   int64     `json:"fromId" gorm:"column:fromId"` //动态作者ID
	AuthorName string    `json:"fromName" gorm:"column:fromName"`
	Favorite   *int64    `json:"likeNum" gorm:"column:likeNum"` // 点赞数
	LikeStatus *int64    `json:"likeStatus" gorm:"column:likeStatus"`
	ShowStatus bool      `json:"showStatus" gorm:"column:showStatus"`
	Agree      int64     `json:"agree" gorm:"column:agree"`
	Content    string    `json:"content" gorm:"column:content"` // 动态内容
	Comment    []Comment `json:"comment" gorm:"-" gorm:"column:comment"`
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

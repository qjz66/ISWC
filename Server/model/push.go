package model

// UserTopic 用来记录用户喜欢的topic
type UserTopic struct {
	ID    int64  `gorm:"primary_key;foreignKey:Users(ID)" json:"id"`
	Topic string `gorm:"primary_key" json:"topic"`
	Num   int64  ` json:"num"`
}

// Rumor 用来记录谣言信息
type Rumor struct {
	Topic    string
	Content  string
	AuthorId string
	Platform string
}

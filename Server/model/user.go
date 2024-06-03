package model

import "gorm.io/gorm"

//用来定义视频流接口所需的model

// MessageReturned 用来存返回状态及信息
type MessageReturned struct {
	gorm.Model
	StatusCode int64  `json:"status"`  // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"message"` // 返回状态描述
}

// User 用来存放用户信息
type User struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	RumorCount *int64 `json:"rumor count"` // 谣言数量，已检测出的谣言数目
	Times      *int64 `json:"times"`       // 数量，检测次数
	Username   string `json:"username"`    // 名字，用户名
	Password   string `json:"password"`
}

// Info 用来存入检测对象的信息
type Info struct {
	InfoId       int64  `gorm:"primaryKey" json:"info_id"` //信息id
	ID           int64  `json:"id"`                        //用户id
	Authenticity string `json:"authenticity"`              // 真实性
	Content      string `json:"content"`                   // 检测内容
	Topic        string `json:"topic"`
}

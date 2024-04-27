package model

import "gorm.io/gorm"

//用来定义视频流接口所需的model

// 用来存返回状态及信息
type MessageReturned struct {
	gorm.Model
	StatusCode int64  `json:"status"`  // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"message"` // 返回状态描述
}

// User用来存放用户信息
type User struct {
	gorm.Model
	RumorCount int64  `json:"rumor count"` // 谣言数量，已检测出的谣言数目
	Times      int64  `json:"times"`       // 数量，检测次数
	Username   string `json:"username"`    // 名字，用户名
}

// Info用来存入检测对象的信息
type Info struct {
	gorm.Model
	Authenticity string                 `json:"authenticity"` // 真实性
	Content      map[string]interface{} `json:"content"`      // 检测内容
}

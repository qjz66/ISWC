package model

import "gorm.io/gorm"

//用来定义视频流接口所需的model

// 用来存返回状态及信息
type MessageReturned struct {
	gorm.Model
	NextTime   int64  `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

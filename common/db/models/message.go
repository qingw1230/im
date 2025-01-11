package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SendID     int64  // 发送者
	ReceiveID  int64  // 接受者
	Type       int    // 发送类型  1私聊  2群聊  3心跳
	Media      int    // 消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string // 消息内容
	CreateTime uint64 // 创建时间
	ReadTime   uint64 // 读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int // 其他数字统计
}

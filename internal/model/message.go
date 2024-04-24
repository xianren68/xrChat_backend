package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Src uint   `json:"src"`
	Tar uint   `json:"tar"`
	Msg string `json:"msg"`
	// 消息类型
	Type int `json:"type"`
}

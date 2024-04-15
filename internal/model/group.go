package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	GroupId uint   `json:"group_id" gorm:"not null"`
	OwnerId uint   `json:"owner_id"`
	Avatar  string `json:"avatar"`
	Desc    string `json:"desc"`
}

package model

import "gorm.io/gorm"

// Relation relation model.
type Relation struct {
	gorm.Model
	OwnerId  uint `json:"owner_id" gorm:"not null"`  // relation src.
	TargetId uint `json:"target_id" gorm:"not null"` // relation target.
	Type     int  `json:"type" gorm:"not null"`      // relation type example(friend,group).
	// if type is friend,the remark is a rename of friend.
	// if type is group,the remark is self name in the group.
	Remark string `json:"remark"`
}

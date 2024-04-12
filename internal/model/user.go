// Package model contains the data models for the application.
package model

import "gorm.io/gorm"

// User represents a user of the application.
type User struct {
	gorm.Model
	Username string `json:"user_name" gorm:"not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required,email"`
	UserId   int64  `json:"user_id" gorm:"auto_increment; not null"`
	Email    string `json:"email" gorm:"not null" validate:"required,email"`
	Phone    string `json:"phone" validate:"matches(^1[3-9]{1}\\d{9}$)"`
	Avatar   string `json:"avatar"`
	Salt     string `json:"salt"`
	Gender   string `json:"gender"`
	Line     string `json:"line"`
}

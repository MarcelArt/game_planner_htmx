package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name   string `gorm:"not null" form:"name"`
	Avatar string `form:"avatar"`
	UserID uint   `gorm:"not null" form:"userId"`

	User *User
}

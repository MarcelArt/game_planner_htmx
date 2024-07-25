package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" form:"username"`
	Email    string `gorm:"not null" form:"email"`
	Password string `gorm:"not null" form:"password"`
}

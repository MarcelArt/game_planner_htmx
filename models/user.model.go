package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" form:"username"`
	Email    string `gorm:"not null;unique" form:"email"`
	Password string `gorm:"not null" form:"password"`
}

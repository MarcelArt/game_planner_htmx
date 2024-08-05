package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Title       string `gorm:"not null" form:"title"`
	Description string `gorm:"not null" form:"description"`
	Picture     string `gorm:"not null" form:"picture"`

	ProfileID uint `gorm:"not null"`

	Profile *Profile
}

package models

import "gorm.io/gorm"

type ConnectedDevice struct {
	gorm.Model
	RefreshToken string `gorm:"not null"`
	UserAgent    string
	Ip           string
	UserID       uint `gorm:"not null"`

	User *User
}

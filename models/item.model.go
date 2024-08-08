package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Picture string `gorm:"not null"`

	GameID uint `gorm:"not null"`

	Game *Game
}

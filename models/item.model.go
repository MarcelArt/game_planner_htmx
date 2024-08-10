package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name    string `gorm:"not null" form:"name"`
	Picture string `gorm:"not null" form:"picture"`

	GameID uint `gorm:"not null" form:"gameId"`

	Game *Game
}

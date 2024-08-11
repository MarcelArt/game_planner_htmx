package models

import "gorm.io/gorm"

type RecipeDetail struct {
	gorm.Model
	InputAmount float64 `gorm:"default:1" form:"inputAmount"`

	RecipeID uint `gorm:"not null" form:"recipeId"`
	ItemID   uint `gorm:"not null" form:"itemId"`

	Recipe *Recipe
	Item   *Item
}

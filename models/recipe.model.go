package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	OutputAmount float64 `gorm:"default:1" form:"outputAmount"`

	ItemID uint `gorm:"not null" form:"itemId"`

	Item *Item
}

type RecipeDto struct {
	gorm.Model
	OutputAmount float64 `gorm:"default:1" form:"outputAmount"`

	ItemID uint `gorm:"not null" form:"itemId"`

	Item          *Item
	RecipeDetails []*RecipeDetail `gorm:"foreignKey:RecipeID" form:"recipeDetails"`
}

func (RecipeDto) TableName() string {
	return "recipes"
}

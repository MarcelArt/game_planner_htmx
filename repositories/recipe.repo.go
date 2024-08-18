package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/gorm"
)

type IRecipeRepo interface {
	IBaseCrudRepo[models.Recipe]
	CreateWithDetail(recipeInput *models.RecipeDto) error
}

type RecipeRepo struct {
	BaseCrudRepo[models.Recipe]
}

func NewRecipeRepo(db *gorm.DB) *RecipeRepo {
	return &RecipeRepo{
		BaseCrudRepo[models.Recipe]{
			db: db,
		},
	}
}

func (r *RecipeRepo) CreateWithDetail(recipeInput *models.RecipeDto) error {
	return r.db.Create(&recipeInput).Error
}

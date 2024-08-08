package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/gorm"
)

type IItemRepo interface {
	IBaseCrudRepo[models.Item]
}

type ItemRepo struct {
	BaseCrudRepo[models.Item]
}

func NewItemRepo(db *gorm.DB) *ItemRepo {
	return &ItemRepo{
		BaseCrudRepo[models.Item]{
			db: db,
		},
	}
}

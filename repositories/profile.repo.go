package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/gorm"
)

type IProfileRepo interface {
	IBaseCrudRepo[models.Profile]
}

type ProfileRepo struct {
	BaseCrudRepo[models.Profile]
}

func NewProfileRepo(db *gorm.DB) *ProfileRepo {
	return &ProfileRepo{
		BaseCrudRepo[models.Profile]{
			db: db,
		},
	}
}

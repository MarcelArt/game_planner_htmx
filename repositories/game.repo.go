package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/gorm"
)

type IGameRepo interface {
	IBaseCrudRepo[models.Game]
}

type GameRepo struct {
	BaseCrudRepo[models.Game]
}

func NewGameRepo(db *gorm.DB) *GameRepo {
	return &GameRepo{
		BaseCrudRepo[models.Game]{
			db: db,
		},
	}
}

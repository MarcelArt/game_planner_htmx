package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type IGameRepo interface {
	IBaseCrudRepo[models.Game]
	GetByProfileID(c *fiber.Ctx, profileID uint) paginate.Page
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

func (r *GameRepo) GetByProfileID(c *fiber.Ctx, profileID uint) paginate.Page {
	var gamesModel []*models.Game
	query := r.db.Where("profile_id = ?", profileID).Model(gamesModel)

	pg := paginate.New()

	return pg.With(query).Request(c.Request()).Response(&gamesModel)
}

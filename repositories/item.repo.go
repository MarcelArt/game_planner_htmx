package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type IItemRepo interface {
	IBaseCrudRepo[models.Item]
	GetByGameID(c *fiber.Ctx, gameID string) paginate.Page
	GetDropdownByGameID(gameID uint) ([]*models.Item, error)
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

func (r *ItemRepo) GetByGameID(c *fiber.Ctx, gameID string) paginate.Page {
	var model []*models.Item
	query := r.db.Model(&model).Where("game_id = ?", gameID)

	pg := paginate.New()
	return pg.With(query).Request(c.Request()).Response(&model)
}

func (r *ItemRepo) GetDropdownByGameID(gameID uint) ([]*models.Item, error) {
	var items []*models.Item
	err := r.db.Where("game_id = ?", gameID).Find(&items).Error

	return items, err
}

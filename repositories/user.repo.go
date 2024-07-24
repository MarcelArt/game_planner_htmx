package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	IBaseCrudRepo[models.User]
}

type UserRepo struct {
	BaseCrudRepo[models.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		BaseCrudRepo[models.User]{
			db: db,
		},
	}
}

func (r *UserRepo) GetByID(id string) (*models.User, error) {
	return &models.User{
		Username: "tes",
	}, nil
}

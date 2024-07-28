package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/gorm"
)

type IProfileRepo interface {
	IBaseCrudRepo[models.Profile]
	GetByUserID(userID uint) (*models.Profile, error)
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

func (r *ProfileRepo) GetByUserID(userID uint) (*models.Profile, error) {
	var profile *models.Profile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	return profile, err
}

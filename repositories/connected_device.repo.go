package repositories

import (
	"github.com/MarcelArt/game_planner_htmx/models"
	"gorm.io/gorm"
)

type IConnectedDeviceRepo interface {
	IBaseCrudRepo[models.ConnectedDevice]
	GetByToken(refreshToken string) (*models.ConnectedDevice, error)
}

type ConnectedDeviceRepo struct {
	BaseCrudRepo[models.ConnectedDevice]
}

func NewConnectedDeviceRepo(db *gorm.DB) *ConnectedDeviceRepo {
	return &ConnectedDeviceRepo{
		BaseCrudRepo[models.ConnectedDevice]{
			db: db,
		},
	}
}

func (r *ConnectedDeviceRepo) GetByToken(refreshToken string) (*models.ConnectedDevice, error) {
	var device *models.ConnectedDevice
	err := r.db.Where("refresh_token = ?", refreshToken).First(&device).Error
	return device, err
}

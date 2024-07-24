package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBaseCrudRepo[TModel any] interface {
	Create(input *TModel) (*TModel, error)
	Read(c *fiber.Ctx, dest interface{}) paginate.Page
	Update(id string, input *TModel) error
	Delete(id string) (*TModel, error)
	GetByID(id string) (*TModel, error)
}

type BaseCrudRepo[TModel any] struct {
	db *gorm.DB
}

func NewBaseCrudRepo[TModel any](db *gorm.DB) *BaseCrudRepo[TModel] {
	return &BaseCrudRepo[TModel]{
		db: db,
	}
}

func (r *BaseCrudRepo[TModel]) Create(input *TModel) (*TModel, error) {
	err := r.db.Create(&input).Error

	return input, err
}

func (r *BaseCrudRepo[TModel]) Read(c *fiber.Ctx, dest interface{}) paginate.Page {
	pg := paginate.New()

	query := r.db.Model(dest)
	page := pg.With(query).Request(c.Request()).Response(&dest)

	return page
}

func (r *BaseCrudRepo[TModel]) Update(id string, input *TModel) error {
	return r.db.Model(input).Where("id = ?", id).Updates(input).Error
}

func (r *BaseCrudRepo[TModel]) Delete(id string) (*TModel, error) {
	var model *TModel
	err := r.db.Clauses(clause.Returning{}).Delete(&model, id).Error

	return model, err
}

func (r *BaseCrudRepo[TModel]) GetByID(id string) (*TModel, error) {
	var model *TModel
	err := r.db.First(&model, id).Error
	return model, err
}

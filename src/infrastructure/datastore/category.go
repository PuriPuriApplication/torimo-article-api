package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/repository"
)

type CategoryDatastore struct {
	db *gorm.DB
}

func NewCategoryDatastore(d *gorm.DB) repository.ICategoryRepository {
	return &CategoryDatastore{
		db: d,
	}
}

func (c *CategoryDatastore) FindByIdIn(ids []uint64) []model.Category {
	category := []model.Category{}
	c.db.Find(&category, "id IN (?)", ids)
	return category
}

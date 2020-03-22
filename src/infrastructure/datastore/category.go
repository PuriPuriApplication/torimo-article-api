package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/article_repository"
)

type CategoryDatastore struct {
	db *gorm.DB
}

func NewCategoryDatastore(d *gorm.DB) article_repository.ICategoryRepository {
	return &CategoryDatastore{
		db: d,
	}
}

func (c *CategoryDatastore) FindByIdIn(ids []uint64) []article_model.Category {
	var category []article_model.Category
	c.db.Find(&category, "id IN (?)", ids)
	return category
}

package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/thoas/go-funk"

	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/repository"
)

type ArticleCategoryDatastore struct {
	db *gorm.DB
}

func NewArticleCategoryDatastore(d *gorm.DB) repository.IArticleCategoryRepository {
	return &ArticleCategoryDatastore{
		db: d,
	}
}

func (ac *ArticleCategoryDatastore) SaveCategoriesForArticle(articleID uint64, categoryIDs []uint64) {
	funk.ForEach(categoryIDs, func(categoryID uint64) {
		article := model.Article{ID: articleID}
		category := model.Category{ID: categoryID}
		ac.db.Model(&article).Association("Categories").Append(&category)
	})
}

func (ac *ArticleCategoryDatastore) DeleteByAriticleID(articleID uint64) {
	article := model.Article{ID: articleID}
	ac.db.Model(&article).Association("Categories").Clear()
}

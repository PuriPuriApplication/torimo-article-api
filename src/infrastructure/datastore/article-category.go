package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/thoas/go-funk"

	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/article_repository"
)

type ArticleCategoryDatastore struct {
	db *gorm.DB
}

func NewArticleCategoryDatastore(d *gorm.DB) article_repository.IArticleCategoryRepository {
	return &ArticleCategoryDatastore{
		db: d,
	}
}

func (ac *ArticleCategoryDatastore) SaveCategoriesForArticle(articleID uint64, categoryIDs []uint64) {
	funk.ForEach(categoryIDs, func(categoryID uint64) {
		article := article_model.Article{ID: articleID}
		category := article_model.Category{ID: categoryID}
		ac.db.Model(&article).Association("Categories").Append(&category)
	})
}

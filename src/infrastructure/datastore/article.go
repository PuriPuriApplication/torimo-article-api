package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/repository"
)

type ArticleDatastore struct {
	db *gorm.DB
}

func NewArticleDatastore(d *gorm.DB) repository.IArticleRepository {
	return &ArticleDatastore{
		db: d,
	}
}

func (a *ArticleDatastore) Save(article *model.Article) {
	a.db.Create(&article)
}

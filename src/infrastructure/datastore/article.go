package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/thoas/go-funk"

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

func (a *ArticleDatastore) Create(article *model.Article) {
	a.db.Create(&article)
}

func (a *ArticleDatastore) Update(article *model.Article)  {
	a.db.Save(&article)
}

func (a *ArticleDatastore) FindAll() []model.Article {
	var articles []model.Article
	a.db.Select("id").Find(&articles)

	results := funk.Map(articles, func(article model.Article) model.Article {
		a.db.First(&article).Related(&article.User, "User").Related(&article.Shop, "Shop").Related(&article.Categories, "Categories")
		return article
	}).([]model.Article)

	return results
}

func (a *ArticleDatastore) FindOne(ID uint64) model.Article {
	article := model.Article{ID: ID}
	a.db.First(&article).Related(&article.User, "User").Related(&article.Shop, "Shop").Related(&article.Categories, "Categories")
	return article
}

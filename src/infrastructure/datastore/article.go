package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/thoas/go-funk"

	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/article_repository"
)

type ArticleDatastore struct {
	db *gorm.DB
}

func NewArticleDatastore(d *gorm.DB) article_repository.IArticleRepository {
	return &ArticleDatastore{
		db: d,
	}
}

func (a *ArticleDatastore) Save(article *article_model.Article) {
	a.db.Create(&article)
}

func (a *ArticleDatastore) FindAll() []article_model.Article {
	var articles []article_model.Article
	a.db.Select("id").Find(&articles)

	results := funk.Map(articles, func(article article_model.Article) article_model.Article {
		a.db.First(&article).Related(&article.User, "User").Related(&article.Shop, "Shop").Related(&article.Categories, "Categories")
		return article
	}).([]article_model.Article)

	return results
}

func (a *ArticleDatastore) FindOne(ID uint64) article_model.Article {
	article := article_model.Article{ID: ID}
	a.db.First(&article).Related(&article.User, "User").Related(&article.Shop, "Shop").Related(&article.Categories, "Categories")
	return article
}

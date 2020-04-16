package repository

import (
	"torimo-article-api/src/domain/model"
)

type IArticleRepository interface {
	Create(article *model.Article)
	Update(article *model.Article)
	FindAll() []model.Article
	FindOne(ID uint64) model.Article
}

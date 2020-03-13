package repository

import (
	"torimo-article-api/src/domain/model"
)

type IArticleRepository interface {
	Save(article *model.Article)
	FindAll() []model.Article
}

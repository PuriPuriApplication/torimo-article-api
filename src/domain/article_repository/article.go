package article_repository

import (
	"torimo-article-api/src/domain/article_model"
)

type IArticleRepository interface {
	Save(article *article_model.Article)
	FindAll() []article_model.Article
	FindOne(ID uint64) article_model.Article
}

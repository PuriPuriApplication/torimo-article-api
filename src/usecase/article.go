package usecase

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/handler/request"
)

type IArticleUsecase interface {
	Create(requestArticle *request.CreateArticleRequest) uint64
	GetAll() []article_model.Article
	GetOne(ID uint64) article_model.Article
}

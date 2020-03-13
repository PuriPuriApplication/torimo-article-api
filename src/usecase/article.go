package usecase

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/handler/request"
)

type IArticleUsecase interface {
	Create(requestArticle *request.CreateArticleRequest) uint64
	GetAll() []model.Article
}

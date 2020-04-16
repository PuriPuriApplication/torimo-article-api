package usecase

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/handler/request"
)

type IArticleUsecase interface {
	Create(requestArticle *request.UpsertArticleRequest) uint64
	Update(id uint64, requestArticle *request.UpsertArticleRequest)
	GetAll() []model.Article
	GetOne(ID uint64) model.Article
}

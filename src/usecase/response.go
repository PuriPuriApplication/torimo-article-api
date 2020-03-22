package usecase

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
)

type IResponseUsecase interface {
	CreateAll(a []article_model.Article) *[]response_model.ResponseArticle
	CreateOne(a *article_model.Article) *response_model.ResponseArticle
}

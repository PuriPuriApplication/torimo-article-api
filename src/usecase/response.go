package usecase

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
)

type IResponseUsecase interface {
	CreateAll(a []model.Article) *[]response.ResponseArticle
	CreateOne(a *model.Article) *response.ResponseArticle
}

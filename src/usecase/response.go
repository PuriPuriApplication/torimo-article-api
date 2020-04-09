package usecase

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
)

type IResponseUsecase interface {
	MappingAll(a []model.Article) *[]response.ResponseArticle
	MappingOne(a *model.Article) *response.ResponseArticle
}

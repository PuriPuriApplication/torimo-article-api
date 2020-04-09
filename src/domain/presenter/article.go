package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/handler/response"
)

type IArticlePresenter interface {
	MappingAll(a *[]model.Article) []*response.ResponseArticle
	MappingOne(a *model.Article) *response.ResponseArticle
}

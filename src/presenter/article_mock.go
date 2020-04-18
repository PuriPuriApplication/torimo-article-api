package presenter

import (
	"fmt"
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/presenter"
	"torimo-article-api/src/handler/response"
)

type ArticleMockPresenter struct{}

func NewArticleMockPresenter() presenter.IArticlePresenter {
	return &ArticleMockPresenter{}
}

func (amp *ArticleMockPresenter) MappingAll(articles *[]model.Article) []*response.ResponseArticle {
	fmt.Println("Mapping of ", articles)
	return []*response.ResponseArticle{
		&response.ResponseArticle{},
	}
}

func (amp *ArticleMockPresenter) MappingOne(a *model.Article) *response.ResponseArticle {
	fmt.Println("Mapping of ", a)
	return &response.ResponseArticle{}
}

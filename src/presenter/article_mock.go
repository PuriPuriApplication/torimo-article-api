package presenter

import (
	"fmt"
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/presenter"
	"torimo-article-api/src/handler/response"

	"github.com/thoas/go-funk"
)

type ArticleMockPresenter struct{}

func NewArticleMockPresenter() presenter.IArticlePresenter {
	return &ArticlePresenter{}
}

func (amp *ArticleMockPresenter) MappingAll(articles *[]model.Article) []*response.ResponseArticle {
	fmt.Println("Mapping of ", articles)
	return []&response.ResponseArticle{}
}

func (amp *ArticleMockPresenter) MappingOne(a *model.Article) *response.ResponseArticle {
	fmt.Println("Mapping of ", articles)
	return &response.ResponseArticle{}
}

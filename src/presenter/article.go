package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/presenter"
	"torimo-article-api/src/handler/response"

	"github.com/thoas/go-funk"
)

type ArticlePresenter struct{}

func NewArticlePresenter() presenter.IArticlePresenter {
	return &ArticlePresenter{}
}

func (ri *ArticlePresenter) MappingAll(articles *[]model.Article) []*response.ResponseArticle {
	return funk.Map(articles, func(a *model.Article) *response.ResponseArticle {
		return &response.ResponseArticle{
			ID:         a.ID,
			Title:      a.Title,
			Body:       a.Body,
			Status:     a.Status,
			User:       convertUserResponse(&a.User),
			Shop:       convertShopResponse(&a.Shop),
			Categories: convertCategoryResponse(a.Categories),
		}
	}).([]*response.ResponseArticle)
}

func (ri *ArticlePresenter) MappingOne(a *model.Article) *response.ResponseArticle {
	return &response.ResponseArticle{
		ID:         a.ID,
		Title:      a.Title,
		Body:       a.Body,
		Status:     a.Status,
		User:       convertUserResponse(&a.User),
		Shop:       convertShopResponse(&a.Shop),
		Categories: convertCategoryResponse(a.Categories),
	}
}

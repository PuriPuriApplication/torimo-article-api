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
			ID:         v.ID,
			Title:      v.Title,
			Body:       v.Body,
			Status:     v.Status,
			User:       convertUserResponse(&v.User),
			Shop:       convertShopResponse(&v.Shop),
			Categories: convertCategoryResponse(v.Categories),
		}
	}).([]*response.ResponseArticle)
}

func (ri *ArticlePresenter) MappingOne(a *model.Article) *response.ResponseArticle {
	return &response.ResponseArticle{
		ID:         articles.ID,
		Title:      articles.Title,
		Body:       articles.Body,
		Status:     articles.Status,
		User:       convertUserResponse(&articles.User),
		Shop:       convertShopResponse(&articles.Shop),
		Categories: convertCategoryResponse(articles.Categories),
	}
}

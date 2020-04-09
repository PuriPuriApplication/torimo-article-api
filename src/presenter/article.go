package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/presenter"
	"torimo-article-api/src/handler/response"
)

type ArticlePresenter struct{}

func NewArticlePresenter() presenter.IArticlePresenter {
	return &ArticlePresenter{}
}

func (ri *ArticlePresenter) MappingAll(a *[]model.Article) []*response.ResponseArticle {
	var responses []*response.ResponseArticle
	for _, v := range *a {
		response := &response.ResponseArticle{
			ID:         v.ID,
			Title:      v.Title,
			Body:       v.Body,
			Status:     v.Status,
			User:       convertUserResponse(&v.User),
			Shop:       convertShopResponse(&v.Shop),
			Categories: convertCategoryResponse(v.Categories),
		}
		responses = append(responses, response)
	}

	return responses
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

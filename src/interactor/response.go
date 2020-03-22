package interactor

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
	"torimo-article-api/src/domain/response_repository"
	"torimo-article-api/src/usecase"
)

type ResponseInteractor struct {
	UserResponseRepository     response_repository.IUserRepository
	ShopResponseRepository     response_repository.IShopRepository
	CategoryResponseRepository response_repository.ICategoryRepository
}

func NewResponseInteractor(
	ur response_repository.IUserRepository,
	sr response_repository.IShopRepository,
	cr response_repository.ICategoryRepository,
) usecase.IResponseUsecase {
	return &ResponseInteractor{
		UserResponseRepository:     ur,
		ShopResponseRepository:     sr,
		CategoryResponseRepository: cr,
	}
}

func (ri *ResponseInteractor) CreateAll(a []article_model.Article) *[]response_model.ResponseArticle {
	var responses []response_model.ResponseArticle
	for _, v := range a {
		response := response_model.ResponseArticle{
			ID:         v.ID,
			Title:      v.Title,
			Body:       v.Body,
			Status:     v.Status,
			User:       ri.UserResponseRepository.Convert(&v.User),
			Shop:       ri.ShopResponseRepository.Convert(&v.Shop),
			Categories: ri.CategoryResponseRepository.Convert(v.Categories),
		}
		responses = append(responses, response)
	}

	return &responses
}

func (ri *ResponseInteractor) CreateOne(a *article_model.Article) *response_model.ResponseArticle {
	return &response_model.ResponseArticle{
		ID:         a.ID,
		Title:      a.Title,
		Body:       a.Body,
		Status:     a.Status,
		User:       ri.UserResponseRepository.Convert(&a.User),
		Shop:       ri.ShopResponseRepository.Convert(&a.Shop),
		Categories: ri.CategoryResponseRepository.Convert(a.Categories),
	}
}

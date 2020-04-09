package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
	"torimo-article-api/src/domain/response_repository"
)

type CategoryDataConvert struct{}

func NewCategoryDataConvert() response_repository.ICategoryRepository {
	return &CategoryDataConvert{}
}

func (c *CategoryDataConvert) Convert(category []model.Category) []*response.ResponseCategory {
	var categories []*response.ResponseCategory
	for _, v := range category {
		elem := &response.ResponseCategory{
			ID:         v.ID,
			Name:       v.Name,
			CreateUser: v.CreateUser,
			IsDeleted:  v.IsDeleted,
		}
		categories = append(categories, elem)
	}

	return categories
}

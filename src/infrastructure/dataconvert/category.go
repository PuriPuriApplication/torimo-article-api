package dataconvert

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
	"torimo-article-api/src/domain/response_repository"
)

type CategoryDataConvert struct{}

func NewCategoryDataConvert() response_repository.ICategoryRepository {
	return &CategoryDataConvert{}
}

func (c *CategoryDataConvert) Convert(category []article_model.Category) []*response_model.ResponseCategory {
	var categories []*response_model.ResponseCategory
	for _, v := range category {
		elem := &response_model.ResponseCategory{
			ID:         v.ID,
			Name:       v.Name,
			CreateUser: v.CreateUser,
			IsDeleted:  v.IsDeleted,
		}
		categories = append(categories, elem)
	}

	return categories
}

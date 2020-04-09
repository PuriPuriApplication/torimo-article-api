package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/handler/response"

	"github.com/thoas/go-funk"
)

func convertCategoryResponse(categories []model.Category) []*response.ResponseCategory {
	results := funk.Map(categories, func(category model.Category) *response.ResponseCategory {
		return &response.ResponseCategory{
			ID:         category.ID,
			Name:       category.Name,
			CreateUser: category.CreateUser,
			IsDeleted:  category.IsDeleted,
		}
	}).([]*response.ResponseCategory)

	return results
}

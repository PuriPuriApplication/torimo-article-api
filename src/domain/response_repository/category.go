package response_repository

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
)

type ICategoryRepository interface {
	Convert(category []model.Category) []*response.ResponseCategory
}

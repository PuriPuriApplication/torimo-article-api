package response_repository

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
)

type ICategoryRepository interface {
	Convert(category []article_model.Category) []*response_model.ResponseCategory
}

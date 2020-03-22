package article_repository

import "torimo-article-api/src/domain/article_model"

type ICategoryRepository interface {
	FindByIdIn(ids []uint64) []article_model.Category
}

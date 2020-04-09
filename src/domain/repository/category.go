package repository

import "torimo-article-api/src/domain/model"

type ICategoryRepository interface {
	FindByIdIn(ids []uint64) []model.Category
}

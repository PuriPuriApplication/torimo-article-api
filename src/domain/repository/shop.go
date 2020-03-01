package repository

import "torimo-article-api/src/domain/model"

type IShopRepository interface {
	FindById(id uint64) model.Shop
}

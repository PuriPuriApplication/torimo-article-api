package response_repository

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
)

type IShopRepository interface {
	Convert(shop *article_model.Shop) *response_model.ResponseShop
}

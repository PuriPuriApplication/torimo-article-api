package response_repository

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
)

type IShopRepository interface {
	Convert(shop *model.Shop) *response.ResponseShop
}

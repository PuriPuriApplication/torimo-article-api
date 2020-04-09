package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
	"torimo-article-api/src/domain/response_repository"
)

type ShopDataConvert struct{}

func NewShopDataConvert() response_repository.IShopRepository {
	return &ShopDataConvert{}
}

func (s *ShopDataConvert) Convert(shop *model.Shop) *response.ResponseShop {
	return &response.ResponseShop{
		ID:         shop.ID,
		Name:       shop.Name,
		StationID:  shop.StationID,
		CreateUser: shop.CreateUser,
		IsDeleted:  shop.IsDeleted,
	}
}

package dataconvert

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
	"torimo-article-api/src/domain/response_repository"
)

type ShopDataConvert struct{}

func NewShopDataConvert() response_repository.IShopRepository {
	return &ShopDataConvert{}
}

func (s *ShopDataConvert) Convert(shop *article_model.Shop) *response_model.ResponseShop {
	return &response_model.ResponseShop{
		ID:         shop.ID,
		Name:       shop.Name,
		StationID:  shop.StationID,
		CreateUser: shop.CreateUser,
		IsDeleted:  shop.IsDeleted,
	}
}

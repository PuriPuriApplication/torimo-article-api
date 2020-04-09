package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/handler/response"
)

func convertShopResponse(shop *model.Shop) *response.ResponseShop {
	return &response.ResponseShop{
		ID:         shop.ID,
		Name:       shop.Name,
		StationID:  shop.StationID,
		CreateUser: shop.CreateUser,
		IsDeleted:  shop.IsDeleted,
	}
}

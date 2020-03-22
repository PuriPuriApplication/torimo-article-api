package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/article_repository"
)

type ShopDatastore struct {
	db *gorm.DB
}

func NewShopDatastore(d *gorm.DB) article_repository.IShopRepository {
	return &ShopDatastore{
		db: d,
	}
}

func (s *ShopDatastore) FindById(id uint64) article_model.Shop {
	shop := article_model.Shop{}
	s.db.Find(&shop, "id=?", id)
	return shop
}

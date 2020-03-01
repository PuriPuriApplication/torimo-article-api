package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/repository"
)

type ShopDatastore struct {
	db *gorm.DB
}

func NewShopDatastore(d *gorm.DB) repository.IShopRepository {
	return &ShopDatastore{
		db: d,
	}
}

func (s *ShopDatastore) FindById(id uint64) model.Shop {
	shop := model.Shop{}
	s.db.Find(&shop, "id=?", id)
	return shop
}

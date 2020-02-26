package model

import (
	"time"
)

type Category struct {
	ID      int `json:"id"`
	Name string `json:"name"`
	CreateUser int64 `json:"createUser"`
	// Articles []*Article `gorm:"many2many:article_categories;"`
	IsDeleted bool `json:"isDeleted"`
	CreateAt *time.Time `json:"createAt"`
	DeleteAt *time.Time `json:"deleteAt"`
}

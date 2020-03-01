package model

import (
	"time"
)

type Category struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	CreateUser uint64     `json:"createUser"`
	IsDeleted  bool       `json:"isDeleted"`
	CreateAt   *time.Time `json:"createAt"`
	DeleteAt   *time.Time `json:"deleteAt"`
	// Articles []*Article `gorm:"many2many:article_categories;"`
}

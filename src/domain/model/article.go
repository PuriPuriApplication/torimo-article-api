package model

import (
	"time"
)

type Article struct {
	ID         uint64
	Title      string
	Body       string
	Status     string
	UserID     uint64
	User       User
	ShopID     uint64
	Shop       Shop
	Categories []Category `gorm:"many2many:article_categories"`
	CreateAt   time.Time
	UpdateAt   *time.Time
}

package model

import (
	"time"
)

type Article struct {
	ID         uint64     `json:"id"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Status     string     `json:"status"`
	UserID     uint64     `json:"-"`
	User       User       `json:"user"`
	ShopID     uint64     `json:"-"`
	Shop       Shop       `json:"shop"`
	Categories []Category `json:"categories" gorm:"many2many:article_categories"`
	CreateAt   time.Time  `json:"createAt"`
	UpdateAt   *time.Time `json:"updateAt"`
}

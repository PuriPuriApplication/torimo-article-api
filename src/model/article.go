package model

import (
	"time"
)

type Article struct {
	ID      int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Status string `json:"status"`
	UserID int64 `json:"-"`
	User User `json:"user"`
	ShopID int64 `json:"-"`
	Shop Shop `json:"shop"`
	Categories []Category `gorm:"many2many:article_categories" json:"categories"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt *time.Time `json:"updateAt"`
}

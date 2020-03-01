package handler

import (
	"github.com/jinzhu/gorm"
)

// Handler is Struct
type Handler struct {
	db *gorm.DB
}

func NewHandler(d *gorm.DB) *Handler {
	return &Handler{
		db: d,
	}
}

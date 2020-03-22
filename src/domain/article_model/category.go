package article_model

import (
	"time"
)

type Category struct {
	ID         uint64
	Name       string
	CreateUser uint64
	IsDeleted  bool
	CreateAt   *time.Time
	DeleteAt   *time.Time
}

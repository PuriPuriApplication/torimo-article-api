package article_model

import (
	"time"
)

type Shop struct {
	ID         uint64
	Name       string
	StationID  uint64
	CreateUser uint64
	IsDeleted  bool
	CreateAt   time.Time
	DeleteAt   *time.Time
}

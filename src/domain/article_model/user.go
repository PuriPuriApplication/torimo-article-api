package article_model

import (
	"time"
)

type User struct {
	ID                  uint64
	Name                string
	ExternalServiceID   string
	ExternalServiceType string
	IsDeleted           bool
	CreateAt            time.Time
	UpdateAt            *time.Time
	DeleteAt            *time.Time
}

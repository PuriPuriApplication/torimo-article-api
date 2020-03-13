package model

import (
	"time"
)

type User struct {
	ID                  uint64     `json:"id"`
	Name                string     `json:"name"`
	ExternalServiceID   string     `json:"externalServiceID"`
	ExternalServiceType string     `json:"externalServiceType"`
	IsDeleted           bool       `json:"isDeleted"`
	CreateAt            time.Time  `json:"createAt"`
	UpdateAt            *time.Time `json:"updateAt"`
	DeleteAt            *time.Time `json:"deleteAt"`
}

package model

import (
	"time"
)

type Shop struct {
	ID      int `json:"id"`
	Name string `json:"name"`
	StationID string `json:"stationId"`
	CreateUser int64 `json:"createUser"`
	IsDeleted bool `json:"isDeleted"`
	CreateAt time.Time `json:"createAt"`
	DeleteAt *time.Time `json:"updateAt"`
}

package model

import (
	"time"
)

type Shop struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	StationID  int64      `json:"stationId"`
	CreateUser int64      `json:"createUser"`
	IsDeleted  bool       `json:"isDeleted"`
	CreateAt   time.Time  `json:"createAt"`
	DeleteAt   *time.Time `json:"updateAt"`
}

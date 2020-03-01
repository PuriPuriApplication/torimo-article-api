package model

import (
	"time"
)

type Shop struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	StationID  uint64     `json:"stationId"`
	CreateUser uint64     `json:"createUser"`
	IsDeleted  bool       `json:"isDeleted"`
	CreateAt   time.Time  `json:"createAt"`
	DeleteAt   *time.Time `json:"updateAt"`
}

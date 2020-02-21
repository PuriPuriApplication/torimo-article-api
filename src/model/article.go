package model

type (
	article struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}
)

var (
	articles = map[int]*article{}
	seq      = 1
)

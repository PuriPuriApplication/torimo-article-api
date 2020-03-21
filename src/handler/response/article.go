package response

import "torimo-article-api/src/domain/model"

type User struct {
	ID                  uint64 `json:"id"`
	Name                string `json:"name"`
	ExternalServiceID   string `json:"externalServiceID"`
	ExternalServiceType string `json:"externalServiceType"`
	IsDeleted           bool   `json:"isDeleted"`
}

func convertUser(u *model.User) *User {
	return &User{
		ID:                  u.ID,
		Name:                u.Name,
		ExternalServiceID:   u.ExternalServiceID,
		ExternalServiceType: u.ExternalServiceType,
		IsDeleted:           u.IsDeleted,
	}
}

type Shop struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	StationID  uint64 `json:"stationId"`
	CreateUser uint64 `json:"createUser"`
	IsDeleted  bool   `json:"isDeleted"`
}

func convertShop(s *model.Shop) *Shop {
	return &Shop{
		ID:         s.ID,
		Name:       s.Name,
		StationID:  s.StationID,
		CreateUser: s.CreateUser,
		IsDeleted:  s.IsDeleted,
	}
}

type Category struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	CreateUser uint64 `json:"createUser"`
	IsDeleted  bool   `json:"isDeleted"`
}

func convertCategory(c []model.Category) []*Category {
	var categories []*Category
	for _, v := range c {
		category := &Category{
			ID:         v.ID,
			Name:       v.Name,
			CreateUser: v.CreateUser,
			IsDeleted:  v.IsDeleted,
		}

		categories = append(categories, category)
	}
	return categories
}

type SelectedArticle struct {
	ID         uint64      `json:"id"`
	Title      string      `json:"title"`
	Body       string      `json:"body"`
	Status     string      `json:"status"`
	User       *User       `json:"user"`
	Shop       *Shop       `json:"shop"`
	Categories []*Category `json:"categories"`
}

func NewSelectedArticle(a *model.Article) *SelectedArticle {
	return &SelectedArticle{
		ID:         a.ID,
		Title:      a.Title,
		Body:       a.Body,
		Status:     a.Status,
		User:       convertUser(&a.User),
		Shop:       convertShop(&a.Shop),
		Categories: convertCategory(a.Categories),
	}
}

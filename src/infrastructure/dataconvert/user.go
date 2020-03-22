package dataconvert

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
	"torimo-article-api/src/domain/response_repository"
)

type UserDataConvert struct{}

func NewUserDataConvert() response_repository.IUserRepository {
	return &UserDataConvert{}
}

func (u *UserDataConvert) Convert(user *article_model.User) *response_model.ResponseUser {
	return &response_model.ResponseUser{
		ID:                  user.ID,
		Name:                user.Name,
		ExternalServiceID:   user.ExternalServiceID,
		ExternalServiceType: user.ExternalServiceType,
		IsDeleted:           user.IsDeleted,
	}
}

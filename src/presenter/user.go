package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
	"torimo-article-api/src/domain/response_repository"
)

type UserDataConvert struct{}

func NewUserDataConvert() response_repository.IUserRepository {
	return &UserDataConvert{}
}

func (u *UserDataConvert) Convert(user *model.User) *response.ResponseUser {
	return &response.ResponseUser{
		ID:                  user.ID,
		Name:                user.Name,
		ExternalServiceID:   user.ExternalServiceID,
		ExternalServiceType: user.ExternalServiceType,
		IsDeleted:           user.IsDeleted,
	}
}

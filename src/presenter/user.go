package presenter

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/handler/response"
)

func convertUserResponse(user *model.User) *response.ResponseUser {
	return &response.ResponseUser{
		ID:                  user.ID,
		Name:                user.Name,
		ExternalServiceID:   user.ExternalServiceID,
		ExternalServiceType: user.ExternalServiceType,
		IsDeleted:           user.IsDeleted,
	}
}

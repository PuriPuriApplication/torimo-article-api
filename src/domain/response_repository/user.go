package response_repository

import (
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/presenter/response"
)

type IUserRepository interface {
	Convert(user *model.User) *response.ResponseUser
}
